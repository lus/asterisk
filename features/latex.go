package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// initializeLaTeXFeature initializes the LaTeX feature
func initializeLaTeXFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'latex' command
	router.RegisterCmd(&dgc.Command{
		Name:        "latex",
		Description: "Renders the given LaTeX expression as an image",
		Usage:       "latex <codeblock>",
		Example:     "latex `$ 1+4^6 $`\n",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     latexCommand,
	})
}

// latexCommand handles the 'latex' command
func latexCommand(ctx *dgc.Ctx) {
	// Validate the arguments
	codeblock := ctx.Arguments.AsCodeblock()
	if codeblock == nil {
		ctx.RespondEmbed(embeds.InvalidUsage(ctx.Command.Usage))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Respond with a loading embed
	msg, _ := ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Loading())

	// Render the given expression and respond with it
	url, err := utils.RenderLaTeX(codeblock.Content)
	if err != nil {
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}
	ctx.Session.ChannelMessageEditEmbed(ctx.Event.ChannelID, msg.ID, embeds.LaTeX(url))
}
