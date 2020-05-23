package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// LaTeX generates a LaTeX result embed
func LaTeX(url string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "LaTeX result",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x00ff00,
		Image: &discordgo.MessageEmbedImage{
			URL: url,
		},
	}
}
