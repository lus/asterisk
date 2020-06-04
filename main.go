package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Lukaesebrot/asterisk/features"
	"github.com/Lukaesebrot/asterisk/utils"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/reminders"
	"github.com/Lukaesebrot/asterisk/users"

	"github.com/Lukaesebrot/asterisk/static"

	"github.com/Lukaesebrot/asterisk/config"
	"github.com/Lukaesebrot/asterisk/database"
	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("Starting this Asterisk instance...")

	// Initialize the configuration
	log.Println("Loading the bot configuration...")
	config.Load()
	log.Println("Successfully loaded the bot configuration.")

	// Connect to the MongoDB host
	log.Println("Connecting to the specified MongoDB server...")
	err := database.Connect()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to the specified MongoDB server.")

	// Initialize the Discord session
	log.Println("Establishing the Discord connection...")
	session, err := discordgo.New("Bot " + config.CurrentConfig.BotToken)
	if err != nil {
		panic(err)
	}
	session.AddHandler(func(session *discordgo.Session, event *discordgo.Ready) {
		go utils.AnimatePresence(session)
	})
	err = session.Open()
	if err != nil {
		panic(err)
	}
	static.Self = session.State.User
	static.StartupTime = time.Now()
	log.Println("Successfully established the Discord connection.")

	// Initialize all features
	log.Println("Initializing all features...")
	features.Initialize(session)
	log.Println("Successfully initialized all features.")

	// Schedule the reminder queue
	log.Println("Scheduling the reminder queue...")
	go reminders.ScheduleQueue(session, func(reminder *reminders.Reminder) {
		session.ChannelMessageSendComplex(reminder.ChannelID, &discordgo.MessageSend{
			Content: fmt.Sprintf("<@%s>", reminder.UserID),
			Embed:   embeds.Reminder(reminder),
		})
	})
	log.Println("Successfully scheduled the reminder queue.")

	// Make the specified initial user a bot admin
	log.Println("Updating the initial user permissions...")
	user, err := users.RetrieveCached(config.CurrentConfig.InitialAdminID)
	if err != nil {
		panic(err)
	}
	user.AssignFlag(users.FlagAdministrator)
	err = user.Update()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully updated the initial user permissions.")

	// Wait for the program to exit
	log.Println("Successfully started this Asterisk instance.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Gracefully close the Discord session
	log.Println("Stopping this Asterisk instance...")
	session.Close()
	database.CurrentClient.Disconnect(context.TODO())
}
