package commands

import (
	"time"

	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// Initialize initializes all commands
func Initialize(router *dgc.Router, session *discordgo.Session) {
	// Define the default rate limiter
	rateLimiter := dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed("Hey! Don't spam!"))
	})

	// Register the default help command
	router.RegisterDefaultHelpCommand(session, rateLimiter)

	// Register the random command
	router.RegisterCmd(&dgc.Command{
		Name:        "random",
		Description: "Generates a random boolean, number, string or choice",
		Usage:       "random <bool | number <interval> | string <int: length> | choice <options...>>",
		Example:     "random bool",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "bool",
				Aliases:     []string{"b"},
				Description: "Generates a random boolean",
				Usage:       "random bool",
				Example:     "random bool",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     RandomBool,
			},
			&dgc.Command{
				Name:        "number",
				Aliases:     []string{"n"},
				Description: "Generates a random number",
				Usage:       "random number <interval>",
				Example:     "random number [0, 200]",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     RandomNumber,
			},
			&dgc.Command{
				Name:        "string",
				Aliases:     []string{"s"},
				Description: "Generates a random string",
				Usage:       "random string <int: length>",
				Example:     "random string 20",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     RandomString,
			},
			&dgc.Command{
				Name:        "choice",
				Aliases:     []string{"c"},
				Description: "Chooses a random element of the given options",
				Usage:       "random choice <options...>",
				Example:     "random choice I am cool lol",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     RandomChoice,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     Random,
	})

	// Register the hash command
	hashingRateLimiter := dgc.NewRateLimiter(30*time.Second, 3*time.Second, func(ctx *dgc.Ctx) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed("You need to wait at least thirty seconds between two hash calculations."))
	})
	router.RegisterCmd(&dgc.Command{
		Name:        "hash",
		Description: "Hashes the given string using the specified algorithm",
		Usage:       "hash <md5> <string>",
		Example:     "hash md5 hello",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "md5",
				Description: "Hashes the given string using the md5 algorithm",
				Usage:       "hash md5 <string>",
				Example:     "hash md5 hello",
				IgnoreCase:  true,
				RateLimiter: hashingRateLimiter,
				Handler:     HashMD5,
			},
		},
		RateLimiter: hashingRateLimiter,
		Handler:     Hash,
	})

	// Register the math command
	router.RegisterCmd(&dgc.Command{
		Name:        "math",
		Description: "Evaluates the given mathematical expression",
		Usage:       "math <codeblock>",
		Example:     "math 10^3",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     Math,
	})

	// Register the latex command
	router.RegisterCmd(&dgc.Command{
		Name:        "latex",
		Description: "Renders the given LaTeX expression",
		Usage:       "latex <codeblock>",
		Example:     "latex 10^3",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     Latex,
	})

	// Register the hastebin feature
	session.AddHandler(HastebinMessageCreateListener)
	session.AddHandler(HastebinReactionAddListener)

	// Register the info command
	router.RegisterCmd(&dgc.Command{
		Name:        "info",
		Description: "Displays some useful information about the bot",
		Usage:       "info",
		Example:     "info",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     Info,
	})

	// Register the stats command
	router.RegisterCmd(&dgc.Command{
		Name:        "stats",
		Description: "Displays some general statistics about the bot",
		Usage:       "stats",
		Example:     "stats",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     Stats,
	})

	// Register the request command
	router.RegisterCmd(&dgc.Command{
		Name:        "request",
		Description: "Creates a feature request",
		Usage:       "request <string>",
		Example:     "request My cool new feature",
		IgnoreCase:  true,
		RateLimiter: dgc.NewRateLimiter(1*time.Hour, 1*time.Minute, func(ctx *dgc.Ctx) {
			ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed("You need to wait at least one hour between two feature requests."))
		}),
		Handler: Request,
	})
	session.AddHandler(RequestReactionListener)

	// Register the settings command
	router.RegisterCmd(&dgc.Command{
		Name:        "settings",
		Description: "Displays the current guild settings or changes them",
		Usage:       "settings [toggleChannelRestriction | toggleCommandChannel <channel mention> | toggleHastebinIntegration]",
		Example:     "settings toggleChannelRestriction",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "toggleChannelRestriction",
				Aliases:     []string{"tcr"},
				Description: "Toggles the current command channel restriction status",
				Usage:       "settings toggleChannelRestriction",
				Example:     "settings toggleChannelRestriction",
				Flags: []string{
					"guildAdminOnly",
				},
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     SettingsToggleChannelRestriction,
			},
			&dgc.Command{
				Name:        "toggleCommandChannel",
				Aliases:     []string{"tcc"},
				Description: "Toggles the command channel status of the mentioned channel",
				Usage:       "settings toggleCommandChannel <channel mention>",
				Example:     "settings toggleCommandChannel #my-channel",
				Flags: []string{
					"guildAdminOnly",
				},
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     SettingsToggleCommandChannel,
			},
			&dgc.Command{
				Name:        "toggleHastebinIntegration",
				Aliases:     []string{"thi"},
				Description: "Toggles the current hastebin integration status",
				Usage:       "settings toggleHastebinIntegration",
				Example:     "settings toggleHastebinIntegration",
				Flags: []string{
					"guildAdminOnly",
				},
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     SettingsToggleHastebinIntegration,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     Settings,
	})

	// Register the say command
	router.RegisterCmd(&dgc.Command{
		Name:        "say",
		Description: "[Bot Admin only] Makes the bot say something",
		Usage:       "say <string>",
		Example:     "say Hello, world!",
		Flags: []string{
			"botAdminOnly",
		},
		IgnoreCase: true,
		Handler:    Say,
	})

	// Register the blacklist command
	router.RegisterCmd(&dgc.Command{
		Name:        "blacklist",
		Description: "[Bot Admin only] Adds/Removes a user to/from the command blacklist",
		Usage:       "blacklist <user mention>",
		Example:     "blacklist @Erik",
		Flags: []string{
			"botAdminOnly",
		},
		IgnoreCase: true,
		Handler:    Blacklist,
	})

	// Register the debug command
	router.RegisterCmd(&dgc.Command{
		Name:        "debug",
		Description: "[Bot Admin only] Executes the given code at runtime",
		Usage:       "debug <codeblock>",
		Example:     "debug fmt.Println(\"Hello, world!\")",
		Flags: []string{
			"botAdminOnly",
		},
		IgnoreCase: true,
		Handler:    Debug,
	})
}
