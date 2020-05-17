package config

import "os"

// Config godoc
type Config struct {
	DbHost               string
	DbName               string
	DbUser               string
	DbPassword           string
	FacebookClientID     string
	FacebookClientSecret string
	FacebookState        string
	FacebookRedirectURL  string
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// New godoc
func New() *Config {
	return &Config{
		DbHost:               getEnv("DB_HOST", ""),
		DbName:               getEnv("DB_NAME", ""),
		DbUser:               getEnv("DB_USER", ""),
		DbPassword:           getEnv("DB_PASSWORD", ""),
		FacebookClientID:     getEnv("FACEBOOK_CLIENT_ID", ""),
		FacebookClientSecret: getEnv("FACEBOOK_CLIENT_SECRET", ""),
		FacebookRedirectURL:  getEnv("FACEBOOK_REDIRECT_URL", ""),
	}
}
