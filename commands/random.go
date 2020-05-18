package commands

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Random handles the random command
func Random(ctx *dgc.Ctx) {
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed(ctx.Command.Usage))
}

// RandomBool handles the random bool command
func RandomBool(ctx *dgc.Ctx) {
	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Respond with the generated random boolean
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed(strconv.FormatBool(rand.Intn(2) == 0)))
}

// RandomNumber handles the random number command
func RandomNumber(ctx *dgc.Ctx) {
	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Define the random number
	number := rand.Int()
	if ctx.Arguments.Amount() > 0 {
		valid, generated := utils.GenerateFromInterval(ctx.Arguments.Raw())
		if !valid {
			ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("The interval you specified is invalid."))
			return
		}
		number = generated
	}

	// Respond with the generated random number
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed(strconv.Itoa(number)))
}

// RandomString handles the random string command
func RandomString(ctx *dgc.Ctx) {
	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Validate the argument length
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("You need to specify a length."))
		return
	}

	// Parse the string length
	length, err := ctx.Arguments.Get(0).AsInt()
	if err != nil || length <= 0 || length > 100 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("The length parameter has to be a number > 0 and < 100."))
		return
	}

	// Generate the random string
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteArray := make([]byte, length)
	for i := range byteArray {
		byteArray[i] = characters[rand.Intn(len(characters))]
	}

	// Respond with the generated random string
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed(string(byteArray)))
}

// RandomChoice handles the random choice command
func RandomChoice(ctx *dgc.Ctx) {
	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Validate the argument length
	if ctx.Arguments.Amount() < 2 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed("You need to specify at least 2 options."))
		return
	}

	// Make a random choice
	option := ctx.Arguments.Get(rand.Intn(ctx.Arguments.Amount())).Raw()

	// Respond with the random piked choice
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed(option))
}
