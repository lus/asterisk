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

// GenerateInvalidUsageEmbed generates an embed for invalid usages
func GenerateInvalidUsageEmbed(usage string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Invalid Usage",
		Description: "That's not how you're supposed to use this command!",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Message",
				Value:  "```" + usage + "```",
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

// GenerateRandomOutputEmbed generates an embed for random outputs
func GenerateRandomOutputEmbed(output string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Random Output",
		Description: "Here's your (pseudo) random output.",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Output",
				Value:  "```" + output + "```",
				Inline: false,
			},
		},
	}
}
