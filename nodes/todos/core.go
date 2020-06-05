package todos

import (
	"context"
	"time"

	"github.com/Lukaesebrot/asterisk/database"
	"github.com/Lukaesebrot/asterisk/static"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create creates a new ToDo object
func Create(userID, content string) (*ToDo, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("toDos")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define the ToDo object
	toDo := &ToDo{
		UserID:  userID,
		Created: time.Now().Unix(),
		Content: content,
	}

	// Insert the ToDo object
	insertResult, err := collection.InsertOne(ctx, toDo)
	if err != nil {
		return nil, err
	}

	// Return the ToDo object
	toDo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return toDo, nil
}

// Get returns the ToDo object of the given user which was created the n'th
func Get(userID string, n int64) (*ToDo, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("toDos")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the n'th ToDo object
	opts := options.FindOne()
	opts.SetSort(bson.M{"created": 1})
	opts.SetSkip(n)
	filter := bson.M{"userID": userID}
	result := collection.FindOne(ctx, filter, opts)
	err := result.Err()
	if err != nil {
		return nil, err
	}

	// Parse the document into a ToDo object and return it
	toDo := new(ToDo)
	err = result.Decode(toDo)
	return toDo, err
}

// GetAll returns all ToDo objects of the given user
func GetAll(userID string) ([]ToDo, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("toDos")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Retrieve the ToDos
	opts := options.Find()
	opts.SetSort(bson.M{"created": 1})
	filter := bson.M{"userID": userID}
	result, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	// Parse the documenst into ToDo objects and return them
	var toDos []ToDo
	err = result.All(ctx, &toDos)
	return toDos, err
}

// Delete deletes the current ToDo object
func (toDo *ToDo) Delete() error {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("toDos")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete the document
	filter := bson.M{"_id": toDo.ID}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
