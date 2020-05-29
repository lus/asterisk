package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
)

// initializeInfoFeature initializes the info feature
func initializeInfoFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'info' command
	router.RegisterCmd(&dgc.Command{
		Name:        "info",
		Description: "Displays some general information about the bot",
		Usage:       "info",
		Example:     "info",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     infoCommand,
	})
}

// infoCommand handles the 'info' command
func infoCommand(ctx *dgc.Ctx) {
	// Check the rate limiter
	if ctx.Command != nil && !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	ctx.RespondEmbed(embeds.Info(ctx))
}
