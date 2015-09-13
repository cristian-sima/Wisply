package config

import (
	"os"
)

type SQLConfiguration struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (config *SQLConfiguration) GetPath() string {
	var path, defaultFile, customFile string
	path = "database/"
	defaultFile = path + "default.json"
	customFile = path + "custom.json"
	if _, err := os.Open("/conf/" + customFile); err == nil {
		return customFile
	}
	return defaultFile
}
