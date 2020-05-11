package commands

import (
	"log"

	"github.com/Lukaesebrot/asterisk/utils"

	"github.com/bwmarrin/discordgo"
)

// Info handles the info command
func Info() func(*discordgo.Session, *discordgo.MessageCreate, []string) {
	return func(session *discordgo.Session, event *discordgo.MessageCreate, args []string) {
		_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, utils.GenerateBotInfoEmbed())
		if err != nil {
			log.Println("[ERR] " + err.Error())
		}
	}
}
