package starboard

import (
	"context"
	"time"

	"github.com/Lukaesebrot/asterisk/database"
	"github.com/Lukaesebrot/asterisk/static"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddEntry creates a new starboard entry
func AddEntry(sourceChannel, sourceMessageID, starboardChannel, starboardMessageID string) (*Entry, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("starboardEntries")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define the starboard entry object
	entry := &Entry{
		SourceChannel:      sourceChannel,
		SourceMessageID:    sourceMessageID,
		StarboardChannel:   starboardChannel,
		StarboardMessageID: starboardMessageID,
	}

	// Insert the starboard entry object
	insertResult, err := collection.InsertOne(ctx, entry)
	if err != nil {
		return nil, err
	}

	// Return the starboard entry object
	entry.ID = insertResult.InsertedID.(primitive.ObjectID)
	return entry, nil
}

// GetEntry returns the starboard entry object of the given source channel and message ID
func GetEntry(sourceChannel, sourceMessageID string) (*Entry, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("starboardEntries")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the corresponding starboard entry document
	filter := bson.M{
		"sourceChannel":   sourceChannel,
		"sourceMessageID": sourceMessageID,
	}
	result := collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	// Parse the document into a starboard entry object and return it
	entry := new(Entry)
	err = result.Decode(entry)
	return entry, err
}
