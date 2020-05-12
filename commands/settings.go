package commands

import (
	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/bwmarrin/discordgo"
)

// Define the usage of this command
var settingsUsage = "$settings <isRestricted | setRestricted <bool> | addCommandChannel <channel mention> | removeCommandChannel <channel mention>>"

// Settings handles the settings command
func Settings() func(*discordgo.Session, *discordgo.MessageCreate, []string, *guildconfig.GuildConfig) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string, guildConfig *guildconfig.GuildConfig) {
		// TODO: Implement settings command structure
	}
}
