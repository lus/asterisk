package todos

import "go.mongodb.org/mongo-driver/bson/primitive"

// ToDo represents a ToDo object
type ToDo struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  string             `bson:"userID"`
	Created int64              `bson:"created"`
	Content string             `bson:"content"`
}
