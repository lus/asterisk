package cmdparser

import "github.com/bwmarrin/discordgo"

// CommandSystem represents a bunch of commands and settings
type CommandSystem struct {
	Prefixes []string
	Commands map[string]*Command
}

// Command represents a command
type Command struct {
	SubCommands map[string]*Command
	Handler     func(session *discordgo.Session, event *discordgo.MessageCreate, args []string)
}
