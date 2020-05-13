package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Lukaesebrot/asterisk/middlewares"

	"github.com/Lukaesebrot/asterisk/commands"
	"github.com/Lukaesebrot/asterisk/static"
	"github.com/Lukaesebrot/dgc"

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
	session, err := discordgo.New("Bot " + config.CurrentConfig.BotToken)
	if err != nil {
		panic(err)
	}
	err = session.Open()
	if err != nil {
		panic(err)
	}
	static.Self = session.State.User
	log.Println("Successfully established the Discord connection.")

	// Initialize the command system
	log.Println("Initializing the command system...")
	router := dgc.Create(&dgc.Router{
		Prefixes: []string{
			"$",
			"<@!" + static.Self.ID + ">",
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
	router.RegisterDefaultHelpCommand(session)
	router.RegisterCmd(&dgc.Command{
		Name:        "info",
		Description: "Displays some useful information about the bot",
		Usage:       "info",
		IgnoreCase:  true,
		Handler:     commands.Info,
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "stats",
		Description: "Displays some general statistics about the bot",
		Usage:       "stats",
		IgnoreCase:  true,
		Handler:     commands.Stats,
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "random",
		Description: "Generates a random bool, number, string or choice",
		Usage:       "random <bool | number <interval> | string <int: length> | choice <options...>>",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "bool",
				Aliases:     []string{"b"},
				Description: "Generates a random boolean",
				IgnoreCase:  true,
				Handler:     commands.RandomBool,
			},
			&dgc.Command{
				Name:        "number",
				Aliases:     []string{"n"},
				Description: "Generates a random number",
				IgnoreCase:  true,
				Handler:     commands.RandomNumber,
			},
			&dgc.Command{
				Name:        "string",
				Aliases:     []string{"s"},
				Description: "Generates a random string",
				IgnoreCase:  true,
				Handler:     commands.RandomString,
			},
			&dgc.Command{
				Name:        "choice",
				Aliases:     []string{"c"},
				Description: "Generates a random choice",
				IgnoreCase:  true,
				Handler:     commands.RandomChoice,
			},
		},
		Handler: commands.Random,
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "md5",
		Description: "Hashes the given arguments using md5",
		Usage:       "md5 <string>",
		IgnoreCase:  true,
		Handler:     commands.MD5,
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "say",
		Description: "[Bot Admin only] Makes me say something",
		Usage:       "say <string>",
		Flags: []string{
			"botAdminOnly",
		},
		IgnoreCase: true,
		Handler:    commands.Say,
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "blacklist",
		Description: "[Bot Admin only] Adds/Removes a user to/from the command blacklist",
		Usage:       "blacklist <user mention>",
		Flags: []string{
			"botAdminOnly",
		},
		IgnoreCase: true,
		Handler:    commands.Blacklist,
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "debug",
		Description: "[Bot Admin only] Executes the given string at runtime",
		Usage:       "debug <codeblock>",
		Flags: []string{
			"botAdminOnly",
		},
		IgnoreCase: true,
		Handler:    commands.Debug,
	})
	log.Println("Successfully registered commands.")

	// Register middlewares
	log.Println("Registering middlewares...")
	router.AddMiddleware("*", middlewares.CheckCommandBlacklist)
	router.AddMiddleware("*", middlewares.InjectGuildConfig)
	router.AddMiddleware("*", middlewares.CheckCommandChannel)
	router.AddMiddleware("botAdminOnly", middlewares.CheckBotAdmin)
	log.Println("Successfully registered middlewares.")

	// Handle incoming console commands
	log.Println("Waiting for console commands. Type 'help' for help.")
	reader := bufio.NewReader(os.Stdin)
	concommands.Handle(reader, session)
}
