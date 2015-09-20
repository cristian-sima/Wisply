// Package sources manages the sources
package sources

import (
	"errors"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
)

// Model contains the main operations for sources
type Model struct {
}

// GetAll returns an array of Source with all sources
func (model *Model) GetAll() []Source {
	var list []Source
	sql := "SELECT id, name, url, description FROM source"
	rows, _ := database.Database.Query(sql)
	for rows.Next() {
		source := Source{}
		rows.Scan(&source.ID, &source.Name, &source.URL, &source.Description)
		list = append(list, source)
	}

	return list
}

// NewSource creates a new source using the ID
func (model *Model) NewSource(ID string) (*Source, error) {

	source := new(Source)
	isValid := isValidID(ID)
	if !isValid.IsValid {
		return source, errors.New("Validation invalid")
	}
	sql := "SELECT id, name, url, description FROM source WHERE id = ?"
	query, err := database.Database.Prepare(sql)
	query.QueryRow(ID).Scan(&source.ID, &source.Name, &source.URL, &source.Description)
	if err != nil {
		return source, errors.New("No source like that")
	}
	return source, nil
}

// InsertNewSource tries to create a new source
func (model *Model) InsertNewSource(sourceDetails map[string]interface{}) (adapter.WisplyError, error) {

	problem := adapter.WisplyError{}

	result := hasValidDetails(sourceDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("Error")
	}

	name := sourceDetails["name"].(string)
	description := sourceDetails["description"].(string)
	url := sourceDetails["url"].(string)
	sql := "INSERT INTO `source` (`name`, `description`, `url`) VALUES (?, ?, ?)"
	query, err := database.Database.Prepare(sql)
	query.Exec(name, description, url)
	if err != nil {
		problem.Message = "No source like that"
		return problem, errors.New("Error")
	}

	return problem, nil
}

// CountSources returns the number of sources
func CountSources() int {
	var number int
	query, _ := database.Database.Prepare("SELECT count(*) FROM source")
	query.QueryRow().Scan(&number)
	return number
}
