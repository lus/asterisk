package commands

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Latex handles the latex command
func Latex(ctx *dgc.Ctx) {
	// Validate the arguments
	codeblock := ctx.Arguments.AsCodeblock()
	if codeblock == nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage(ctx.Command.Usage))
		return
	}

	// Render the given expression and respond with it
	url, err := utils.RenderLaTeX(codeblock.Content)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.LaTeX(url))
}
