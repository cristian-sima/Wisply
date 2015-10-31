package config

import (
	"os"
)

// SQLConfiguration encapsulates the details to connect to the database
type SQLConfiguration struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// GetPath checks if there is a custom configuration for database it loads it
// Otherwise, it loads the defualt configuration
func (config *SQLConfiguration) GetPath() string {
	var path, defaultFile, customFile string
	path = "database/"
	defaultFile = path + "default.json"
	customFile = path + "custom.json"
	if _, err := os.Open("conf/" + customFile); err == nil {
		return customFile
	}
	return defaultFile
}
