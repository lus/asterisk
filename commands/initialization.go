package commands

import (
	"time"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// Initialize initializes all commands
func Initialize(router *dgc.Router, session *discordgo.Session) {
	// Define the default rate limiter
	rateLimiter := dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("Hey! Don't spam!"))
	})

	// Register the default help command
	router.RegisterDefaultHelpCommand(session, rateLimiter)

	// Register the settings command
	router.RegisterCmd(&dgc.Command{
		Name:        "settings",
		Description: "Displays the current guild settings or changes them",
		Usage:       "settings [commandChannel <channel id or mention>]",
		Example:     "settings commandChannel 573374556409032800",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "commandChannel",
				Aliases:     []string{"cc"},
				Description: "Toggles the command channel status of the given channel",
				Usage:       "settings commandChannel <channel id or mention>",
				Example:     "settings commandChannel 573374556409032800",
				Flags: []string{
					"guild_admin",
					"ignore_command_channel",
				},
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     SettingsCommandChannel,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     Settings,
	})

	// Register the reminder command
	router.RegisterCmd(&dgc.Command{
		Name:        "reminder",
		Description: "Lists the current reminders",
		Usage:       "reminder [create <duration> <message> | delete <id>]",
		Example:     "reminder",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			&dgc.Command{
				Name:        "create",
				Aliases:     []string{"c"},
				Description: "Creates a reminder",
				Usage:       "reminder create <duration> <message>",
				Example:     "reminder create 1h That's my reminder!",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     ReminderCreate,
			},
			&dgc.Command{
				Name:        "delete",
				Aliases:     []string{"d", "rm"},
				Description: "Deletes a reminder",
				Usage:       "reminder delete <id>",
				Example:     "reminder delete 1",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     ReminderDelete,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     Reminder,
	})

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
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("You need to wait at least thirty seconds between two hash calculations."))
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

	// Register the bug command
	router.RegisterCmd(&dgc.Command{
		Name:        "bug",
		Description: "Reports a bug",
		Usage:       "bug <string>",
		Example:     "bug The bot spams",
		IgnoreCase:  true,
		RateLimiter: dgc.NewRateLimiter(1*time.Hour, 1*time.Minute, func(ctx *dgc.Ctx) {
			ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("You need to wait at least one hour between two bug reports."))
		}),
		Handler: Bug,
	})
	session.AddHandler(BugReactionListener)

	// Register the request command
	router.RegisterCmd(&dgc.Command{
		Name:        "request",
		Description: "Creates a feature request",
		Usage:       "request <string>",
		Example:     "request My cool new feature",
		IgnoreCase:  true,
		RateLimiter: dgc.NewRateLimiter(1*time.Hour, 1*time.Minute, func(ctx *dgc.Ctx) {
			ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("You need to wait at least one hour between two feature requests."))
		}),
		Handler: Request,
	})
	session.AddHandler(RequestReactionListener)

	// Register the say command
	router.RegisterCmd(&dgc.Command{
		Name:        "say",
		Description: "[Bot Admin only] Makes the bot say something",
		Usage:       "say <string>",
		Example:     "say Hello, world!",
		Flags: []string{
			"bot_admin",
		},
		IgnoreCase: true,
		Handler:    Say,
	})

	// Register the debug command
	router.RegisterCmd(&dgc.Command{
		Name:        "debug",
		Description: "[Bot Admin only] Executes the given code at runtime",
		Usage:       "debug <codeblock>",
		Example:     "debug fmt.Println(\"Hello, world!\")",
		Flags: []string{
			"bot_admin",
		},
		IgnoreCase: true,
		Handler:    Debug,
	})
}
