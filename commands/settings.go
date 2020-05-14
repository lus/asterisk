package commands

import (
	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Settings handles the settings command
func Settings(ctx *dgc.Ctx) {
	// Respond with the current guild configuration
	guildConfig := ctx.CustomObjects["guildConfig"].(*guildconfig.GuildConfig)
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateGuildSettingsEmbed(guildConfig))
}

// SettingsToggleChannelRestriction handles the settings toggleChannelRestriction command
func SettingsToggleChannelRestriction(ctx *dgc.Ctx) {
	// Toggle the channel restriction status
	guildConfig := ctx.CustomObjects["guildConfig"].(*guildconfig.GuildConfig)
	guildConfig.ChannelRestriction = !guildConfig.ChannelRestriction
	err := guildConfig.Update()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}

	// Respond with a success message
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("The command channel restriction has been "+utils.PrettifyBool(guildConfig.ChannelRestriction)+"."))
}
