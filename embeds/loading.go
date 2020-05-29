package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Loading generates a loading embed
func Loading() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Loading",
		Description: "Your result is being loaded...",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
	}
}
