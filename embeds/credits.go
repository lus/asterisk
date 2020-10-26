package embeds

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Credits generates a credits embed
func Credits() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:     "Credits",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x0893d8,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "3rd-party libraries",
				Value: "[github.com/bwmarrin/discordgo](https://github.com/bwmarrin/discordgo)" +
					"\n[github.com/Lukaesebrot/dgc](https://github.com/Lukaesebrot/dgc)" +
					"\n[github.com/containous/yaegi](https://github.com/containous/yaegi)" +
					"\n[github.com/joho/godotenv](https://github.com/joho/godotenv)" +
					"\n[github.com/valyala/fasthttp](https://github.com/valyala/fasthttp)" +
					"\n[github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver)",
			},
			{
				Name: "Cool people",
				Value: "`das_#9677` for `testing`" +
					"\n`Cerus#5149` for `testing`",
			},
		},
	}
}
