package embeds

import (
	"fmt"
	"time"

	"github.com/Lukaesebrot/asterisk/reminders"
	"github.com/bwmarrin/discordgo"
)

// Reminder generates a reminder embed
func Reminder(reminder *reminders.Reminder) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "Reminder",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "User",
				Value: fmt.Sprintf("<@%s> (`%s`)", reminder.UserID, reminder.UserID),
			},
			{
				Name:  "Message",
				Value: "```" + reminder.Message + "```",
			},
		},
	}
}

// Reminders generates a reminder list embed
func Reminders(userReminders []reminders.Reminder) *discordgo.MessageEmbed {
	// Define the reminders string
	remindersString := "No scheduled reminders"
	if len(userReminders) > 0 {
		remindersString = ""
		counter := 1
		for _, reminder := range userReminders {
			if counter != 1 {
				remindersString += "\n"
			}
			remindersString += fmt.Sprintf("`%d.` In: `%s`", counter, time.Until(time.Unix(reminder.Exceeds, 0)).Round(time.Second).String())
			counter++
		}
	}

	return &discordgo.MessageEmbed{
		Title:     "Reminder List",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Reminders",
				Value: remindersString,
			},
		},
	}
}
