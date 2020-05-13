package commands

import (
	"time"

	"github.com/Lukaesebrot/asterisk/config"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// lastRequestCalls holds the timestamp at which the given user ID (key) last executed the request command
var lastRequestCalls = make(map[string]int64)

// Request handles the request command
func Request(ctx *dgc.Ctx) {
	// Check if the user is being time-outed
	lastCall := lastRequestCalls[ctx.Event.Member.User.ID]
	currentTime := time.Now().UnixNano() / 1e6
	if (lastCall != 0) && (currentTime-lastCall) < 3600000 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed("You need to wait at least one hour between two feature requests."))
		return
	}

	// Validate the input
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("You need to specify a feature you want to request."))
		return
	}

	// Send the feature request to the feature request channel and add the delete emote
	message, err := ctx.Session.ChannelMessageSendEmbed(config.CurrentConfig.FeatureRequestChannel, utils.GenerateFeatureRequestEmbed(ctx))
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed("Your feature request couldn't be submitted. Please try again later."))
		return
	}
	ctx.Session.MessageReactionAdd(config.CurrentConfig.FeatureRequestChannel, message.ID, "✅")

	// Save the current execution time
	lastRequestCalls[ctx.Event.Member.User.ID] = currentTime

	// Confirm the creation of the feature request
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("Your feature request was submitted."))
}

// RequestReactionListener has to be registered to enable the tick reaction on feature requests
func RequestReactionListener(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	// Check if the channel is the feature request channel
	if event.ChannelID != config.CurrentConfig.FeatureRequestChannel {
		return
	}

	// Check if the user is a bot admin
	if !utils.StringArrayContains(config.CurrentConfig.BotAdmins, event.UserID) {
		return
	}

	// Check of the reaction is the tick reaction
	if event.Emoji.Name != "✅" {
		return
	}

	// Delete the message
	session.ChannelMessageDelete(event.ChannelID, event.MessageID)
}
