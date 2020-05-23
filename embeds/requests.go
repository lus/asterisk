package embeds

import (
	"strings"
	"time"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// FeatureRequest generates a feature request embed
func FeatureRequest(ctx *dgc.Ctx) *discordgo.MessageEmbed {
	author := ctx.Event.Author
	return &discordgo.MessageEmbed{
		Title:     "Feature Request",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Description",
				Value: "```" + strings.ReplaceAll(ctx.Arguments.Raw(), "`", "'") + "```",
			},
			{
				Name:  "Requester",
				Value: "```" + author.String() + " (" + author.ID + ")```",
			},
		},
	}
}

// BugReport generates a bug report embed
func BugReport(ctx *dgc.Ctx) *discordgo.MessageEmbed {
	author := ctx.Event.Author
	return &discordgo.MessageEmbed{
		Title:     "Bug Report",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Description",
				Value: "```" + strings.ReplaceAll(ctx.Arguments.Raw(), "`", "'") + "```",
			},
			{
				Name:  "Reporter",
				Value: "```" + author.String() + " (" + author.ID + ")```",
			},
		},
	}
}
