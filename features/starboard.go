package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/guilds"
	"github.com/bwmarrin/discordgo"
)

// initializeStarboardFeature initializes the starboard feature
func initializeStarboardFeature(session *discordgo.Session) {
	// Register the starboardReactionListener
	session.AddHandler(starboardReactionListener)
}

// starboardReactionListener has to be registered to enable the starboard feature
func starboardReactionListener(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	// Check if the reaction is a star reaction
	if event.Emoji.Name != "⭐" && event.Emoji.Name != "🌟" {
		return
	}

	// Retrieve the guild object
	guild, err := guilds.RetrieveCached(event.GuildID)
	if err != nil {
		return
	}

	// Retrieve the starboard channel ID
	starboardChannelID := guild.Settings.StarboardChannel
	if starboardChannelID == "" {
		return
	}

	// Retrieve the message object
	message, err := session.ChannelMessage(event.ChannelID, event.MessageID)
	if err != nil {
		return
	}
	message.GuildID = event.GuildID

	// Retrieve the amount of star emojis and check if the message already got posted into the starboard
	amount := 0
	alreadyPosted := false
	for _, reactions := range message.Reactions {
		if reactions.Emoji.Name != "⭐" && reactions.Emoji.Name != "🌟" {
			continue
		}
		if reactions.Me {
			alreadyPosted = true
		}
		amount += reactions.Count
	}

	// Check if the amount of star emojis is at least 2 and if the message already got posted into the starboard
	if amount < 2 || alreadyPosted {
		return
	}

	// Add the message to the starboard
	session.ChannelMessageSendEmbed(starboardChannelID, embeds.Starboard(message))
	session.MessageReactionAdd(event.ChannelID, event.MessageID, "⭐")
}
