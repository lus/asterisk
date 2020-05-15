package commands

import (
	"strings"

	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// Hastebin handles the hastebin command
func Hastebin(ctx *dgc.Ctx) {
	// Validate the arguments
	codeblock := ctx.Arguments.AsCodeblock()
	if codeblock == nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed(ctx.Command.Usage))
		return
	}
	content := strings.TrimSpace(codeblock.Content)
	if content == "" {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("You don't want to create an empty haste, do you?"))
		return
	}

	// Create the haste
	hasteURL, err := utils.CreateHaste(content)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}

	// Respond with the haste URL
	ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, hasteURL)
}

// HastebinMessageCreateListener defines the message create listener for the hastebin feature
func HastebinMessageCreateListener(session *discordgo.Session, event *discordgo.MessageCreate) {
	// Parse the message content into a codeblock
	arguments := dgc.ParseArguments(event.Content)
	codeblock := arguments.AsCodeblock()
	if codeblock == nil || strings.TrimSpace(codeblock.Content) == "" {
		return
	}

	// Add the clipboard reaction
	session.MessageReactionAdd(event.ChannelID, event.Message.ID, "📋")
}

// HastebinReactionAddListener defines the reaction add listener for the hastebin feature
func HastebinReactionAddListener(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	// Check if the reaction is the clipboard reaction
	if event.Emoji.Name != "📋" {
		return
	}

	// Check if the emoji was added by a bot
	user, err := session.User(event.UserID)
	if err != nil || user.Bot {
		return
	}

	// Retrieve the message
	message, err := session.State.Message(event.ChannelID, event.MessageID)
	if err != nil {
		message, err = session.ChannelMessage(event.ChannelID, event.MessageID)
		if err != nil {
			return
		}
	}

	// Parse the message content into a codeblock
	arguments := dgc.ParseArguments(message.Content)
	codeblock := arguments.AsCodeblock()
	if codeblock == nil {
		return
	}
	content := strings.TrimSpace(codeblock.Content)
	if content == "" {
		return
	}

	// Create the haste
	hasteURL, err := utils.CreateHaste(content)
	if err != nil {
		return
	}

	// Respond with the haste URL
	session.ChannelMessageSend(event.ChannelID, hasteURL)

	// Remove the reaction
	session.MessageReactionRemove(event.ChannelID, event.MessageID, "📋", event.UserID)
}
