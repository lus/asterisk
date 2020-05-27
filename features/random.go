package features

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// initializeRandomFeature initializes the random feature
func initializeRandomFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'random' command
	router.RegisterCmd(&dgc.Command{
		Name:        "random",
		Description: "Generates a random boolean, number, string or choice",
		Usage:       "random <bool | number <interval> | string <length> | choice <options...>>",
		Example:     "random number [0,5]",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			{
				Name:        "bool",
				Description: "Generates a random boolean",
				Usage:       "random bool",
				Example:     "random bool",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     randomBoolCommand,
			},
			{
				Name:        "number",
				Description: "Generates a random number respecting the given interval",
				Usage:       "random number <interval>",
				Example:     "random number [0,5]",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     randomNumberCommand,
			},
			{
				Name:        "string",
				Description: "Generates a random string with the given length",
				Usage:       "random string <length>",
				Example:     "random string 32",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     randomStringCommand,
			},
			{
				Name:        "choice",
				Description: "Generates a random choice",
				Usage:       "random choice <options...>",
				Example:     "random choice \"Coice one\" \"Choice two\"",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     randomChoiceCommand,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     nil,
	})
}

// randomCommand handles the 'random' command
func randomCommand(ctx *dgc.Ctx) {
	ctx.RespondEmbed(embeds.InvalidUsage(ctx.Command.Usage))
}

// randomBoolCommand handles the 'random bool' command
func randomBoolCommand(ctx *dgc.Ctx) {
	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Respond with the generated random boolean
	ctx.RespondEmbed(embeds.Success(strconv.FormatBool(rand.Intn(2) == 0)))
}

// randomNumberCommand handles the 'random number' command
func randomNumberCommand(ctx *dgc.Ctx) {
	// Define the random number
	number := rand.Int()
	if ctx.Arguments.Amount() > 0 {
		valid, generated := utils.GenerateFromInterval(ctx.Arguments.Raw())
		if !valid {
			ctx.RespondEmbed(embeds.InvalidUsage("The interval you specified is invalid."))
			return
		}
		number = generated
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Respond with the generated random number
	ctx.RespondEmbed(embeds.Success(strconv.Itoa(number)))
}

// randomStringCommand handles the 'random string' command
func randomStringCommand(ctx *dgc.Ctx) {
	// Validate the argument length
	if ctx.Arguments.Amount() == 0 {
		ctx.RespondEmbed(embeds.InvalidUsage("You need to specify a length."))
		return
	}

	// Parse the string length
	length, err := ctx.Arguments.Get(0).AsInt()
	if err != nil || length <= 0 || length > 100 {
		ctx.RespondEmbed(embeds.InvalidUsage("The length parameter has to be a number > 0 and <= 100."))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Generate the random string
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteArray := make([]byte, length)
	for i := range byteArray {
		byteArray[i] = characters[rand.Intn(len(characters))]
	}

	// Respond with the generated random string
	ctx.RespondEmbed(embeds.Success(string(byteArray)))
}

// randomChoiceCommand handles the 'random choice' command
func randomChoiceCommand(ctx *dgc.Ctx) {
	// Validate the argument length
	if ctx.Arguments.Amount() < 2 {
		ctx.RespondEmbed(embeds.InvalidUsage("You need to specify at least 2 options."))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Make a random choice
	option := ctx.Arguments.Get(rand.Intn(ctx.Arguments.Amount())).Raw()

	// Respond with the random piked choice
	ctx.RespondEmbed(embeds.Success(option))
}
