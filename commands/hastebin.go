package commands

import (
	"strings"

	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// HastebinMessageCreateListener defines the message create listener for the hastebin feature
func HastebinMessageCreateListener(session *discordgo.Session, event *discordgo.MessageCreate) {
	// Parse the message content into a codeblock
	arguments := dgc.ParseArguments(event.Content)
	codeblock := arguments.AsCodeblock()
	if codeblock == nil || strings.TrimSpace(codeblock.Content) == "" || len(strings.TrimSpace(codeblock.Content)) < 500 {
		return
	}

	// Check if the guild activated the hastebin feature
	guildConfig, err := guildconfig.Retrieve(event.GuildID)
	if err != nil || !guildConfig.HastebinIntegration {
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

	// Check if the message already has a reaction by the bot itself
	valid := false
	for _, reaction := range message.Reactions {
		if reaction.Me {
			valid = true
			break
		}
	}
	if !valid {
		return
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

	// Remove the reaction of the user and the own one
	session.MessageReactionRemove(event.ChannelID, event.MessageID, "📋", event.UserID)
	session.MessageReactionRemove(event.ChannelID, event.MessageID, "📋", "@me")
}
