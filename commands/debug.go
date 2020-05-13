package commands

import (
	"reflect"
	"strings"

	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

// Debug handles the debug command
func Debug(ctx *dgc.Ctx) {
	// Check if the executor is a bot admin
	if !utils.IsBotAdmin(ctx.Event.Author.ID) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInsufficientPermissionsEmbed("You need to be a bot admin to use this command."))
		return
	}

	// Validate the arguments
	if ctx.Arguments.Amount() == 0 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed(ctx.Command.Usage))
		return
	}

	// Define the evaluation string and create the interpreter
	evaluationString := strings.ReplaceAll(ctx.Arguments.Raw(), "```", "")
	interpreter := interp.New(interp.Options{})

	// Inject the custom variables
	custom := make(map[string]map[string]reflect.Value)
	custom["main"] = map[string]reflect.Value{
		"ctx": reflect.ValueOf(ctx),
	}
	interpreter.Use(stdlib.Symbols)
	interpreter.Use(custom)
	_, err := interpreter.Eval("import (\n. \"main\"\n\"fmt\"\n)")
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}

	// Evaluate the given string and output the result
	_, err = interpreter.Eval(evaluationString)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateErrorEmbed(err.Error()))
		return
	}
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("Evaluation succeeded."))
}
