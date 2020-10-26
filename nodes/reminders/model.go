package reminders

import "go.mongodb.org/mongo-driver/bson/primitive"

// Reminder represents a reminder
type Reminder struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Exceeds   int64              `bson:"exceeds"`
	UserID    string             `bson:"userID"`
	ChannelID string             `bson:"channelID"`
	Message   string             `bson:"message"`
}
