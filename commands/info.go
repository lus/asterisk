package commands

import (
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Info handles the info command
func Info(ctx *dgc.Ctx) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateBotInfoEmbed())
}
