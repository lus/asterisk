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
	BotToken              string `json:"botToken"`
	FeatureRequestChannel string `json:"featureRequestChannel"`
	BugReportChannel      string `json:"bugReportChannel"`
	MongoConnectionString string `json:"mongoConnectionString"`
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
		FeatureRequestChannel: os.Getenv("ASTERISK_FEATURE_REQUEST_CHANNEL"),
		BugReportChannel:      os.Getenv("ASTERISK_BUG_REPORT_CHANNEL"),
		MongoConnectionString: os.Getenv("ASTERISK_MONGO_CONNECTION_STRING"),
	}
}
