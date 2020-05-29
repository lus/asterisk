package features

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
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
			{
				Name:        "sha256",
				Description: "Hashes the given string using the sha256 algorithm",
				Usage:       "hash sha256 <string>",
				Example:     "hash sha256 Hello, world!",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     hashSHA256Command,
			},
			{
				Name:        "sha512",
				Description: "Hashes the given string using the sha512 algorithm",
				Usage:       "hash sha512 <string>",
				Example:     "hash sha512 Hello, world!",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     hashSHA512Command,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     hashCommand,
	})
}

// hashCommand handles the 'hash' command
func hashCommand(ctx *dgc.Ctx) {
	ctx.RespondEmbed(embeds.InvalidUsage(ctx.Command.Usage))
}

// hashMD5Command handles the 'hash md5' command
func hashMD5Command(ctx *dgc.Ctx) {
	// Validate the arguments
	raw := ctx.Arguments.Raw()
	if raw == "" {
		ctx.RespondEmbed(embeds.InvalidUsage("You need to specify a string you want to hash."))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Hash the given string
	hashBytes := md5.Sum([]byte(raw))
	hashString := hex.EncodeToString(hashBytes[:])

	// Respond with the hash
	ctx.RespondEmbed(embeds.Success(hashString))
}

// hashSHA256Command handles the 'hash sha256' command
func hashSHA256Command(ctx *dgc.Ctx) {
	// Validate the arguments
	raw := ctx.Arguments.Raw()
	if raw == "" {
		ctx.RespondEmbed(embeds.InvalidUsage("You need to specify a string you want to hash."))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Hash the given string
	hashBytes := sha256.Sum256([]byte(raw))
	hashString := hex.EncodeToString(hashBytes[:])

	// Respond with the hash
	ctx.RespondEmbed(embeds.Success(hashString))
}

// hashSHA512Command handles the 'hash sha512' command
func hashSHA512Command(ctx *dgc.Ctx) {
	// Validate the arguments
	raw := ctx.Arguments.Raw()
	if raw == "" {
		ctx.RespondEmbed(embeds.InvalidUsage("You need to specify a string you want to hash."))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Hash the given string
	hashBytes := sha512.Sum512([]byte(raw))
	hashString := hex.EncodeToString(hashBytes[:])

	// Respond with the hash
	ctx.RespondEmbed(embeds.Success(hashString))
}
