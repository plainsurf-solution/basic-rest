// config.go

package common

import (
	"os"

	"github.com/joho/godotenv"
)

// DatabaseConfig represents the configuration for the database
type DatabaseConfig struct {
	Type           string
	Connection     string
	DbName         string
	collectionName string
}

// LoadConfig loads the configuration from the environment file
func LoadConfig() (*DatabaseConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	dbConfig := &DatabaseConfig{
		Type:       os.Getenv("DB_TYPE"),
		Connection: os.Getenv("DB_CONNECTION_STRING"),
	}

	return dbConfig, nil
}
