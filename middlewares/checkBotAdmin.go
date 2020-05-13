package middlewares

import (
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// CheckBotAdmin checks if the executor is a bot admin
func CheckBotAdmin(ctx *dgc.Ctx) bool {
	if !utils.IsBotAdmin(ctx.Event.Author.ID) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInsufficientPermissionsEmbed("You need to be a bot admin to use this command."))
		return false
	}
	return true
}
