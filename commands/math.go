package commands

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Math handles the math command
func Math(ctx *dgc.Ctx) {
	// Validate the arguments
	codeblock := ctx.Arguments.AsCodeblock()
	if codeblock == nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed(ctx.Command.Usage))
		return
	}

	// Create the evaluable expression
	expression, err := govaluate.NewEvaluableExpression(codeblock.Content)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}

	// Evaluate the expression and respond with the result
	params := make(map[string]interface{})
	result, err := expression.Evaluate(params)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed(fmt.Sprintf("%v", result)))
}
