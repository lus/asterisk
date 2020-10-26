package embeds

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Starboard generates a starboard embed
func Starboard(message *discordgo.Message) *discordgo.MessageEmbed {
	desc := fmt.Sprintf(
		"%s\n\n[*jump to message*](https://discordapp.com/channels/%s/%s/%s)",
		message.Content, message.GuildID, message.ChannelID, message.ID)

	emb := &discordgo.MessageEmbed{
		Description: desc,
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    message.Author.String(),
			IconURL: message.Author.AvatarURL("512"),
		},
	}

	if len(message.Attachments) > 0 {
		emb.Image = &discordgo.MessageEmbedImage{
			URL: message.Attachments[0].URL,
		}
	}

	return emb
}
