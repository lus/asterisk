package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
)

// initializeSayFeature initializes the say feature
func initializeSayFeature(router *dgc.Router) {
	// Register the 'say' command
	router.RegisterCmd(&dgc.Command{
		Name:        "say",
		Description: "[Bot Admin only] Makes the bot say something inside the current channel",
		Usage:       "say <string>",
		Example:     "say Hello, world!",
		Flags: []string{
			"bot_admin",
		},
		IgnoreCase:  true,
		RateLimiter: nil,
		Handler:     sayCommand,
	})
}

// sayCommand handles the 'say' command
func sayCommand(ctx *dgc.Ctx) {
	// Validate the input
	if ctx.Arguments.Amount() == 0 {
		ctx.RespondEmbed(embeds.InvalidUsage("You need to specify a string that I should say."))
		return
	}
	ctx.RespondText(ctx.Arguments.Raw())
}
