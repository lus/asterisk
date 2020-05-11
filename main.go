package main

import (
	"log"

	"github.com/Lukaesebrot/asterisk/config"
)

func main() {
	log.Println("Starting this Asterisk instance...")

	// Initialize the configuration
	log.Println("Loading the bot configuration...")
	err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully loaded the bot configuration.")

	// TODO: Implement MongoDB integration

	// TODO: Implement command system

	// TODO: Implement console command system

	log.Println("Waiting for console commands. Type 'help' for help.")
}
