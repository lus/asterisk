package middlewares

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// CheckGuildPermissions checks if the current executor has got the given guild permission(s)
func CheckGuildPermissions(flag, format string, permissions ...int) dgc.Middleware {
	return func(next dgc.ExecutionHandler) dgc.ExecutionHandler {
		return func(ctx *dgc.Ctx) {
			// Check if the command has got the specified flag
			if !utils.StringArrayContains(ctx.Command.Flags, flag) {
				next(ctx)
				return
			}

			// Check if the executer has the required permissions
			hasPermissions := true
			for _, permission := range permissions {
				hasPermission, _ := hasPermission(ctx.Session, ctx.Event.GuildID, ctx.Event.Author.ID, permission)
				if !hasPermission {
					hasPermissions = false
					break
				}
			}
			if !hasPermissions {
				ctx.RespondEmbed(embeds.InsufficientPermissions("You need to have the guild-related '" + format + "' permission(s)."))
				return
			}
			next(ctx)
		}
	}
}

// hasPermission checks whether or not the given user has the given permission
func hasPermission(session *discordgo.Session, guildID string, userID string, permission int) (bool, error) {
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
