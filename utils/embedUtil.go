package utils

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/static"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
	"github.com/c2h5oh/datasize"
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
	heap := datasize.ByteSize(memStats.HeapInuse)
	stack := datasize.ByteSize(memStats.StackInuse)

	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Stats",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Application specs",
				Value: "Mode: `" + static.Mode + "`" +
					"\nVersion: `" + static.Version + "`",
				Inline: false,
			},
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
					"\nHeap in use: `" + heap.HumanReadable() + "`" +
					"\nStack in use: `" + stack.HumanReadable() + "`",
				Inline: false,
			},
		},
	}
}

// GenerateFeatureRequestEmbed generates the embed which contains the information about a feature request
func GenerateFeatureRequestEmbed(ctx *dgc.Ctx) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Feature Request",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Description",
				Value:  "```" + strings.ReplaceAll(ctx.Arguments.Raw(), "`", "'") + "```",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Requester",
				Value:  "```" + ctx.Event.Author.Username + "#" + ctx.Event.Author.Discriminator + " (" + ctx.Event.Author.ID + ")```",
				Inline: false,
			},
		},
	}
}

// GenerateBugReportEmbed generates the embed which contains the information about a bug report
func GenerateBugReportEmbed(ctx *dgc.Ctx) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Bug Report",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Description",
				Value:  "```" + strings.ReplaceAll(ctx.Arguments.Raw(), "`", "'") + "```",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Reporter",
				Value:  "```" + ctx.Event.Author.Username + "#" + ctx.Event.Author.Discriminator + " (" + ctx.Event.Author.ID + ")```",
				Inline: false,
			},
		},
	}
}

// GenerateGuildSettingsEmbed generates the embed which contains the current settings of a guild
func GenerateGuildSettingsEmbed(guildConfig *guildconfig.GuildConfig) *discordgo.MessageEmbed {
	// Define the command channel string
	commandChannels := "No command channels"
	if len(guildConfig.CommandChannels) > 0 {
		commandChannels = "<#" + strings.Join(guildConfig.CommandChannels, ">, <#") + ">"
	}

	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "Guild Settings Overview",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Command Channel Restriction",
				Value:  "`" + PrettifyBool(guildConfig.ChannelRestriction) + "`",
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Command Channels",
				Value:  commandChannels,
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Hastebin Integration",
				Value:  "`" + PrettifyBool(guildConfig.HastebinIntegration) + "`",
				Inline: false,
			},
		},
	}
}

// GenerateLaTeXResultEmbed generates an embed for LaTeX render results
func GenerateLaTeXResultEmbed(url string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Type:      "rich",
		Title:     "LaTeX Result",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x00ff00,
		Image: &discordgo.MessageEmbedImage{
			URL: url,
		},
	}
}
