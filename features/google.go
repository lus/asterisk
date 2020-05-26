package features

import (
	"fmt"
	"net/url"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
)

// initializeGoogleFeature initializes the google feature
func initializeGoogleFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'google' command
	router.RegisterCmd(&dgc.Command{
		Name:        "google",
		Aliases:     []string{"lmgtfy", "search"},
		Description: "Creates a lmgtfy link with the given parameters",
		Usage:       "google <parameters>",
		Example:     "google How to bake an apple cake",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     googleCommand,
	})
}

// googleCommand handles the 'google' command
func googleCommand(ctx *dgc.Ctx) {
	// Validate the input
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage("You need to specify some search parameters."))
		return
	}

	// Respond with the the lmgtfy URL
	ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, fmt.Sprintf("https://lmgtfy.com/?q=%s", url.QueryEscape(ctx.Arguments.Raw())))
}
