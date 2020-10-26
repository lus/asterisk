package embeds

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Success generates a basic success embed
func Success(output string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "Success",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Output",
				Value: "```" + strings.ReplaceAll(output, "`", "'") + "```",
			},
		},
	}
}

// Error generates a basic error embed
func Error(message string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "Error",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Message",
				Value: "```" + strings.ReplaceAll(message, "`", "'") + "```",
			},
		},
	}
}

// InvalidUsage generates a basic invalid usage embed
func InvalidUsage(usage string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "Invalid Usage",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Usage",
				Value: "```" + strings.ReplaceAll(usage, "`", "'") + "```",
			},
		},
	}
}

// InsufficientPermissions generates an insufficient permissions embed
func InsufficientPermissions(message string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "Insufficient Permissions",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Message",
				Value: "```" + strings.ReplaceAll(message, "`", "'") + "```",
			},
		},
	}
}
