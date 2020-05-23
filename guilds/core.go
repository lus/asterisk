package guilds

import (
	"context"
	"time"

	"github.com/Lukaesebrot/asterisk/database"
	"github.com/Lukaesebrot/asterisk/static"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// guildCache defines the current guild cache
var guildCache = make(map[string]*Guild)

// Retrieve retrieves a guild from the database
func Retrieve(guildID string) (*Guild, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("guilds")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try to retrieve the corresponding guild document
	filter := bson.M{"discordID": guildID}
	result := collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		// Create the guild document and return it instantly if it doesn't exist
		if err == mongo.ErrNoDocuments {
			guild := &Guild{
				DiscordID: guildID,
				Settings: GuildSettings{
					CommandChannels: []string{},
				},
			}
			insertResult, err := collection.InsertOne(ctx, guild)
			if err != nil {
				return nil, err
			}
			guild.ID = insertResult.InsertedID.(primitive.ObjectID)
			return guild, nil
		}
		return nil, err
	}

	// Return the retrieved guild object
	guild := new(Guild)
	err = result.Decode(guild)
	return guild, err
}

// RetrieveCached retrieves a guild but checks the cache first
func RetrieveCached(guildID string) (*Guild, error) {
	// Check if the cache contains the corresponding guild object
	if guildCache[guildID] != nil {
		return guildCache[guildID], nil
	}

	// Retrieve the guild object from the database and cache it
	guild, err := Retrieve(guildID)
	if err != nil {
		return nil, err
	}
	guildCache[guildID] = guild
	return guild, nil
}

// RemoveFromCache removes a guild from the cache
func RemoveFromCache(guildID string) {
	delete(guildCache, guildID)
}

// Update overrides the databases variables with the local ones
func (guild *Guild) Update() error {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("guilds")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the MongoDB document
	filter := bson.M{"_id": guild.ID}
	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": guild})
	return err
}
