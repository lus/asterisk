package commands

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// MD5 handles the md5 command
func MD5(ctx *dgc.Ctx) {
	// Validate the arguments
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed(ctx.Command.Usage))
		return
	}

	// Hash the given arguments and respond with the hash
	hash := md5.Sum([]byte(ctx.Arguments.Raw()))
	hashString := hex.EncodeToString(hash[:])
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed(hashString))
}
