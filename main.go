package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

type conf struct {
	GuildID  string `yaml:"guildId"`
	BotToken string `yaml:"botToken"`
}

var (
	s        *discordgo.Session
	filePath string
	config   *conf
	initErr  error
)

func init() {
	const (
		defaultValue = ""
		usage        = "path to config file"
	)
	flag.StringVar(&filePath, "filename", defaultValue, usage)
	flag.StringVar(&filePath, "f", defaultValue, usage)
	flag.Parse()
}

func init() {
	if filePath == "" {
		log.Fatalf("Error: must provide a path to config file")
		flag.PrintDefaults()
		os.Exit(1)
	}
	config, initErr = getConfigValues(filePath)
	if initErr != nil {
		log.Fatalf("Error: unable to parse config YAML file %v", initErr)
		flag.PrintDefaults()
		os.Exit(2)
	}
}

func init() {
	s, initErr = discordgo.New("Bot " + config.BotToken)
	if initErr != nil {
		log.Fatalf("Error: unable to create session %v", initErr)
		os.Exit(3)
	}
	fmt.Println(s)
}

func getConfigValues(filePath string) (*conf, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config *conf
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func main() {
	s.Close()
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
