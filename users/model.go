package users

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a bot user
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	DiscordID   string             `bson:"discordID"`
	Permissions int                `bson:"permissions"`
}
