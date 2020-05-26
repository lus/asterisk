package embeds

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Starboard generates a starboard embed
func Starboard(message *discordgo.Message) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         fmt.Sprintf("https://discord.com/channels/%s/%s/%s/", message.GuildID, message.ChannelID, message.ID),
		Title:       "Jump To Message",
		Description: message.Content,
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    message.Author.String(),
			IconURL: message.Author.AvatarURL("512"),
		},
	}
}
