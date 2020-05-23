package middlewares

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/users"
	"github.com/Lukaesebrot/dgc"
)

// CheckInternalPermissions checks if the current user has got the given internal permission(s)
func CheckInternalPermissions(format string, permissions ...users.Permission) func(*dgc.Ctx) bool {
	return func(ctx *dgc.Ctx) bool {
		// Retrieve the current user
		user := ctx.CustomObjects.MustGet("user").(*users.User)

		// Check if the user has the required permissions
		hasPermission := user.HasPermission(permissions...)
		if !hasPermission {
			ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InsufficientPermissions("You need to have the internal '"+format+"' permission(s)."))
		}
		return hasPermission
	}
}
