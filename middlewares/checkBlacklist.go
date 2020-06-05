package middlewares

import (
	"github.com/Lukaesebrot/asterisk/nodes/users"
	"github.com/Lukaesebrot/dgc"
)

// CheckBlacklist checks if the current user is blacklisted
func CheckBlacklist(next dgc.ExecutionHandler) dgc.ExecutionHandler {
	return func(ctx *dgc.Ctx) {
		// Retrieve the current user
		user := ctx.CustomObjects.MustGet("user").(*users.User)

		// Check if the user is a staff member
		if user.HasFlag(users.FlagAdministrator) || user.HasFlag(users.FlagModerator) {
			next(ctx)
			return
		}

		// Check if the user has the 'blacklisted' flag
		isBlacklisted := user.HasFlag(users.FlagBlacklisted)
		if isBlacklisted {
			return
		}
		next(ctx)
	}
}
