package utils

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// GenerateInternalErrorEmbed generates an embed for internal errors
func GenerateInternalErrorEmbed(errorMessage string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Internal Error",
		Description: "An internal error occured :/",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Error message",
				Value:  "```" + errorMessage + "```",
				Inline: false,
			},
		},
	}
}
