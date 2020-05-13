package config

import (
	"encoding/json"
	"io/ioutil"
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
	MongoConnectionString string        `json:"mongoConnectionString"`
	JDoodle               JDoodleConfig `json:"jdoodle"`
}

// JDoodleConfig represents the configuration of a JDoodle subscription
type JDoodleConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// Load loads the bot configuration
func Load() error {
	// Load the configuration using a .env file if the bot runs in development mode
	if static.Mode != "prod" {
		godotenv.Load()
		CurrentConfig = &Config{
			BotToken:              os.Getenv("ASTERISK_BOT_TOKEN"),
			BotAdmins:             strings.Split(os.Getenv("ASTERISK_BOT_ADMINS"), ","),
			MongoConnectionString: os.Getenv("ASTERISK_MONGO_CONNECTION_STRING"),
			JDoodle: JDoodleConfig{
				ClientID:     os.Getenv("ASTERISK_JDOODLE_CLIENT_ID"),
				ClientSecret: os.Getenv("ASTERISK_JDOODLE_CLIENT_SECRET"),
			},
		}
		return nil
	}

	// Load the configuration using a .json file if the bot runs in production mode
	file, err := os.Open("data/config.json")
	defer file.Close()
	if err != nil {
		// Create the file if it does not exist
		if os.IsNotExist(err) {
			// Create the file itself
			file, err := os.Create("data/config.json")
			if err != nil {
				return err
			}

			// Define the JSON content
			json, err := json.MarshalIndent(Config{}, "", "    ")
			if err != nil {
				return err
			}

			// Write the JSON content into the file
			_, err = file.Write(json)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Read all bytes out of the file
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Unmarshal the configuration and return it
	return json.Unmarshal(raw, CurrentConfig)
}
