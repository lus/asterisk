package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Lukaesebrot/asterisk/cmdparser"
	"github.com/Lukaesebrot/asterisk/commands"
	"github.com/Lukaesebrot/asterisk/utils"

	"github.com/Lukaesebrot/asterisk/concommands"
	"github.com/Lukaesebrot/asterisk/config"
	"github.com/Lukaesebrot/asterisk/database"
	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("Starting this Asterisk instance...")

	// Initialize the configuration
	log.Println("Loading the bot configuration...")
	err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully loaded the bot configuration.")

	// Connect to the MongoDB host
	log.Println("Connecting to the specified MongoDB server...")
	err = database.Connect()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to the specified MongoDB server.")

	// Initialize the Discord session
	log.Println("Establishing the Discord connection...")
	session, err := discordgo.New("Bot " + config.CurrentConfig.Token)
	if err != nil {
		panic(err)
	}
	self, err := session.User("@me")
	if err != nil {
		panic(err)
	}
	log.Println("Successfully established the Discord connection.")

	// Initialize the command system
	log.Println("Initializing the command system...")
	commandSystem := &cmdparser.CommandSystem{
		BotUser: self,
		Prefixes: []string{
			"<@" + self.ID + ">",
			"$",
			"as!",
			"你",
		},
		Commands: map[string]*cmdparser.Command{
			"info": &cmdparser.Command{
				Handler: commands.Info(self),
			},
		},
		PingHandler: func(session *discordgo.Session, event *discordgo.MessageCreate) {
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateBotInfoEmbed(self))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
		},
	}
	session.AddHandler(commandSystem.Handler())
	log.Println("Successfully initialized the command system.")

	// Handle incoming console commands
	log.Println("Waiting for console commands. Type 'help' for help.")
	reader := bufio.NewReader(os.Stdin)
	concommands.Handle(reader, session)
}
