package config

import (
	"encoding/json"
	"os"
)

// GetDatabase returns the configuration of database
func GetDatabase() *SQLConfiguration {
	conf := new(SQLConfiguration)
	path := conf.GetPath()
	file, _ := getFile(path)
	decoder := json.NewDecoder(file)
	decoder.Decode(&conf)
	return conf
}

func getFile(pathToFile string) (*os.File, error) {
	var directory string
	directory = "conf/"
	pathToFile = directory + pathToFile
	file, err := os.Open(pathToFile)
	return file, err
}
