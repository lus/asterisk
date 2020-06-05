package middlewares

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/nodes/guilds"
	"github.com/Lukaesebrot/dgc"
)

// InjectGuildObject injects the guild object of the current guild into the custom context objects
func InjectGuildObject(next dgc.ExecutionHandler) dgc.ExecutionHandler {
	return func(ctx *dgc.Ctx) {
		// Retrieve the guild object
		guild, err := guilds.RetrieveCached(ctx.Event.GuildID)
		if err != nil {
			ctx.RespondEmbed(embeds.Error(err.Error()))
			return
		}
		ctx.CustomObjects.Set("guild", guild)
		next(ctx)
	}
}
