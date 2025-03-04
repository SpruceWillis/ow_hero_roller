package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

// config for the overall application
type secretsConfig struct {
	GuildID  string `yaml:"guildId"`
	BotToken string `yaml:"botToken"`
}

type heroConfig struct {
	Heroes *[]heroData `yaml:heroes`
}

// config for each hero
type heroData struct {
	Name      string `yaml:"name"`
	Role      string `yaml:"role"`
	ImageLink string `yaml:"imageLink"`
}

var (
	s              *discordgo.Session
	appFilePath    string
	dataFilePath   string
	appConfig      *secretsConfig
	heroDataConfig *heroConfig
	initErr        error

	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "roll",
			Description: "roll a random hero",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "hero_role",
					Description: "Hero roles to consider. Valid values are 'all' (default) 'dps' 'support' 'tank'",
					Required:    false,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"roll": func(_ *discordgo.Session, i *discordgo.InteractionCreate) {
			// parse option data into a useable form
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}
			margs := make([]interface{}, 0, len(options))
			msgformat := "You learned how to use command options! " +
				"Take a look at the value(s) you entered:\n"

			if option, ok := optionMap["hero_role"]; ok {
				margs = append(margs, option.StringValue())
				msgformat += "> hero role option %s\n"
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
		},
	}
)

// establish flags and read config values
func init() {
	const (
		defaultValue = ""
		usage        = "path to config file"
	)
	flag.StringVar(&appFilePath, "file", defaultValue, usage)
	flag.StringVar(&appFilePath, "f", defaultValue, usage)
	flag.StringVar(&dataFilePath, "data-file-name", defaultValue, usage)
	flag.StringVar(&dataFilePath, "d", defaultValue, usage)
	flag.Parse()
}

// read config to get file path and get config values
// TODO: refactor this to make config parsing more generic
func init() {
	if appFilePath == "" {
		log.Fatalf("Error: must provide a path to application config file")
		flag.PrintDefaults()
		os.Exit(1)
	}
	appConfig, initErr = getAppConfigValues(appFilePath)
	if initErr != nil {
		log.Fatalf("Error: unable to parse application config YAML file %v", initErr)
		flag.PrintDefaults()
		os.Exit(2)
	}
}

// TODO refactor this to make parsing more generic
func getAppConfigValues(filePath string) (*secretsConfig, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config *secretsConfig
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// connect to discord
func init() {
	fmt.Println("initializing discord connection")
	s, initErr = discordgo.New("Bot " + appConfig.BotToken)
	if initErr != nil {
		log.Fatalf("Error: unable to create session %v", initErr)
		os.Exit(3)
	}
}

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	log.Println("opening session")
	err := s.Open()
	if err != nil {
		log.Fatalln("unable to open session")
		os.Exit(4)
	}

	defer s.Close()

	log.Println("adding commands")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, appConfig.GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	for _, command := range registeredCommands {
		s.ApplicationCommandDelete(command.ApplicationID, appConfig.GuildID, command.ID)
	}
}

// func slashCommands() []*discordgo.ApplicationCommand {
// 	commands := []*discordgo.ApplicationCommand{
// 		{
// 			Name:       "roll-hero",
// 			Description: "Roll an Overwatch hero",
// 		},
// 	}
// 	return commands
// }
