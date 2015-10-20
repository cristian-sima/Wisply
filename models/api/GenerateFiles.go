package api

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var pathOfAPITables = "/cache/api/tables/"

// GenerateTableFile creates the sql table
// It is a factory of files
// After it creates a file, it calls its generate method
func GenerateTableFile(tableName, format string) {
	var file writter
	switch format {
	case "csv":
		file = createCSVFile(tableName)
		break
	}
	file.generate()
}

type writter interface {
	generate()
}

type file struct {
	writter
	name   string
	format string
}

func (file *file) getFullPath() string {
	var getWisplyPath = func() string {
		wisplyPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		return wisplyPath
	}
	pathToFolder := pathOfAPITables
	rawPath := getWisplyPath() + pathToFolder + file.name + "." + file.format
	return strings.Replace(rawPath, "\\", "/", -1)
}

func createFile(tableName, format string) *file {
	return &file{
		name:   tableName,
		format: format,
	}
}
