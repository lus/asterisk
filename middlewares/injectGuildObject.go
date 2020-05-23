package middlewares

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/guilds"
	"github.com/Lukaesebrot/dgc"
)

// InjectGuildObject injects the guild object of the current guild into the custom context objects
func InjectGuildObject(ctx *dgc.Ctx) bool {
	// Retrieve the guild object
	guild, err := guilds.RetrieveCached(ctx.Event.GuildID)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return false
	}
	ctx.CustomObjects.Set("guild", guild)
	return true
}
