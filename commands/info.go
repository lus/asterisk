package commands

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
)

// Info handles the info command
func Info(ctx *dgc.Ctx) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Info(ctx))
}
