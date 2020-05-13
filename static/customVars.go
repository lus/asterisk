package static

import (
	"github.com/bwmarrin/discordgo"
)

var (
	// Self holds the bot user and gets initialized on the bot startup
	Self *discordgo.User

	// Blacklist holds all the blacklisted users
	Blacklist = map[string]bool{}
)
