package presence

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Animate animates the presence using the given session
func Animate(session *discordgo.Session) {
	session.UpdateStatusComplex(presences[rand.Intn(len(presences))])
	time.Sleep(10 * time.Second)
	Animate(session)
}
