import (
	"os"
)

// Config represents the application's configuration.
type Config struct {
	Port     string // The port number on which the server will listen.
	Database string // The connection string for the database.
}

// LoadConfig loads and returns a new Config instance with values from environment variables.
func LoadConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Database: getEnv("DATABASE_URL", "sqlite://aicraft.db"),
	}
}

// getEnv retrieves the value of an environment variable or a default value if it's not set.
func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}