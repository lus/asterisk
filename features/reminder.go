package features

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/reminders"
	"github.com/Lukaesebrot/dgc"
	"go.mongodb.org/mongo-driver/mongo"
)

// initializeReminderFeature initializes the reminder feature
func initializeReminderFeature(router *dgc.Router, rateLimiter dgc.RateLimiter) {
	// Register the 'reminder' command
	router.RegisterCmd(&dgc.Command{
		Name:        "reminder",
		Aliases:     []string{"remind", "reminders"},
		Description: "Lists your current reminders or manages them",
		Usage:       "reminder [create <duration> <message> | delete <id>]",
		Example:     "reminder create 2h Hello, world!",
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{
			{
				Name:        "create",
				Aliases:     []string{"c"},
				Description: "Creates a new reminder",
				Usage:       "reminder create <duration> <message>",
				Example:     "reminder create 2h Hello, world!",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     reminderCreateCommand,
			},
			{
				Name:        "delete",
				Aliases:     []string{"d", "rm"},
				Description: "Deletes the reminder with the given ID",
				Usage:       "reminder delete <id>",
				Example:     "reminder delete 1",
				IgnoreCase:  true,
				RateLimiter: rateLimiter,
				Handler:     reminderDeleteCommand,
			},
		},
		RateLimiter: rateLimiter,
		Handler:     reminderCommand,
	})
}

// reminderCommand handles the 'reminder' command
func reminderCommand(ctx *dgc.Ctx) {
	// Retrieve the users reminders
	userReminders, err := reminders.GetAll(ctx.Event.Author.ID)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}

	// Respond with the list reminders embed
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Reminders(userReminders))
}

// reminderCreateCommand handles the 'reminder create' command
func reminderCreateCommand(ctx *dgc.Ctx) {
	// Validate the argument length
	if ctx.Arguments.Amount() < 2 {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage(ctx.Command.Usage))
		return
	}

	// Parse the first argument into a duration
	duration, err := ctx.Arguments.Get(0).AsDuration()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage(ctx.Command.Usage))
		return
	}

	// Use the remaining arguments as the reminder message
	ctx.Arguments.Remove(0)
	message := ctx.Arguments.Raw()

	// Create the reminder and restart the reminder queue
	_, err = reminders.Create(ctx.Event.Author.ID, ctx.Event.ChannelID, duration, message)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}
	reminders.RestartQueue()

	// Respond with a success embed
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Success("Your reminder has been created."))
}

// reminderDeleteCommand handles the 'reminder delete' command
func reminderDeleteCommand(ctx *dgc.Ctx) {
	// Parse the arguments into an integer
	id, err := ctx.Arguments.AsSingle().AsInt64()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.InvalidUsage(ctx.Command.Usage))
		return
	}

	// Retrieve the reminder
	reminder, err := reminders.Get(ctx.Event.Author.ID, id-1)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("There is no reminder with the specified ID."))
			return
		}
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}

	// Delete the reminder and restart the reminder queue
	err = reminder.Delete()
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}
	reminders.RestartQueue()

	// Respond with a success embed
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Success("Your reminder has been deleted."))
}
