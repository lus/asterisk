package users

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a bot user
type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	DiscordID   string             `bson:"discordID"`
	Permissions int                `bson:"permissions"`
}
