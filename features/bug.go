package features

import (
	"github.com/Lukaesebrot/asterisk/config"
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/users"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// initializeBugFeature initializes the bug feature
func initializeBugFeature(router *dgc.Router, rateLimiter dgc.RateLimiter, session *discordgo.Session) {
	// Register the 'bug' command
	router.RegisterCmd(&dgc.Command{
		Name:        "bug",
		Description: "Reports a bug to the developers",
		Usage:       "bug <description>",
		Example:     "bug The bot spams!",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     bugCommand,
	})

	// Register the bugReactionListener
	session.AddHandler(bugReactionListener)
}

// bugCommand handles the 'bug' command
func bugCommand(ctx *dgc.Ctx) {
	// Validate the input
	if ctx.Arguments.Amount() == 0 {
		ctx.RespondEmbed(embeds.InvalidUsage("You need to specify a description of the bug you want to report."))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Send the bug report to the bug report channel and add the delete emote
	message, err := ctx.Session.ChannelMessageSendEmbed(config.CurrentConfig.Channels.BugReports, embeds.BugReport(ctx))
	if err != nil {
		ctx.RespondEmbed(embeds.Error("Your bug report couldn't be submitted. Please try again later."))
		return
	}
	ctx.Session.MessageReactionAdd(config.CurrentConfig.Channels.BugReports, message.ID, "✅")

	// Confirm the creation of the feature request
	ctx.RespondEmbed(embeds.Success("Your bug report got submitted."))
}

// bugReactionListener has to be registered to enable the tick reaction on bug reports
func bugReactionListener(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	// Check if the channel is the bug report channel
	if event.ChannelID != config.CurrentConfig.Channels.BugReports {
		return
	}

	// Check if the user is a bot admin
	user, err := users.RetrieveCached(event.UserID)
	if err != nil || !user.HasFlag(users.FlagAdministrator) {
		return
	}

	// Check if the reaction is the tick reaction
	if event.Emoji.Name != "✅" {
		return
	}

	// Delete the message
	session.ChannelMessageDelete(event.ChannelID, event.MessageID)
}
