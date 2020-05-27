package middlewares

import (
	"github.com/Lukaesebrot/asterisk/guilds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// CheckCommandChannel checks whether or not the current channel is a command channel
func CheckCommandChannel(next dgc.ExecutionHandler) dgc.ExecutionHandler {
	return func(ctx *dgc.Ctx) {
		guild := ctx.CustomObjects.MustGet("guild").(*guilds.Guild)
		if len(guild.Settings.CommandChannels) == 0 || utils.StringArrayContains(guild.Settings.CommandChannels, ctx.Event.ChannelID) || utils.StringArrayContains(ctx.Command.Flags, "ignore_command_channel") {
			next(ctx)
		}
	}
}
