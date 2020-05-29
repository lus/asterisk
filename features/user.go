package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/users"
	"github.com/Lukaesebrot/dgc"
)

// initializeUserFeature initializes the user feature
func initializeUserFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'user' command
	router.RegisterCmd(&dgc.Command{
		Name:        "user",
		Description: "Displays some general information about an user",
		Usage:       "user <user mention or ID>",
		Example:     "user @Erik",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     userCommand,
	})
}

// userCommand handles the 'user' command
func userCommand(ctx *dgc.Ctx) {
	// Validate the user ID
	userID := ctx.Arguments.AsSingle().AsUserMentionID()
	if userID == "" {
		userID = ctx.Arguments.Raw()
	}

	// Validate the user itself
	dcUser, err := ctx.Session.User(userID)
	if err != nil {
		ctx.RespondEmbed(embeds.Error("The given user couldn't be found."))
		return
	}

	// Retrieve the internal user and inject it into the custom objects
	intUser, err := users.RetrieveCached(userID)
	if err != nil {
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}

	// Respond with the user embed
	ctx.RespondEmbed(embeds.User(dcUser, intUser))
}
