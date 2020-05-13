package middlewares

import (
	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// CheckCommandChannel checks if the current channel is a valid command channel
func CheckCommandChannel(ctx *dgc.Ctx) bool {
	guildConfig := ctx.CustomObjects["guildConfig"].(*guildconfig.GuildConfig)
	return !guildConfig.ChannelRestriction || utils.StringArrayContains(guildConfig.CommandChannels, ctx.Event.ChannelID)
}
