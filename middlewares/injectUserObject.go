package middlewares

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/nodes/users"
	"github.com/Lukaesebrot/dgc"
)

// InjectUserObject injects the user object of the message author into the custom context objects
func InjectUserObject(next dgc.ExecutionHandler) dgc.ExecutionHandler {
	return func(ctx *dgc.Ctx) {
		// Retrieve the user object
		user, err := users.RetrieveCached(ctx.Event.Author.ID)
		if err != nil {
			ctx.RespondEmbed(embeds.Error(err.Error()))
			return
		}
		ctx.CustomObjects.Set("user", user)
		next(ctx)
	}
}
