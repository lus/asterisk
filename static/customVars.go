package static

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	// Self holds the bot user and gets initialized on the bot startup
	Self *discordgo.User

	// StartupTime holds the time when the bot started
	StartupTime time.Time
)
