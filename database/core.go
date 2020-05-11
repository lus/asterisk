package database

import (
	"context"
	"time"

	"github.com/Lukaesebrot/asterisk/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// CurrentClient holds the current MongoDB client
var CurrentClient *mongo.Client

// Connect conntects to the configured MongoDB host
func Connect() error {
	// Define the context for the following database operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB host
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.CurrentConfig.MongoConnectionString))
	if err != nil {
		return err
	}

	// Ping the MongoDB host
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	CurrentClient = client
	return nil
}
