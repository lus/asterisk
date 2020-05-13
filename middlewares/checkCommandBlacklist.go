package middlewares

import (
	"github.com/Lukaesebrot/asterisk/static"
	"github.com/Lukaesebrot/dgc"
)

// CheckCommandBlacklist checks if the executor is blacklisted
func CheckCommandBlacklist(ctx *dgc.Ctx) bool {
	return !static.Blacklist[ctx.Event.Author.ID]
}
