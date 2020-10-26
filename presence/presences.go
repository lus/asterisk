package presence

import "github.com/bwmarrin/discordgo"

// presences contains the presences the bot uses
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
