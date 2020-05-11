package cmdparser

import (
	"log"
	"strings"

	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/bwmarrin/discordgo"
)

// Handler returns the handler of a command system
func (system *CommandSystem) Handler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate) {
		// Define useful variables
		message := event.Message
		content := message.Content

		// Check if the message only contains the bot ping
		if content == "<@"+system.BotUser.ID+">" {
			system.PingHandler(session, event)
			return
		}

		// Check if the message starts with one of the defined prefixes
		content, valid := utils.StringHasPrefix(content, system.Prefixes, true)
		if !valid {
			return
		}

		// Retrieve the corresponding guild configuration
		guildConfig, err := guildconfig.Retrieve(message.GuildID)
		if err != nil {
			_, err := session.ChannelMessageSendEmbed(message.ChannelID, utils.GenerateInternalErrorEmbed(err.Error()))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			return
		}

		// Check if the current channel is a valid command channel
		if !guildConfig.ValidateChannel(message.ChannelID) {
			return
		}

		// Trigger the command if it exists
		split := strings.Split(content, " ")
		command := system.Commands[strings.ToLower(split[0])]
		if command != nil {
			command.Trigger(session, event, split[1:])
		}
	}
}

// Trigger triggers a command
func (command *Command) Trigger(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
	// Handle this command if no arguments were specified
	if len(args) == 0 {
		command.Handler(session, event, args)
		return
	}

	// Trigger possible subcommands
	subCommand := command.SubCommands[strings.ToLower(args[0])]
	if subCommand != nil {
		subCommand.Trigger(session, event, args[1:])
		return
	}
	command.Handler(session, event, args)
}
