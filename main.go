package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

// config for the overall application
type discordConfig struct {
	GuildID  string `yaml:"guildId"`
	BotToken string `yaml:"botToken"`
	TenorKey string `yaml:"tenorKey"`
}

type heroConfig struct {
	Heroes *[]*hero `yaml:"heroes"`
}

// config for each hero
type hero struct {
	Name    string `yaml:"name"`
	Role    string `yaml:"role"`
	Stadium bool   `yaml:"stadium"`
}

var (
	tenorSearchUrl = "https://tenor.googleapis.com/v2/search"
	commands       = []*discordgo.ApplicationCommand{
		{
			Name:        "hero_roll",
			Description: "roll a random hero",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "role",
					Description: "Roles to consider",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "all",
							Value: "all",
						},
						{
							Name:  "dps",
							Value: "dps",
						},
						{
							Name:  "support",
							Value: "support",
						},
						{
							Name:  "tank",
							Value: "tank",
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "stadium",
					Description: "whether to consider only stadium heroes or not",
					Required:    false,
				},
			},
		},
	}
)

// establish flags and read config values
func readAppFlags() (string, string) {
	const (
		defaultValue    = ""
		discordOptUsage = "path to Discord config file"
		heroOptUsage    = "path to hero data file"
	)

	var discordConfigPath, heroConfigPath string
	flag.StringVar(&discordConfigPath, "file", defaultValue, discordOptUsage)
	flag.StringVar(&discordConfigPath, "f", defaultValue, discordOptUsage)
	flag.StringVar(&heroConfigPath, "data-file-name", defaultValue, heroOptUsage)
	flag.StringVar(&heroConfigPath, "d", defaultValue, heroOptUsage)
	flag.Parse()
	return discordConfigPath, heroConfigPath
}

func readConfigValues[T interface{}](filePath string) (*T, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config *T
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// connect to discord
func initializeDiscordSession(config *discordConfig) (*discordgo.Session, error) {
	fmt.Println("initializing discord connection")
	return discordgo.New("Bot " + config.BotToken)
}

func getOptionsFromDiscordInteraction(i *discordgo.InteractionCreate) map[string]any {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]any, len(options))
	for _, opt := range options {
		switch opt.Type {
		case discordgo.ApplicationCommandOptionString:
			optionMap[opt.Name] = opt.StringValue()
		case discordgo.ApplicationCommandOptionBoolean:
			optionMap[opt.Name] = opt.BoolValue()
		}
	}
	return optionMap
}

// TODO improve json parsing with structured data
func getGifFromJson(jsonBytes []byte) (string, error) {
	var result map[string]any
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		fmt.Println("error parsing JSON", err)
	}
	results, ok := result["results"]
	if !ok {
		return "", fmt.Errorf("no results found in json bytes %v", string(jsonBytes[:]))
	}
	resultArray := results.([]interface{})
	if len(resultArray) == 0 {
		return "", fmt.Errorf("no results found in json bytes %v", string(jsonBytes[:]))
	}
	mediaFormats := resultArray[0].(map[string]any)["media_formats"]
	gifBlock, ok := mediaFormats.(map[string]any)["gif"]
	if !ok {
		return "", fmt.Errorf("no gif results found in first result of json bytes %v", string(jsonBytes[:]))
	}
	gifUrl, ok := gifBlock.(map[string]any)["url"]
	if !ok {
		return "", fmt.Errorf("no gif url found in first result of json bytes %v", string(jsonBytes[:]))
	}
	stringUrl := gifUrl.(string)
	return stringUrl, nil
}

func getHeroGif(heroName string, apiKey string) (string, error) {
	queryString := fmt.Sprintf("%v overwatch", heroName)
	req, err := http.NewRequest("GET", tenorSearchUrl, nil)
	if err != nil {
		return "", err
	}
	queryParams := map[string]string{
		"q":             queryString,
		"key":           apiKey,
		"locale":        "en_US",
		"media_filter":  "gif",
		"limit":         "1",
		"contentFilter": "off",
		"ar_range":      "all",
		"random":        "true",
	}
	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req.Close = false
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error requesting gif %v\n", err)
		return "", err
	}
	return getGifFromJson(resBody)
}

func buildRollCommandHandler(heroData *[]*hero, tenorKey string) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var responseContent string
		// parse option data into a useable form
		optionMap := getOptionsFromDiscordInteraction(i)
		roleValue, ok := optionMap["role"].(string)
		if !ok {
			roleValue = "all"
		}
		isStadium, ok := optionMap["stadium"].(bool)
		if !ok {
			isStadium = false
		}
		validHeroes := make([]*hero, 0)
		isAllRoles := roleValue == "all"
		for _, heroDatum := range *heroData {
			roleMatch := isAllRoles || heroDatum.Role == roleValue
			stadiumMatch := !isStadium || heroDatum.Stadium
			if roleMatch && stadiumMatch {
				fmt.Printf("found match %v\n", heroDatum.Name)
				validHeroes = append(validHeroes, heroDatum)
			}
		}
		// valid role provided but no hero data found
		// this should never happen given the data, but is included for completeness
		if len(validHeroes) == 0 {
			responseContent = fmt.Sprintf("no heroes found for role %v", roleValue)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: responseContent,
				},
			})
			return
		}
		heroChoice := validHeroes[rand.Intn(len(validHeroes))]
		gifUrl, err := getHeroGif(heroChoice.Name, tenorKey)
		var embeds []*discordgo.MessageEmbed
		if err != nil {
			embeds = []*discordgo.MessageEmbed{}
		} else {
			embeds = []*discordgo.MessageEmbed{
				{
					Type: discordgo.EmbedTypeGifv,
					Image: &discordgo.MessageEmbedImage{
						URL: gifUrl,
					},
				},
			}
		}

		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: heroChoice.Name,
				Embeds:  embeds,
			},
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func main() {
	discordConfigPath, heroDataConfigPath := readAppFlags()
	if discordConfigPath == "" || heroDataConfigPath == "" {
		log.Fatalf("Error: missing required file path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	appConfigValues, err := readConfigValues[discordConfig](discordConfigPath)
	if err != nil {
		log.Fatalf("unable to read Discord connection configuration: %v", err)
		os.Exit(1)
	}

	heroData, err := readConfigValues[heroConfig](heroDataConfigPath)
	if err != nil {
		log.Fatalf("unable to read hero data: %v", err)
		os.Exit(1)
	}
	heroRollCommandHandler := buildRollCommandHandler(heroData.Heroes, appConfigValues.TenorKey)
	s, err := initializeDiscordSession(appConfigValues)
	if err != nil {
		log.Fatalf("unable to initialize discord session: %v", err)
	}

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	log.Println("opening session")
	err = s.Open()
	if err != nil {
		log.Fatalln("unable to open session")
		os.Exit(4)
	}

	defer s.Close()

	log.Println("adding commands")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, appConfigValues.GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
			return
		}
		registeredCommands[i] = cmd
	}
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		heroRollCommandHandler(s, i)
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	for _, command := range registeredCommands {
		s.ApplicationCommandDelete(command.ApplicationID, appConfigValues.GuildID, command.ID)
	}
}
