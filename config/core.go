package config

import (
	"os"

	"github.com/Lukaesebrot/asterisk/static"
	"github.com/joho/godotenv"
)

// CurrentConfig holds the current bot configuration
var CurrentConfig *Config = new(Config)

// Config represents the bot configuration
type Config struct {
	BotToken              string
	InitialAdminID        string
	MongoConnectionString string
	Channels              *SectionChannels
	Redeployment          *SectionRedeployment
}

// SectionChannels represents the 'channels' configuration section
type SectionChannels struct {
	FeatureRequests string
	BugReports      string
}

// SectionRedeployment represents the 'redeployment' configuration section
type SectionRedeployment struct {
	WebhookURL string
	Token      string
}

// Load loads the bot configuration
func Load() {
	// Load the .env file if the bot runs in development mode
	if static.Mode != "prod" {
		godotenv.Load()
	}

	// Set the current configuration
	CurrentConfig = &Config{
		BotToken:              os.Getenv("ASTERISK_BOT_TOKEN"),
		InitialAdminID:        os.Getenv("ASTERISK_INITIAL_ADMIN_ID"),
		MongoConnectionString: os.Getenv("ASTERISK_MONGO_CONNECTION_STRING"),
		Channels: &SectionChannels{
			FeatureRequests: os.Getenv("ASTERISK_FEATURE_REQUEST_CHANNEL"),
			BugReports:      os.Getenv("ASTERISK_BUG_REPORT_CHANNEL"),
		},
		Redeployment: &SectionRedeployment{
			WebhookURL: os.Getenv("ASTERISK_REDEPLOYMENT_WEBHOOK_URL"),
			Token:      os.Getenv("ASTERISK_REDEPLOYMENT_TOKEN"),
		},
	}
}
