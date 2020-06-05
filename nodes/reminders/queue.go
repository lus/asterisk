package reminders

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/bwmarrin/discordgo"
)

// restartQueue gets set to true if the queue should restart
var restartQueue = false

// ScheduleQueue starts the queue that triggers the reminders
func ScheduleQueue(session *discordgo.Session, onTrigger func(reminder *Reminder)) {
	// Retrieve the next reminder
	reminder, err := Next()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			for !restartQueue {
				time.Sleep(time.Millisecond)
			}
			ScheduleQueue(session, onTrigger)
			return
		}
		log.Println("[ERR] ", err.Error())
		ScheduleQueue(session, onTrigger)
		return
	}

	// Wait until the reminder exceeds or the queue is instructed to restart
	for reminder.Exceeds > time.Now().Unix() {
		if restartQueue {
			restartQueue = false
			ScheduleQueue(session, onTrigger)
			return
		}
		time.Sleep(time.Millisecond)
	}

	// Trigger and delete the reminder
	onTrigger(reminder)
	err = reminder.Delete()
	if err != nil {
		log.Println("[ERR] ", err.Error())
	}
	ScheduleQueue(session, onTrigger)
}

// RestartQueue instructs the queue to restart
func RestartQueue() {
	restartQueue = true
}
