package middlewares

import (
	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// InjectGuildConfig injects the guild configuration into the custom context objects
func InjectGuildConfig(ctx *dgc.Ctx) bool {
	// Retrieve the guild configuration
	guildConfig, err := guildconfig.RetrieveCached(ctx.Event.GuildID)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return false
	}
	ctx.CustomObjects.Set("guildConfig", guildConfig)
	return true
}
