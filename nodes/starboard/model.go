package starboard

import "go.mongodb.org/mongo-driver/bson/primitive"

// Entry represents a starboard entry
type Entry struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	SourceChannel      string             `bson:"sourceChannel"`
	SourceMessageID    string             `bson:"sourceMessageID"`
	StarboardChannel   string             `bson:"starboardChannel"`
	StarboardMessageID string             `bson:"starboardMessageID"`
}
