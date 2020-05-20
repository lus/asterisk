package commands

import (
	"github.com/Lukaesebrot/asterisk/config"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// Bug handles the bug command
func Bug(ctx *dgc.Ctx) {
	// Validate the input
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("You need to specify a description of the bug you want to report."))
		return
	}

	// Send the bug report to the bug report channel and add the delete emote
	message, err := ctx.Session.ChannelMessageSendEmbed(config.CurrentConfig.BugReportChannel, utils.GenerateBugReportEmbed(ctx))
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed("Your bug report couldn't be submitted. Please try again later."))
		return
	}
	ctx.Session.MessageReactionAdd(config.CurrentConfig.BugReportChannel, message.ID, "✅")

	// Confirm the creation of the feature request
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("Your bug report got submitted."))
}

// BugReactionListener has to be registered to enable the tick reaction on bug reports
func BugReactionListener(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	// Check if the channel is the bug report channel
	if event.ChannelID != config.CurrentConfig.BugReportChannel {
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
