package commands

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/bwmarrin/discordgo"
)

// define the usage of this command
var usage = "$random <bool | number | string | choice>"

// Random handles the random command
func Random() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed(usage))
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

		// Define the random number
		number := rand.Int()
		if len(args) > 0 {
			valid, generated := utils.FormatInterval(strings.Join(args, " "))
			if !valid {
				_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("The interval you specified is invalid"))
				if err != nil {
					log.Println("[ERR] " + err.Error())
				}
				return
			}
			number = generated
		}

		// Respond with the generated random number
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateRandomOutputEmbed(strconv.Itoa(number)))
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}

// RandomString handles the random string command
func RandomString() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		// Seed the random generator
		rand.Seed(time.Now().UnixNano())

		// Validate the argument length
		if len(args) == 0 {
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("You need to specify a length"))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			return
		}

		// Parse the string length
		length, err := strconv.Atoi(args[0])
		if err != nil || length <= 0 {
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("The length parameter has to be a number > 0"))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			return
		}

		// Generate the random string
		characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		byteArray := make([]byte, length)
		for i := range byteArray {
			byteArray[i] = characters[rand.Intn(len(characters))]
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
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("You need to specify at least 2 options"))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			return
		}

		// Make a random choice
		option := args[rand.Intn(len(args))]

		// Respond with the random piked choice
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateRandomOutputEmbed(option))
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}
