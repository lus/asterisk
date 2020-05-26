package guilds

import "go.mongodb.org/mongo-driver/bson/primitive"

// Guild represents a bot guild
type Guild struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	DiscordID string             `bson:"discordID"`
	Settings  GuildSettings      `bson:"settings"`
}

// GuildSettings represents the settings for a bot guild
type GuildSettings struct {
	CommandChannels  []string `bson:"commandChannels"`
	StarboardChannel string   `bson:"starboardChannel"`
}
