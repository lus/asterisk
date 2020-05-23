package middlewares

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/users"
	"github.com/Lukaesebrot/dgc"
)

// InjectUserObject injects the user object of the message author into the custom context objects
func InjectUserObject(ctx *dgc.Ctx) bool {
	// Retrieve the user object
	user, err := users.RetrieveCached(ctx.Event.Author.ID)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return false
	}
	ctx.CustomObjects.Set("user", user)
	return true
}
