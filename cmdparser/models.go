package cmdparser

import "github.com/bwmarrin/discordgo"

// CommandSystem represents a bunch of commands and settings
type CommandSystem struct {
	BotUser     *discordgo.User
	Prefixes    []string
	Commands    map[string]*Command
	PingHandler func(session *discordgo.Session, event *discordgo.MessageCreate)
}

// Command represents a command
type Command struct {
	SubCommands map[string]*Command
	Handler     func(session *discordgo.Session, event *discordgo.MessageCreate, args []string)
}
