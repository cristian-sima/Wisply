package developer

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var pathOfDownloadTables = "/cache/developer/tables/"

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

// an example of a fullpath is
// W:\go-workspace\src\github.com\cristian-sima\Wisply\cache\tables\account.csv
func (file *file) getFullPath() string {
	var getWisplyPath = func() string {
		wisplyPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		return wisplyPath
	}
	pathToFolder := pathOfDownloadTables
	rawPath := getWisplyPath() + pathToFolder + file.getFileName()
	return strings.Replace(rawPath, "\\", "/", -1)
}

func (file *file) getFileName() string {
	return file.name + "." + file.format
}

func createFile(tableName, format string) *file {
	return &file{
		name:   tableName,
		format: format,
	}
}

func copyFile(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

func deleteFile(path string) {
	if err := os.Remove(path); err != nil {
		fmt.Println("No file there")
	}
}
