package features

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
)

// initializeHashFeature initializes the hash feature
func initializeHashFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'hash' command
	router.RegisterCmd(&dgc.Command{
		Name:        "hash",
		Description: "Hashes the given string using the specified algorithm",
		Usage:       "hash <md5> <string>",
		Example:     "hash md5 Hello, world!",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			{
				Name:        "md5",
				Description: "Hashes the given string using the md5 algorithm",
				Usage:       "hash md5 <string>",
				Example:     "hash md5 Hello, world!",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     hashMD5Command,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     hashCommand,
	})
}

// hashCommand handles the 'hash' command
func hashCommand(ctx *dgc.Ctx) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage(ctx.Command.Usage))
}

// hashMD5Command handles the 'hash md5' command
func hashMD5Command(ctx *dgc.Ctx) {
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
