package middlewares

import (
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// CheckGuildAdmin checks if the executor is a guild admin
func CheckGuildAdmin(ctx *dgc.Ctx) bool {
	isAdmin, _ := utils.HasPermission(ctx.Session, ctx.Event.GuildID, ctx.Event.Author.ID, discordgo.PermissionAdministrator)
	if !isAdmin {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInsufficientPermissionsEmbed("You need to be a guild admin to use this command."))
		return false
	}
	return true
}
