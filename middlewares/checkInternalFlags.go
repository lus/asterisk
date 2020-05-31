package middlewares

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/users"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// CheckInternalFlags checks if the current user has got the given internal flag(s)
func CheckInternalFlags(flag, format string, flags ...int) dgc.Middleware {
	return func(next dgc.ExecutionHandler) dgc.ExecutionHandler {
		return func(ctx *dgc.Ctx) {
			// Check if the command has got the specified flag
			if !utils.StringArrayContains(ctx.Command.Flags, flag) {
				next(ctx)
				return
			}

			// Retrieve the current user
			user := ctx.CustomObjects.MustGet("user").(*users.User)

			// Check if the user has the required flags
			hasFlags := user.HasFlag(flags...)
			if !hasFlags {
				ctx.RespondEmbed(embeds.InsufficientPermissions("You need to have the internal '" + format + "' flag(s)."))
				return
			}
			next(ctx)
		}
	}
}
