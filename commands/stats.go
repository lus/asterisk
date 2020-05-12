package commands

import (
	"log"

	"github.com/Lukaesebrot/asterisk/utils"

	"github.com/bwmarrin/discordgo"
)

// Stats handles the stats command
func Stats() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateStatsEmbed(session))
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}
