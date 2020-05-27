package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
)

// initializeCreditsFeature initializes the credits feature
func initializeCreditsFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'credits' command
	router.RegisterCmd(&dgc.Command{
		Name:        "credits",
		Description: "Displays some credits",
		Usage:       "credits",
		Example:     "credits",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     creditsCommand,
	})
}

// creditsCommand handles the 'credits' command
func creditsCommand(ctx *dgc.Ctx) {
	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	ctx.RespondEmbed(embeds.Credits())
}
