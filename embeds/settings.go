package embeds

import (
	"fmt"
	"time"

	"github.com/Lukaesebrot/asterisk/guilds"
	"github.com/bwmarrin/discordgo"
)

// Settings generates a guild settings embed
func Settings(guild *guilds.Guild) *discordgo.MessageEmbed {
	// Define the command channels string
	commandChannels := "`*`"
	if len(guild.Settings.CommandChannels) > 0 {
		commandChannels = ""
		for index, channelID := range guild.Settings.CommandChannels {
			commandChannels += fmt.Sprintf("<#%s> (`%s`)", channelID, channelID)
			if len(guild.Settings.CommandChannels) > index+1 {
				commandChannels += ", "
			}
		}
	}

	return &discordgo.MessageEmbed{
		Title:     "Guild Settings",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Command channels",
				Value: commandChannels,
			},
		},
	}
}
