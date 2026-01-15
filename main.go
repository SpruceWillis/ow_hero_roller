package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

type heroConfig struct {
	Heroes *[]*Hero `yaml:"heroes"`
}

// config for each hero
type Hero struct {
	Name    string `yaml:"name"`
	Role    string `yaml:"role"`
	Stadium bool   `yaml:"stadium"`
}

var (
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

const (
	tenorSearchUrl = "https://tenor.googleapis.com/v2/search"
	TENOR_API_KEY  = "TENOR_API_KEY"
	BOT_TOKEN      = "BOT_TOKEN"
	PUBLIC_KEY     = "PUBLIC_KEY"
)

// establish flags and read config values
func readAppFlags() string {
	const (
		defaultValue    = ""
		discordOptUsage = "path to Discord config file"
		heroOptUsage    = "path to hero data file"
	)

	var heroConfigPath string
	flag.StringVar(&heroConfigPath, "data-file-name", defaultValue, heroOptUsage)
	flag.StringVar(&heroConfigPath, "d", defaultValue, heroOptUsage)
	flag.Parse()
	return heroConfigPath
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

func getValidHeroes(heroData *[]*Hero, role string, isStadium bool) []*Hero {
	validHeroes := []*Hero{}
	isAllRoles := role == "all"
	for _, heroDatum := range *heroData {
		roleMatch := isAllRoles || heroDatum.Role == role
		stadiumMatch := !isStadium || heroDatum.Stadium
		if roleMatch && stadiumMatch {
			validHeroes = append(validHeroes, heroDatum)
		}
	}
	return validHeroes
}

func handleRollCommand(heroData *[]*Hero, tenorKey string, i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
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

	validHeroes := getValidHeroes(heroData, roleValue, isStadium)

	if len(validHeroes) == 0 {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("no heroes found for role %v", roleValue),
			},
		}
	}

	heroChoice := validHeroes[rand.Intn(len(validHeroes))]
	gifUrl, err := getHeroGif(heroChoice.Name, tenorKey)
	var embeds []*discordgo.MessageEmbed
	if err == nil {
		embeds = []*discordgo.MessageEmbed{
			{
				Type: discordgo.EmbedTypeGifv,
				Image: &discordgo.MessageEmbedImage{
					URL: gifUrl,
				},
			},
		}
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: heroChoice.Name,
			Embeds:  embeds,
		},
	}
}

func createCommandIfNeeded(botToken string) {
	// Initialize Discord session for REST API calls
	s, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}

	// Get the bot's application ID
	u, err := s.User("@me")
	if err != nil {
		log.Fatalf("unable to retrieve bot user: %v", err)
	}
	appID := u.ID

	// Check if the command exists
	existingCommands, err := s.ApplicationCommands(appID, "")
	if err != nil {
		log.Printf("unable to retrieve commands: %v", err)
	}

	commandExists := false
	for _, cmd := range existingCommands {
		if cmd.Name == "hero_roll" {
			commandExists = true
			log.Println("command 'hero_roll' already exists, skipping creation")
			break
		}
	}

	if !commandExists {
		log.Println("command 'hero_roll' not found, creating it...")
		_, err = s.ApplicationCommandCreate(appID, "", commands[0])
		if err != nil {
			log.Fatalf("Cannot create 'hero_roll' command: %v", err)
		}
		log.Println("command 'hero_roll' created successfully")
	}

	s.Close()

}

func verifySignature(r *http.Request, key string) bool {
	signature := r.Header.Get("X-Signature-Ed25519")
	timestamp := r.Header.Get("X-Signature-Timestamp")

	if signature == "" || timestamp == "" {
		return false
	}

	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return false
	}

	return discordgo.VerifyInteraction(r, ed25519.PublicKey(keyBytes))
}

func main() {
	heroDataConfigPath := readAppFlags()
	if heroDataConfigPath == "" {
		log.Println("Error: missing required file path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	heroData, err := readConfigValues[heroConfig](heroDataConfigPath)
	if err != nil {
		log.Fatalf("unable to read hero data: %v", err)
		os.Exit(1)
	}

	tenorKey := strings.TrimSpace(os.Getenv(TENOR_API_KEY))
	if tenorKey == "" {
		log.Fatalf("unable to read tenor API key from environment variable %v", TENOR_API_KEY)
	}

	publicKey := strings.TrimSpace(os.Getenv(PUBLIC_KEY))
	if publicKey == "" {
		log.Fatalf("unable to read public key from environment variable %v", PUBLIC_KEY)
	}

	botToken := os.Getenv(BOT_TOKEN)
	if botToken == "" {
		log.Fatalf("unable to read bot token from environment variable %v", BOT_TOKEN)
	}

	createCommandIfNeeded(botToken)

	http.HandleFunc("/interactions", func(w http.ResponseWriter, r *http.Request) {
		if !verifySignature(r, publicKey) {
			http.Error(w, "invalid signature", http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error reading request body", http.StatusInternalServerError)
			return
		}

		var interaction discordgo.InteractionCreate
		if err := json.Unmarshal(body, &interaction); err != nil {
			http.Error(w, "error parsing json", http.StatusBadRequest)
			return
		}

		switch interaction.Type {
		case discordgo.InteractionPing:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(discordgo.InteractionResponse{
				Type: discordgo.InteractionResponsePong,
			})
		case discordgo.InteractionApplicationCommand:
			if interaction.ApplicationCommandData().Name == "hero_roll" {
				response := handleRollCommand(heroData.Heroes, tenorKey, &interaction)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			}
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK\n")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("starting server on port %v", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
