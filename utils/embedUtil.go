package utils

import (
	"runtime"
	"strconv"
	"strings"
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
				Value:  "```" + strings.Replace(errorMessage, "`", "'", -1) + "```",
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
				Value:  "```" + strings.Replace(usage, "`", "'", -1) + "```",
				Inline: false,
			},
		},
	}
}

// GenerateInsufficientPermissionsEmbed generates an embed for insufficient permissions
func GenerateInsufficientPermissionsEmbed(message string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Insufficient Permissions",
		Description: "Looks like you're not cool enough to use that command :/",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Message",
				Value:  "```" + strings.Replace(message, "`", "'", -1) + "```",
				Inline: false,
			},
		},
	}
}

// GenerateHelpEmbed generates the embed which contains all the commands
func GenerateHelpEmbed(commandNames []string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Help",
		Description: "Here you will find all available command. Use `$help <command>` to find out more.",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Commands",
				Value:  "`" + strings.Join(commandNames, "`, `") + "`",
				Inline: false,
			},
		},
	}
}

// GenerateSingleHelpEmbed generates the embed which contains the description of a single
func GenerateSingleHelpEmbed(commandName, commandDescription string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Help",
		Description: "Here comes your help for the `" + commandName + "` command.",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Description",
				Value:  "```" + strings.Replace(commandDescription, "`", "'", -1) + "```",
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

// GenerateStatsEmbed generates the embed which contains all the useful stats
func GenerateStatsEmbed(session *discordgo.Session) *discordgo.MessageEmbed {
	// Read the memory stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Stats",
		Description: "Here you will find some useful stats.",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Discord specs",
				Value:  "Guilds: `" + strconv.Itoa(len(session.State.Guilds)) + "`",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name: "System specs",
				Value: "Current Goroutines: `" + strconv.Itoa(runtime.NumGoroutine()) + "`" +
					"\nUsable logical CPUs: `" + strconv.Itoa(runtime.NumCPU()) + "`" +
					"\nHeap in use: `" + strconv.FormatUint(memStats.HeapInuse, 10) + "`" +
					"\nStack in use: `" + strconv.FormatUint(memStats.StackInuse, 10) + "`",
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
				Value:  "```" + strings.Replace(output, "`", "'", -1) + "```",
				Inline: false,
			},
		},
	}
}

// GenerateArbitraryOutputEmbed generates an embed for arbitrary outputs
func GenerateArbitraryOutputEmbed(output string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		URL:         "https://github.com/Lukaesebrot/asterisk",
		Type:        "rich",
		Title:       "Arbitrary Expression Output",
		Description: "Here's the output of your arbitrary expression..",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Output",
				Value:  "```" + strings.Replace(output, "`", "'", -1) + "```",
				Inline: false,
			},
		},
	}
}
