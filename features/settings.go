package features

import (
	"strconv"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/guilds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// initializeSettingsFeature initializes the settings feature
func initializeSettingsFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'settings' command
	router.RegisterCmd(&dgc.Command{
		Name:        "settings",
		Description: "Displays the current guild settings or changes them",
		Usage:       "settings [commandChannel <channel mention>]",
		Example:     "settings commandChannel #my-channel",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			{
				Name:        "commandChannel",
				Description: "Toggles the command channel status for the mentioned channel",
				Usage:       "settings commandChannel <channel mention or ID>",
				Example:     "settings commandChannel #my-channel",
				Flags: []string{
					"guild_admin",
					"ignore_command_channel",
				},
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     settingsCommandChannelCommand,
			},
			{
				Name:        "starboard",
				Description: "Sets the starboard channel, the minimum star amount or disables the feature",
				Usage:       "settings starboard <disable | channel <channel mention or ID> | minimum <number>>",
				Example:     "settings starboard channel #my-channel",
				Flags: []string{
					"guild_admin",
				},
				IgnoreCase: true,
				SubCommands: []*dgc.Command{
					{
						Name:        "disable",
						Description: "Disables the starboard feature",
						Usage:       "settings starboard disable",
						Example:     "settings starboard disable",
						Flags: []string{
							"guild_admin",
						},
						IgnoreCase:  true,
						RateLimiter: rateLimiter,
						Handler:     settingsStarboardDisableCommand,
					},
					{
						Name:        "channel",
						Description: "Sets the starboard channel",
						Usage:       "settings starboard channel <channel mention or ID>",
						Example:     "settings starboard channel #my-channel",
						Flags: []string{
							"guild_admin",
						},
						IgnoreCase:  true,
						RateLimiter: rateLimiter,
						Handler:     settingsStarboardChannelCommand,
					},
					{
						Name:        "minimum",
						Description: "Sets the minimum star amount",
						Usage:       "settings starboard minimum <number>",
						Example:     "settings starboard minimum 5",
						Flags: []string{
							"guild_admin",
						},
						IgnoreCase:  true,
						RateLimiter: rateLimiter,
						Handler:     settingsStarboardMinimumCommand,
					},
				},
				RateLimiter: rateLimiter,
				Handler: func(ctx *dgc.Ctx) {
					ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage(ctx.Command.Usage))
				},
			},
		},
		RateLimiter: rateLimiter,
		Handler:     settingsCommand,
	})
}

// settingsCommand handles the 'settings' command
func settingsCommand(ctx *dgc.Ctx) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Settings(ctx.CustomObjects.MustGet("guild").(*guilds.Guild)))
}

// settingsCommandChannelCommand handles the 'settings commandChannel' command
func settingsCommandChannelCommand(ctx *dgc.Ctx) {
	// Validate the command channel ID
	channelID := ctx.Arguments.AsSingle().AsChannelMentionID()
	if channelID == "" {
		channelID = ctx.Arguments.Raw()
	}

	// Validate the channel itself
	_, err := ctx.Session.Channel(channelID)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("The given channel couldn't be found."))
		return
	}

	// Retrieve the guild object
	guild := ctx.CustomObjects.MustGet("guild").(*guilds.Guild)

	// Add/remove the channel ID to/from the command channel list
	commandChannels := guild.Settings.CommandChannels
	contains := utils.StringArrayContains(commandChannels, channelID)
	if contains {
		newCommandChannels := make([]string, len(commandChannels)-1)
		counter := 0
		for _, value := range commandChannels {
			if value != channelID {
				newCommandChannels[counter] = value
				counter++
			}
		}
		guild.Settings.CommandChannels = newCommandChannels
	} else {
		guild.Settings.CommandChannels = append(guild.Settings.CommandChannels, channelID)
	}

	// Respond with a success message
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Success("The command channel status for the mentioned channel has been "+utils.PrettifyBool(!contains)+"."))
}

// settingsStarboardDisableCommand handles the 'settings starboard disable' command
func settingsStarboardDisableCommand(ctx *dgc.Ctx) {
	// Retrieve the guild object
	guild := ctx.CustomObjects.MustGet("guild").(*guilds.Guild)

	// Disable the starboard feature
	guild.Settings.Starboard.Channel = ""
	err := guild.Update()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}

	// Respond with a success message
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Success("The starboard feature got disabled."))
}

// settingsStarboardChannelCommand handles the 'settings starboard channel' command
func settingsStarboardChannelCommand(ctx *dgc.Ctx) {
	// Validate the starboard channel ID
	channelID := ctx.Arguments.AsSingle().AsChannelMentionID()
	if channelID == "" {
		channelID = ctx.Arguments.Raw()
	}

	// Retrieve the guild object
	guild := ctx.CustomObjects.MustGet("guild").(*guilds.Guild)

	// Validate the channel itself
	_, err := ctx.Session.Channel(channelID)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}

	// Set the starboard channel ID
	guild.Settings.Starboard.Channel = channelID
	err = guild.Update()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}

	// Respond with a success message
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Success("The starboard channel has been set to "+channelID+"."))
}

// settingsStarboardMinimumCommand handles the 'settings starboard minimum' command
func settingsStarboardMinimumCommand(ctx *dgc.Ctx) {
	// Validate the minimum star amount
	minimum, err := ctx.Arguments.AsSingle().AsInt()
	if err != nil {
		if err != nil {
			ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage("You need to specify a minimum star amount."))
			return
		}
	}
	if minimum <= 0 {
		minimum = 1
	}

	// Retrieve the guild object
	guild := ctx.CustomObjects.MustGet("guild").(*guilds.Guild)

	// Set the minimum star amount
	guild.Settings.Starboard.Minimum = minimum
	err = guild.Update()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}

	// Respond with a success message
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Success("The minimum star amount has been set to "+strconv.Itoa(minimum)+"."))
}
