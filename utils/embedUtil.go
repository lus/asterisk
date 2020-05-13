package utils

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Lukaesebrot/asterisk/static"
	"github.com/bwmarrin/discordgo"
)

// GenerateSuccessEmbed generates a general success embed
func GenerateSuccessEmbed(output string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Success",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x00ff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Output",
				Value:  "```" + output + "```",
				Inline: false,
			},
		},
	}
}

// GenerateErrorEmbed generates an embed for errors
func GenerateErrorEmbed(message string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Error",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Message",
				Value:  "```" + strings.Replace(message, "`", "'", -1) + "```",
				Inline: false,
			},
		},
	}
}

// GenerateInvalidUsageEmbed generates an embed for invalid usages
func GenerateInvalidUsageEmbed(message string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Invalid Usage",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Message",
				Value:  "```" + strings.Replace(message, "`", "'", -1) + "```",
				Inline: false,
			},
		},
	}
}

// GenerateInsufficientPermissionsEmbed generates an embed for insufficient permissions
func GenerateInsufficientPermissionsEmbed(message string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Insufficient Permissions",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Message",
				Value:  "```" + strings.Replace(message, "`", "'", -1) + "```",
				Inline: false,
			},
		},
	}
}

// GenerateBotInfoEmbed generates the embed which contains all the neccessary bot information
func GenerateBotInfoEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Bot Information",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
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
		Type:      "rich",
		Title:     "Stats",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Discord specs",
				Value: "Guilds: `" + strconv.Itoa(len(session.State.Guilds)) + "`" +
					"\nAPI latency: `" + strconv.FormatInt(session.HeartbeatLatency().Milliseconds(), 10) + "ms`",
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
