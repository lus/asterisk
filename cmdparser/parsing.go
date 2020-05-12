package cmdparser

import (
	"log"
	"strings"

	"github.com/Lukaesebrot/asterisk/guildconfig"
	"github.com/Lukaesebrot/asterisk/static"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/bwmarrin/discordgo"
)

// Handler returns the handler of a command system
func (system *CommandSystem) Handler() func(*discordgo.Session, *discordgo.MessageCreate) {
	// Initialize the help command
	system.Commands["help"] = &Command{
		Description: "Lists all available commands or shows detailed information about a single command",
		Handler: func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
			// Check the argument length
			if len(args) > 0 {
				commandName := strings.ToLower(args[0])
				if system.Commands[commandName] == nil {
					_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateInvalidUsageEmbed("The command '"+commandName+"' does not exist"))
					if err != nil {
						log.Println("[ERR] " + err.Error())
					}
					return
				}
				_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateSingleHelpEmbed(commandName, system.Commands[commandName].Description))
				if err != nil {
					log.Println("[ERR] " + err.Error())
				}
				return
			}

			// Fetch all command names
			commandNames := make([]string, len(system.Commands))
			index := 0
			for commandName := range system.Commands {
				commandNames[index] = commandName
				index++
			}

			// Respond with the help embed
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateHelpEmbed(commandNames))
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
		},
	}

	return func(session *discordgo.Session, event *discordgo.MessageCreate) {
		// Define useful variables
		message := event.Message
		content := message.Content

		// Check if the message only contains the bot ping
		if content == "<@!"+static.Self.ID+">" {
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
