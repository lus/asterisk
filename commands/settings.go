package commands

import (
	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Settings handles the settings command
func Settings(ctx *dgc.Ctx) {
	// Respond with the current guild configuration
	guildConfig := ctx.CustomObjects.MustGet("guildConfig").(*guildconfig.GuildConfig)
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateGuildSettingsEmbed(guildConfig))
}

// SettingsToggleChannelRestriction handles the settings toggleChannelRestriction command
func SettingsToggleChannelRestriction(ctx *dgc.Ctx) {
	// Toggle the channel restriction status
	guildConfig := ctx.CustomObjects.MustGet("guildConfig").(*guildconfig.GuildConfig)
	guildConfig.ChannelRestriction = !guildConfig.ChannelRestriction
	err := guildConfig.Update()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}

	// Respond with a success message
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("The command channel restriction has been "+utils.PrettifyBool(guildConfig.ChannelRestriction)+"."))
}

// SettingsToggleCommandChannel handles the settings toggleCommandChannel command
func SettingsToggleCommandChannel(ctx *dgc.Ctx) {
	// Validate the argument
	channelID := ctx.Arguments.Get(0).AsChannelMentionID()
	if channelID == "" {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed(ctx.Command.Usage))
		return
	}

	// Define the guild configuration
	guildConfig := ctx.CustomObjects.MustGet("guildConfig").(*guildconfig.GuildConfig)

	// Toggle the command channel status
	contains := utils.StringArrayContains(guildConfig.CommandChannels, channelID)
	if contains {
		newArray := make([]string, len(guildConfig.CommandChannels)-1)
		counter := 0
		for _, channel := range guildConfig.CommandChannels {
			if channel == channelID {
				continue
			}
			newArray[counter] = channel
			counter++
		}
		guildConfig.CommandChannels = newArray
	} else {
		guildConfig.CommandChannels = append(guildConfig.CommandChannels, channelID)
	}
	err := guildConfig.Update()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}

	// Respond with a success message
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("The command channel status for the mentioned channel has been "+utils.PrettifyBool(!contains)+"."))
}

// SettingsToggleHastebinIntegration handles the settings toggleHastebinIntegration command
func SettingsToggleHastebinIntegration(ctx *dgc.Ctx) {
	// Toggle the hastebin integration status
	guildConfig := ctx.CustomObjects.MustGet("guildConfig").(*guildconfig.GuildConfig)
	guildConfig.HastebinIntegration = !guildConfig.HastebinIntegration
	err := guildConfig.Update()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}

	// Respond with a success message
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("The hastebin integration has been "+utils.PrettifyBool(guildConfig.HastebinIntegration)+"."))
}
