package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"slices"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

// config for the overall application
type discordConfig struct {
	GuildID  string `yaml:"guildId"`
	BotToken string `yaml:"botToken"`
}

type heroConfig struct {
	Heroes *[]hero `yaml:heroes`
}

// config for each hero
type hero struct {
	Name      string `yaml:"name"`
	Role      string `yaml:"role"`
	ImageLink string `yaml:"imageLink"`
}

var (
	validRoles = []string{
		"all", "tank", "dps", "support",
	}
	commands = []*discordgo.ApplicationCommand{
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

func mapByRole(heroData *heroConfig) *map[string][]*hero {
	result := map[string][]*hero{
		"all": make([]*hero, 0),
	}
	for _, heroDatum := range *heroData.Heroes {
		if !slices.Contains(validRoles, heroDatum.Role) {
			log.Fatalf("error: found invalid role %v at hero %v", heroDatum.Role, heroDatum.Name)
			os.Exit(1)
		}
		result["all"] = append(result["all"], &heroDatum)
		currentRoleHeroes, ok := result[heroDatum.Role]
		var newRoleValue []*hero
		if ok {
			newRoleValue = append(currentRoleHeroes, &heroDatum)
		} else {
			newRoleValue = []*hero{
				&heroDatum,
			}
		}
		result[heroDatum.Role] = newRoleValue
	}

	return &result
}

// connect to discord
func initializeDiscordSession(config *discordConfig) (*discordgo.Session, error) {
	fmt.Println("initializing discord connection")
	return discordgo.New("Bot " + config.BotToken)
}

func getOptionsFromDiscordInteraction(i *discordgo.InteractionCreate) map[string]string {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]string, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt.StringValue()
	}
	return optionMap
}

func getImageFileType(imageUrl string) string {
	return filepath.Ext(imageUrl)
}

func toProtocol(protocol string, original string) (string, error) {
	u, err := url.Parse(original)
	if err != nil {
		return "", err
	}
	u.Scheme = protocol
	return u.String(), nil
}

func buildRollCommandHandler(heroesByRole *map[string][]*hero) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var responseContent string
		// parse option data into a useable form
		heroMap := *heroesByRole
		optionMap := getOptionsFromDiscordInteraction(i)
		roleValue, ok := optionMap["role"]
		// default value
		if !ok {
			roleValue = "all"
		}
		// TODO add a thumbnail/profile to the bot via web portal
		validHeroes, ok := heroMap[roleValue]
		// valid role provided but no hero data found
		// this should never happen given the data, but is included for completeness
		if !ok || len(validHeroes) == 0 {
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
		responseContent = heroChoice.Name

		if heroChoice.ImageLink == "" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: responseContent,
				},
			})
			return
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: responseContent,
				// Embeds:  embeds,
			},
		})

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
	mappedHeroData := mapByRole(heroData)
	heroRollCommandHandler := buildRollCommandHandler(mappedHeroData)
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
