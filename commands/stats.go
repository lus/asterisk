package commands

import (
	"github.com/Lukaesebrot/dgc"

	"github.com/Lukaesebrot/asterisk/utils"
)

// Stats handles the stats command
func Stats(ctx *dgc.Ctx) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateStatsEmbed(ctx.Session))
}
