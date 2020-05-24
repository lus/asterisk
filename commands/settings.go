package commands

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/guilds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Settings handles the settings command
func Settings(ctx *dgc.Ctx) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Settings(ctx.CustomObjects.MustGet("guild").(*guilds.Guild)))
}

// SettingsCommandChannel handles the settings commandChannel command
func SettingsCommandChannel(ctx *dgc.Ctx) {
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
