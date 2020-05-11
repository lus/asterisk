package guildconfig

import (
	"context"
	"time"

	"github.com/Lukaesebrot/asterisk/utils"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Lukaesebrot/asterisk/database"
	"github.com/Lukaesebrot/asterisk/static"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GuildConfig represents the configuration of a guild
type GuildConfig struct {
	ID                 primitive.ObjectID `bson:"_id"`
	GuildID            string             `bson:"guildID"`
	ChannelRestriction bool               `bson:"channelRestriction"`
	CommandChannels    []string           `bson:"commandChannels"`
}

// Retrieve fetches the current configuration of the given guild or creates it if it doesn't exist
func Retrieve(guildID string) (*GuildConfig, error) {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("guilds")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try to retrieve the current guild configuration
	filter := bson.M{"guildID": guildID}
	result := collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		// Create the guild configuration and return it instantly if it doesn't exist
		if err == mongo.ErrNoDocuments {
			insertResult, err := collection.InsertOne(ctx, GuildConfig{
				GuildID: guildID,
			})
			if err != nil {
				return nil, err
			}
			return &GuildConfig{
				ID:      insertResult.InsertedID.(primitive.ObjectID),
				GuildID: guildID,
			}, nil
		}
		return nil, err
	}

	// Return the retrieved guild configuration
	guildConfig := new(GuildConfig)
	err = result.Decode(guildConfig)
	return guildConfig, err
}

// Update overrides the databases variables with the local ones
func (guildConfig *GuildConfig) Update() error {
	// Define the collection to use for this database operation
	collection := database.CurrentClient.Database(static.MongoDatabase).Collection("guilds")

	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the MongoDB document
	filter := bson.M{"_id": guildConfig.ID}
	_, err := collection.UpdateOne(ctx, filter, guildConfig)
	return err
}

// ValidateChannel checks whether or not commands may get executed in the given channel
func (guildConfig *GuildConfig) ValidateChannel(channelID string) bool {
	// Return true if the guild did not enable the command restriction
	if !guildConfig.ChannelRestriction {
		return true
	}

	// Return true if the channel is in the list of command channels
	return utils.StringArrayContains(guildConfig.CommandChannels, channelID)
}
