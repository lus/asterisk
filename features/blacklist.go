package features

import (
	"fmt"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/nodes/users"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// initializeBlacklistFeature initializes the blacklist feature
func initializeBlacklistFeature(router *dgc.Router) {
	// Register the 'blacklist' command
	router.RegisterCmd(&dgc.Command{
		Name:        "blacklist",
		Description: "[Bot Moderator only] Toggles the blacklist status of the given user",
		Usage:       "blacklist <user mention or ID>",
		Example:     "blacklist @Erik",
		Flags: []string{
			"bot_mod",
		},
		IgnoreCase:  true,
		RateLimiter: nil,
		Handler:     blacklistCommand,
	})
}

// blacklistCommand handles the 'blacklist' command
func blacklistCommand(ctx *dgc.Ctx) {
	// Validate the user ID
	userID := ctx.Arguments.AsSingle().AsUserMentionID()
	if userID == "" {
		userID = ctx.Arguments.Raw()
	}

	// Validate the user itself
	_, err := ctx.Session.User(userID)
	if err != nil {
		ctx.RespondEmbed(embeds.Error("The given user couldn't be found."))
		return
	}

	// Retrieve the internal user object
	user, err := users.RetrieveCached(userID)
	if err != nil {
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}

	// Toggle the blacklist status of the given user
	if user.HasFlag(users.FlagBlacklisted) {
		user.RevokeFlag(users.FlagBlacklisted)
	} else {
		user.AssignFlag(users.FlagBlacklisted)
	}

	// Update the user object
	err = user.Update()
	if err != nil {
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}
	ctx.RespondEmbed(embeds.Success(fmt.Sprintf("The blacklist status of the given user has been %s.", utils.PrettifyBool(user.HasFlag(users.FlagBlacklisted)))))
}
