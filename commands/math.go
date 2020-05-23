package commands

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Math handles the math command
func Math(ctx *dgc.Ctx) {
	// Validate the arguments
	codeblock := ctx.Arguments.AsCodeblock()
	if codeblock == nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage(ctx.Command.Usage))
		return
	}

	// Evaluate the expression and respond with the result
	result, err := utils.EvaluateMathematicalExpression(codeblock.Content)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Success(result))
}
