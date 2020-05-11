package utils

import (
	"time"

	"github.com/Lukaesebrot/asterisk/static"
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

// GenerateBotInfoEmbed generates the embed which contains all the neccessary bot information
func GenerateBotInfoEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Information",
		Description: "Here you will find some information about me :)",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Developer(s)",
				Value:  "`Lukaesebrot#8001`",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "GitHub Repository",
				Value:  "http://github.com/Lukaesebrot/asterisk",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Invite me",
				Value:  "https://discord.com/api/oauth2/authorize?client_id=" + static.Self.ID + "&permissions=0&scope=bot",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Support guild",
				Value:  "https://discord.gg/ddz9b86",
				Inline: false,
			},
		},
	}
}
