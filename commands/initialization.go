package commands

import (
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// Initialize initializes all commands
func Initialize(router *dgc.Router, session *discordgo.Session) {
	// Register the default help command
	router.RegisterDefaultHelpCommand(session)

	// Register the info command
	router.RegisterCmd(&dgc.Command{
		Name:        "info",
		Description: "Displays some useful information about the bot",
		Usage:       "info",
		IgnoreCase:  true,
		Handler:     Info,
	})

	// Register the stats command
	router.RegisterCmd(&dgc.Command{
		Name:        "stats",
		Description: "Displays some general statistics about the bot",
		Usage:       "stats",
		IgnoreCase:  true,
		Handler:     Stats,
	})

	// Register the request command
	router.RegisterCmd(&dgc.Command{
		Name:        "request",
		Description: "Creates a feature request",
		Usage:       "request <feature request>",
		IgnoreCase:  true,
		Handler:     Request,
	})
	session.AddHandler(RequestReactionListener)

	// Register the random command
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
				Handler:     RandomBool,
			},
			&dgc.Command{
				Name:        "number",
				Aliases:     []string{"n"},
				Description: "Generates a random number",
				IgnoreCase:  true,
				Handler:     RandomNumber,
			},
			&dgc.Command{
				Name:        "string",
				Aliases:     []string{"s"},
				Description: "Generates a random string",
				IgnoreCase:  true,
				Handler:     RandomString,
			},
			&dgc.Command{
				Name:        "choice",
				Aliases:     []string{"c"},
				Description: "Generates a random choice",
				IgnoreCase:  true,
				Handler:     RandomChoice,
			},
		},
		Handler: Random,
	})

	// Register the md5 command
	router.RegisterCmd(&dgc.Command{
		Name:        "md5",
		Description: "Hashes the given arguments using md5",
		Usage:       "md5 <string>",
		IgnoreCase:  true,
		Handler:     MD5,
	})

	// Register the say command
	router.RegisterCmd(&dgc.Command{
		Name:        "say",
		Description: "[Bot Admin only] Makes me say something",
		Usage:       "say <string>",
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
		Flags: []string{
			"botAdminOnly",
		},
		IgnoreCase: true,
		Handler:    Blacklist,
	})

	// Register the debug command
	router.RegisterCmd(&dgc.Command{
		Name:        "debug",
		Description: "[Bot Admin only] Executes the given string at runtime",
		Usage:       "debug <codeblock>",
		Flags: []string{
			"botAdminOnly",
		},
		IgnoreCase: true,
		Handler:    Debug,
	})
}
