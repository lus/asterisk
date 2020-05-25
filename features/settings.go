package features

import (
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
				Usage:       "settings commandChannel <channel mention>",
				Example:     "settings commandChannel #my-channel",
				Flags: []string{
					"guild_admin",
					"ignore_command_channel",
				},
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     settingsCommandChannelCommand,
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
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
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
