package utils

import (
	"github.com/Lukaesebrot/asterisk/config"
	"github.com/bwmarrin/discordgo"
)

// IsBotAdmin checks wheter or not the given user ID is part of the list of bot admins
func IsBotAdmin(userID string) bool {
	return StringArrayContains(config.CurrentConfig.BotAdmins, userID)
}

// HasPermission checks whether or not the given user has the given permission
func HasPermission(session *discordgo.Session, guildID string, userID string, permission int) (bool, error) {
	// Check if the user is the guild owner
	guild, err := session.State.Guild(guildID)
	if err != nil {
		guild, err = session.Guild(guildID)
		if err != nil {
			return false, err
		}
	}
	if guild.OwnerID == userID {
		return true, nil
	}

	// Retrieve the member object
	member, err := session.State.Member(guildID, userID)
	if err != nil {
		member, err = session.GuildMember(guildID, userID)
		if err != nil {
			return false, err
		}
	}

	// Check every role for the required permission
	for _, roleID := range member.Roles {
		role, err := session.State.Role(guildID, roleID)
		if err != nil {
			return false, err
		}
		if role.Permissions&permission != 0 {
			return true, nil
		}
	}
	return false, nil
}
