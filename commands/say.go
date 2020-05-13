package commands

import (
	"github.com/Lukaesebrot/dgc"

	"github.com/Lukaesebrot/asterisk/utils"
)

// Say handles the say command
func Say(ctx *dgc.Ctx) {
	// Check if the executor is a bot admin
	if !utils.IsBotAdmin(ctx.Event.Author.ID) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInsufficientPermissionsEmbed("You need to be a bot admin to use this command."))
		return
	}

	// Repeat the input
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("You need to specify a string."))
		return
	}
	ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, ctx.Arguments.Raw())
}
