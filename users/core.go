package users

import (
	"context"
	"time"

	"github.com/Lukaesebrot/asterisk/database"
	"github.com/Lukaesebrot/asterisk/static"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// userCache defines the current user cache
var userCache = make(map[string]*User)

// Retrieve retrieves a user from the database
func Retrieve(userID string) (*User, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("users")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try to retrieve the current user document
	filter := bson.M{"discordID": userID}
	result := collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		// Create the user document and return it instantly if it doesn't exist
		if err == mongo.ErrNoDocuments {
			insertResult, err := collection.InsertOne(ctx, User{
				DiscordID: userID,
			})
			if err != nil {
				return nil, err
			}
			return &User{
				ID:        insertResult.InsertedID.(primitive.ObjectID),
				DiscordID: userID,
			}, nil
		}
		return nil, err
	}

	// Return the retrieved user object
	user := new(User)
	err = result.Decode(user)
	return user, err
}

// RetrieveCached retrieves a user but checks the cache first
func RetrieveCached(userID string) (*User, error) {
	// Check if the cache contains the corresponding user object
	if userCache[userID] != nil {
		return userCache[userID], nil
	}

	// Retrieve the user object from the database and cache it
	user, err := Retrieve(userID)
	if err != nil {
		return nil, err
	}
	userCache[userID] = user
	return user, nil
}

// RemoveFromCache removes a user from the cache
func RemoveFromCache(userID string) {
	delete(userCache, userID)
}

// Update overrides the databases variables with the local ones
func (user *User) Update() error {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("users")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the MongoDB document
	filter := bson.M{"_id": user.ID}
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": user})
	return err
}
