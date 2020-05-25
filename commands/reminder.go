package commands

import (
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/reminders"
	"github.com/Lukaesebrot/dgc"
	"go.mongodb.org/mongo-driver/mongo"
)

// Reminder handles the reminder command
func Reminder(ctx *dgc.Ctx) {
	// Retrieve the users reminders
	userReminders, err := reminders.GetAll(ctx.Event.Author.ID)
	if err != nil {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error(err.Error()))
		return
	}

	// Respond with the list reminders embed
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Reminders(userReminders))
}

// ReminderCreate handles the reminder create command
func ReminderCreate(ctx *dgc.Ctx) {
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

// ReminderDelete handles the reminder delete command
func ReminderDelete(ctx *dgc.Ctx) {
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
