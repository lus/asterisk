package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
)

// initializeSayFeature initializes the say feature
func initializeSayFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
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
		RateLimiter: rateLimiter,
		Handler:     sayCommand,
	})
}

// sayCommand handles the 'say' command
func sayCommand(ctx *dgc.Ctx) {
	// Validate the input
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage("You need to specify a string that I should say."))
		return
	}
	ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, ctx.Arguments.Raw())
}
