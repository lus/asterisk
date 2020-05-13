package commands

import (
	"github.com/Lukaesebrot/dgc"

	"github.com/Lukaesebrot/asterisk/utils"
)

// Say handles the say command
func Say(ctx *dgc.Ctx) {
	// Validate the input
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("You need to specify a string that I should say."))
		return
	}
	ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, ctx.Arguments.Raw())
}
