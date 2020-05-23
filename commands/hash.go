package commands

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
)

// Hash handles the hash command
func Hash(ctx *dgc.Ctx) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage(ctx.Command.Usage))
}

// HashMD5 handles the hash md5 command
func HashMD5(ctx *dgc.Ctx) {
	// Validate the arguments
	raw := ctx.Arguments.Raw()
	if raw == "" {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage("You need to specify a string you want to hash."))
		return
	}

	// Hash the given string
	hashBytes := md5.Sum([]byte(raw))
	hashString := hex.EncodeToString(hashBytes[:])

	// Respond with the hash
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Success(hashString))
}
