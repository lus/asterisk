package static

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	// Self holds the bot user and gets initialized on the bot startup
	Self *discordgo.User

	// StartTime holds the time when the bot started
	StartTime time.Time

	// Blacklist holds all the blacklisted users
	Blacklist = map[string]bool{}
)
