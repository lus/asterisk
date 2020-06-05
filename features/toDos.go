package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/nodes/todos"
	"github.com/Lukaesebrot/dgc"
	"go.mongodb.org/mongo-driver/mongo"
)

// initializeToDoFeature initializes the toDo feature
func initializeToDoFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'toDo' command
	router.RegisterCmd(&dgc.Command{
		Name:        "toDo",
		Aliases:     []string{"toDos", "task", "tasks"},
		Description: "Lists your current ToDo objects or manages them",
		Usage:       "toDo [create <content> | done <id>]",
		Example:     "toDo create Fix some bugs.",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			{
				Name:        "create",
				Aliases:     []string{"c"},
				Description: "Creates a new ToDo object",
				Usage:       "toDo create <message>",
				Example:     "toDo create Fix some bugs.",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     toDoCreateCommand,
			},
			{
				Name:        "done",
				Aliases:     []string{"delete", "d", "rm"},
				Description: "Marks the ToDo item with the given ID as done",
				Usage:       "toDo done <id>",
				Example:     "toDO done 1",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     toDoDoneCommand,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     toDoCommand,
	})
}

// toDoCommand handles the 'toDo' command
func toDoCommand(ctx *dgc.Ctx) {
	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Retrieve the users ToDos
	userToDos, err := todos.GetAll(ctx.Event.Author.ID)
	if err != nil {
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}

	// Respond with the list ToDos embed
	ctx.RespondEmbed(embeds.ToDos(userToDos))
}

// toDoCreateCommand handles the 'toDo create' command
func toDoCreateCommand(ctx *dgc.Ctx) {
	// Validate the argument length
	if ctx.Arguments.Amount() < 1 {
		ctx.RespondEmbed(embeds.InvalidUsage(ctx.Command.Usage))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Create the ToDo object
	_, err := todos.Create(ctx.Event.Author.ID, ctx.Arguments.Raw())
	if err != nil {
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}

	// Respond with a success embed
	ctx.RespondEmbed(embeds.Success("Your ToDo object has been created."))
}

// toDoDoneCommand handles the 'toDo done' command
func toDoDoneCommand(ctx *dgc.Ctx) {
	// Parse the arguments into an integer
	id, err := ctx.Arguments.AsSingle().AsInt64()
	if err != nil {
		ctx.RespondEmbed(embeds.InvalidUsage(ctx.Command.Usage))
		return
	}

	// Retrieve the ToDo object
	toDo, err := todos.Get(ctx.Event.Author.ID, id-1)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.RespondEmbed(embeds.Error("There is no ToDO object with the specified ID."))
			return
		}
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}

	// Check the rate limiter
	if !ctx.Command.NotifyRateLimiter(ctx) {
		return
	}

	// Delete the ToDo object
	err = toDo.Delete()
	if err != nil {
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}

	// Respond with a success embed
	ctx.RespondEmbed(embeds.Success("Your ToDo object has been marked as done."))
}
