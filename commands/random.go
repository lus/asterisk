package commands

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/bwmarrin/discordgo"
)

// Random handles the random command
func Random() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("$random <bool | number [max number (int)] | string <length (int)> | choice <options... (min 2)>>"))
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}

// RandomBool handles the random bool command
func RandomBool() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		// Seed the random generator
		rand.Seed(time.Now().UnixNano())

		// Respond with the generated random boolean
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateRandomOutputEmbed(strconv.FormatBool(rand.Intn(2) == 0)))
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}

// RandomNumber handles the random number command
func RandomNumber() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		// Seed the random generator
		rand.Seed(time.Now().UnixNano())

		// Define the maximum number
		max := -1
		if len(args) != 0 {
			rawMax, err := strconv.Atoi(args[0])
			if err != nil {
				_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("$random <bool | number [max number (int)] | string <length (int)> | choice <options... (min 2)>>"))
				if err != nil {
					log.Println("[ERR] " + err.Error())
				}
				return
			}
			max = rawMax
		}

		// Generate the random number
		rnd := rand.Int()
		if max > 0 {
			rnd = rand.Intn(max)
		}

		// Respond with the generated random number
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateRandomOutputEmbed(strconv.Itoa(rnd)))
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}

// stringCharacters holds the characters that may be part of a string
const stringCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString handles the random string command
func RandomString() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		// Seed the random generator
		rand.Seed(time.Now().UnixNano())

		// Validate the argument length
		if len(args) == 0 {
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("$random <bool | number [max number (int)] | string <length (int)> | choice <options... (min 2)>>"))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			return
		}

		// Define the length
		length, err := strconv.Atoi(args[0])
		if err != nil {
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("$random <bool | number [max number (int)] | string <length (int)> | choice <options... (min 2)>>"))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			return
		}

		// Generate the random string
		byteArray := make([]byte, length)
		for i := range byteArray {
			byteArray[i] = stringCharacters[rand.Intn(len(stringCharacters))]
		}

		// Respond with the generated random string
		_, err = session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateRandomOutputEmbed(string(byteArray)))
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}

// RandomChoice handles the random choice command
func RandomChoice() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		// Seed the random generator
		rand.Seed(time.Now().UnixNano())

		// Validate the argument length
		if len(args) < 2 {
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("$random <bool | number [max number (int)] | string <length (int)> | choice <options... (min 2)>>"))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			return
		}

		// Choose a random option
		choice := args[rand.Intn(len(args))]

		// Respond with the generated random choice
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateRandomOutputEmbed(choice))
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}
