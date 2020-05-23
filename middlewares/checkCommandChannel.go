package middlewares

import (
	"github.com/Lukaesebrot/asterisk/guilds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// CheckCommandChannel checks whether or not the current channel is a command channel
func CheckCommandChannel(ctx *dgc.Ctx) bool {
	guild := ctx.CustomObjects.MustGet("guild").(*guilds.Guild)
	return utils.StringArrayContains(guild.Settings.CommandChannels, ctx.Event.ChannelID)
}
