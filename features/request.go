package features

import (
	"github.com/Lukaesebrot/asterisk/config"
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/users"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// initializeRequestFeature initializes the request feature
func initializeRequestFeature(router *dgc.Router, rateLimiter dgc.RateLimiter, session *discordgo.Session) {
	// Register the 'request' command
	router.RegisterCmd(&dgc.Command{
		Name:        "request",
		Description: "Sends a feature request to the developers",
		Usage:       "request <description>",
		Example:     "request More hashing algorithms.",
		IgnoreCase:  true,
		RateLimiter: rateLimiter,
		Handler:     requestCommand,
	})

	// Register the requestReactionListener
	session.AddHandler(requestReactionListener)
}

// requestCommand handles the 'request' command
func requestCommand(ctx *dgc.Ctx) {
	// Validate the input
	if ctx.Arguments.Amount() == 0 {
		ctx.RespondEmbed(embeds.InvalidUsage("You need to specify a feature you want to request."))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Send the feature request to the feature request channel and add the delete emote
	message, err := ctx.Session.ChannelMessageSendEmbed(config.CurrentConfig.FeatureRequestChannel, embeds.FeatureRequest(ctx))
	if err != nil {
		ctx.RespondEmbed(embeds.Error("Your feature request couldn't be submitted. Please try again later."))
		return
	}
	ctx.Session.MessageReactionAdd(config.CurrentConfig.FeatureRequestChannel, message.ID, "✅")

	// Confirm the creation of the feature request
	ctx.RespondEmbed(embeds.Success("Your feature request got submitted."))
}

// requestReactionListener has to be registered to enable the tick reaction on feature requests
func requestReactionListener(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	// Check if the channel is the feature request channel
	if event.ChannelID != config.CurrentConfig.FeatureRequestChannel {
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
