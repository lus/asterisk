package reminders

import (
	"context"
	"strings"
	"time"

	"github.com/Lukaesebrot/asterisk/database"
	"github.com/Lukaesebrot/asterisk/static"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create creates a new reminder
func Create(userID, channelID string, duration time.Duration, message string) (*Reminder, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("reminders")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define the reminder object
	reminder := &Reminder{
		UserID:    userID,
		ChannelID: channelID,
		Exceeds:   int64(time.Now().Unix() + int64(duration.Round(time.Second).Seconds())),
		Message:   strings.ReplaceAll(message, "`", "'"),
	}

	// Insert the reminder object
	insertResult, err := collection.InsertOne(ctx, reminder)
	if err != nil {
		return nil, err
	}

	// Return the reminder object
	reminder.ID = insertResult.InsertedID.(primitive.ObjectID)
	return reminder, nil
}

// Next returns the reminder which exceeds next
func Next() (*Reminder, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("reminders")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the next reminder
	opts := options.FindOne()
	opts.SetSort(bson.M{"exceeds": 1})
	result := collection.FindOne(ctx, bson.M{}, opts)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	// Parse the document into a reminder object and return it
	reminder := new(Reminder)
	err = result.Decode(reminder)
	return reminder, err
}

// Get returns the reminder of the given user which exceeds the n'th
func Get(userID string, n int64) (*Reminder, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("reminders")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the n'th reminder
	opts := options.FindOne()
	opts.SetSort(bson.M{"exceeds": 1})
	opts.SetSkip(n)
	filter := bson.M{"userID": userID}
	result := collection.FindOne(ctx, filter, opts)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	// Parse the document into a reminder object and return it
	reminder := new(Reminder)
	err = result.Decode(reminder)
	return reminder, err
}

// GetAll returns all reminders of the given user
func GetAll(userID string) ([]Reminder, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("reminders")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the reminders
	opts := options.Find()
	opts.SetSort(bson.M{"exceeds": 1})
	filter := bson.M{"userID": userID}
	result, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	// Parse the documenst into reminder objects and return them
	var reminders []Reminder
	err = result.All(ctx, &reminders)
	return reminders, err
}

// Delete deletes the current reminder
func (reminder *Reminder) Delete() error {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("reminders")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete the document
	filter := bson.M{"_id": reminder.ID}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
