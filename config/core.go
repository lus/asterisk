package config

import (
	"os"
	"strings"

	"github.com/Lukaesebrot/asterisk/static"
	"github.com/joho/godotenv"
)

// CurrentConfig holds the current bot configuration
var CurrentConfig *Config = new(Config)

// Config represents the bot configuration
type Config struct {
	BotToken              string        `json:"botToken"`
	BotAdmins             []string      `json:"botAdmins"`
	FeatureRequestChannel string        `json:"featureRequestChannel"`
	MongoConnectionString string        `json:"mongoConnectionString"`
	JDoodle               JDoodleConfig `json:"jdoodle"`
}

// JDoodleConfig represents the configuration of a JDoodle subscription
type JDoodleConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
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
		BotAdmins:             strings.Split(os.Getenv("ASTERISK_BOT_ADMINS"), ","),
		FeatureRequestChannel: os.Getenv("ASTERISK_FEATURE_REQUEST_CHANNEL"),
		MongoConnectionString: os.Getenv("ASTERISK_MONGO_CONNECTION_STRING"),
		JDoodle: JDoodleConfig{
			ClientID:     os.Getenv("ASTERISK_JDOODLE_CLIENT_ID"),
			ClientSecret: os.Getenv("ASTERISK_JDOODLE_CLIENT_SECRET"),
		},
	}
}
