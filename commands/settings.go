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
