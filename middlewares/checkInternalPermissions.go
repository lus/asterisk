package middlewares

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/users"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// CheckInternalPermissions checks if the current user has got the given internal permission(s)
func CheckInternalPermissions(flag, format string, permissions ...users.Permission) dgc.Middleware {
	return func(next dgc.ExecutionHandler) dgc.ExecutionHandler {
		return func(ctx *dgc.Ctx) {
			// Check if the command has got the specified flag
			if !utils.StringArrayContains(ctx.Command.Flags, flag) {
				next(ctx)
				return
			}

			// Retrieve the current user
			user := ctx.CustomObjects.MustGet("user").(*users.User)

			// Check if the user has the required permissions
			hasPermission := user.HasPermission(permissions...)
			if !hasPermission {
				ctx.RespondEmbed(embeds.InsufficientPermissions("You need to have the internal '" + format + "' permission(s)."))
				return
			}
			next(ctx)
		}
	}
}
