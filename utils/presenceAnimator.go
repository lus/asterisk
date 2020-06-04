package utils

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

// presences defines all available presences
var presences = []discordgo.UpdateStatusData{
	{
		Game: &discordgo.Game{
			Name: "$help",
			Type: discordgo.GameTypeListening,
		},
	},
	{
		Game: &discordgo.Game{
			Name: "you",
			Type: discordgo.GameTypeWatching,
		},
	},
}

// AnimatePresence animates the bots presence
func AnimatePresence(session *discordgo.Session) {
	session.UpdateStatusComplex(presences[rand.Intn(len(presences))])
	time.Sleep(10 * time.Second)
	AnimatePresence(session)
}
