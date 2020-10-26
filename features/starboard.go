package features

import (
	"fmt"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/nodes/guilds"
	"github.com/Lukaesebrot/asterisk/nodes/starboard"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
)

// initializeStarboardFeature initializes the starboard feature
func initializeStarboardFeature(session *discordgo.Session) {
	// Register the starboardReactionAddListener
	session.AddHandler(starboardReactionAddListener)

	// Register the starboardReactionRemoveListener
	session.AddHandler(starboardReactionRemoveListener)
}

// starboardReactionAddListener has to be registered to enable the starboard feature
func starboardReactionAddListener(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	// Check if the reaction is a star reaction
	if event.Emoji.Name != "⭐" {
		return
	}

	// Retrieve the message object
	message, err := session.State.Message(event.ChannelID, event.MessageID)
	if err != nil {
		message, err = session.ChannelMessage(event.ChannelID, event.MessageID)
		if err != nil {
			return
		}
	}
	message.GuildID = event.GuildID

	// Check if the message author is a bot
	if message.Author.Bot {
		return
	}

	// Get the amount of stars
	amount := 0
	for _, reaction := range message.Reactions {
		if reaction.Emoji.Name == "⭐" {
			amount = reaction.Count
			break
		}
	}
	count := fmt.Sprintf("%d ⭐", amount)

	// Check if the message is already marked as a starboard entry
	entry, err := starboard.GetEntry(event.ChannelID, event.MessageID)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}
	if entry != nil {
		// Get the starboard message object
		sbMessage, err := session.State.Message(entry.StarboardChannel, entry.StarboardMessageID)
		if err != nil {
			sbMessage, err = session.ChannelMessage(entry.StarboardChannel, entry.StarboardMessageID)
			if err != nil {
				return
			}
		}

		// Update the amount of stars
		session.ChannelMessageEditComplex(&discordgo.MessageEdit{
			Channel: entry.StarboardChannel,
			ID:      entry.StarboardMessageID,
			Content: &count,
			Embed:   sbMessage.Embeds[0],
		})
		return
	}

	// Retrieve the guild object
	guild, err := guilds.RetrieveCached(event.GuildID)
	if err != nil {
		return
	}

	// Retrieve the starboard channel ID
	starboardChannelID := guild.Settings.Starboard.Channel
	if starboardChannelID == "" {
		return
	}

	// Validate the amount of stars
	if amount < guild.Settings.Starboard.Minimum {
		return
	}

	// Create a starboard entry for the message
	msg, err := session.ChannelMessageSendComplex(starboardChannelID, &discordgo.MessageSend{
		Content: count,
		Embed:   embeds.Starboard(message),
	})
	if err != nil {
		return
	}
	starboard.AddEntry(event.ChannelID, event.MessageID, starboardChannelID, msg.ID)
}

// starboardReactionRemoveListener has to be registered to enable the starboard feature
func starboardReactionRemoveListener(session *discordgo.Session, event *discordgo.MessageReactionRemove) {
	// Check if the reaction is a star reaction
	if event.Emoji.Name != "⭐" {
		return
	}

	// Retrieve the message object
	message, err := session.State.Message(event.ChannelID, event.MessageID)
	if err != nil {
		message, err = session.ChannelMessage(event.ChannelID, event.MessageID)
		if err != nil {
			return
		}
	}
	message.GuildID = event.GuildID

	// Get the amount of stars
	amount := 0
	for _, reaction := range message.Reactions {
		if reaction.Emoji.Name == "⭐" {
			amount = reaction.Count
			break
		}
	}
	count := fmt.Sprintf("%d ⭐", amount)

	// Check if the message is already marked as a starboard entry
	entry, err := starboard.GetEntry(event.ChannelID, event.MessageID)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}
	if entry != nil {
		// Get the starboard message object
		sbMessage, err := session.State.Message(entry.StarboardChannel, entry.StarboardMessageID)
		if err != nil {
			sbMessage, err = session.ChannelMessage(entry.StarboardChannel, entry.StarboardMessageID)
			if err != nil {
				return
			}
		}

		// Update the amount of stars
		session.ChannelMessageEditComplex(&discordgo.MessageEdit{
			Channel: entry.StarboardChannel,
			ID:      entry.StarboardMessageID,
			Content: &count,
			Embed:   sbMessage.Embeds[0],
		})
		return
	}
}
