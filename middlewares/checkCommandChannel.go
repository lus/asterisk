package middlewares

import (
	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// CheckCommandChannel checks if the current channel is a valid command channel
func CheckCommandChannel(ctx *dgc.Ctx) bool {
	guildConfig := ctx.CustomObjects.MustGet("guildConfig").(*guildconfig.GuildConfig)
	if guildConfig.ChannelRestriction && !utils.StringArrayContains(guildConfig.CommandChannels, ctx.Event.ChannelID) {
		isAdmin, _ := hasPermission(ctx.Session, ctx.Event.GuildID, ctx.Event.Author.ID, discordgo.PermissionAdministrator)
		return isAdmin
	}
	return true
}
