package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Lukaesebrot/asterisk/middlewares"

	"github.com/Lukaesebrot/asterisk/commands"
	"github.com/Lukaesebrot/asterisk/static"
	"github.com/Lukaesebrot/dgc"

	"github.com/Lukaesebrot/asterisk/config"
	"github.com/Lukaesebrot/asterisk/database"
	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("Starting this Asterisk instance...")

	// Initialize the configuration
	log.Println("Loading the bot configuration...")
	config.Load()
	log.Println("Successfully loaded the bot configuration.")

	// Connect to the MongoDB host
	log.Println("Connecting to the specified MongoDB server...")
	err := database.Connect()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to the specified MongoDB server.")

	// Initialize the Discord session
	log.Println("Establishing the Discord connection...")
	session, err := discordgo.New("Bot " + config.CurrentConfig.BotToken)
	if err != nil {
		panic(err)
	}
	session.AddHandler(func(session *discordgo.Session, event *discordgo.Ready) {
		session.UpdateListeningStatus("$help")
	})
	err = session.Open()
	if err != nil {
		panic(err)
	}
	static.Self = session.State.User
	static.StartTime = time.Now()
	log.Println("Successfully established the Discord connection.")

	// Initialize the command system
	log.Println("Initializing the command system...")
	router := dgc.Create(&dgc.Router{
		Prefixes: []string{
			"$",
			"<@!" + static.Self.ID + ">",
			"<@" + static.Self.ID + ">",
			"as!",
			"你",
		},
		IgnorePrefixCase: true,
		BotsAllowed:      false,
		PingHandler:      commands.Info,
	})
	router.Initialize(session)
	log.Println("Successfully initialized the command system.")

	// Register commands
	log.Println("Registering commands...")
	commands.Initialize(router, session)
	log.Println("Successfully registered commands.")

	// Register middlewares
	log.Println("Registering middlewares...")
	router.AddMiddleware("*", middlewares.CheckCommandBlacklist)
	router.AddMiddleware("*", middlewares.InjectGuildConfig)
	router.AddMiddleware("*", middlewares.CheckCommandChannel)
	router.AddMiddleware("botAdminOnly", middlewares.CheckBotAdmin)
	router.AddMiddleware("guildAdminOnly", middlewares.CheckGuildAdmin)
	log.Println("Successfully registered middlewares.")

	// Wait for the program to exit
	log.Println("Successfully started this Asterisk instance.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Gracefully close the Discord session
	log.Println("Stopping this Asterisk instance...")
	session.Close()
	database.CurrentClient.Disconnect(context.TODO())
}
