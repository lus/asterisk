package embeds

import (
	"fmt"
	"time"

	"github.com/Lukaesebrot/asterisk/nodes/users"
	"github.com/bwmarrin/discordgo"
)

// User generates an user information embed
func User(dcUser *discordgo.User, intUser *users.User) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "User Information",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x0893d8,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "ID",
				Value: fmt.Sprintf("`%s`", dcUser.ID),
			},
			{
				Name:  "Tag",
				Value: fmt.Sprintf("`%s`", dcUser.String()),
			},
			{
				Name:  "Avatar hash",
				Value: fmt.Sprintf("`%s`", dcUser.Avatar),
			},
			{
				Name:  "Avatar URL",
				Value: fmt.Sprintf("%s", dcUser.AvatarURL("512")),
			},
			{
				Name:  "Internal flag integer",
				Value: fmt.Sprintf("`%d`", intUser.Flags),
			},
		},
	}
}
