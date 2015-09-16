package sources

import (
	"errors"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	wisply "github.com/cristian-sima/Wisply/models/wisply"
)

// Model contains the main operations for sources
type Model struct {
}

// GetAll returns an array of Source with all sources
func (model *Model) GetAll() []Source {
	var list []Source
	wisply.Database.Raw("SELECT id, name, url, description FROM source").QueryRows(&list)
	return list
}

// NewSource creates a new source using the ID
func (model *Model) NewSource(ID string) (*Source, error) {

	source := new(Source)
	isValid := isValidID(ID)
	if !isValid.IsValid {
		return source, errors.New("Validation invalid")
	}
	err := wisply.Database.Raw("SELECT id, name, url, description FROM source WHERE id = ?", ID).QueryRow(&source)
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

	stringElements := []string{sourceDetails["name"].(string),
		sourceDetails["description"].(string),
		sourceDetails["url"].(string)}
	_, err := wisply.Database.Raw("INSERT INTO `source` (`name`, `description`, `url`) VALUES (?, ?, ?)", stringElements).Exec()

	if err != nil {
		problem.Message = "No source like that"
		return problem, errors.New("Error")
	}

	return problem, nil
}

// CountSources returns the number of sources
func CountSources() int {
	var number int
	wisply.Database.Raw("SELECT count(*) FROM source").QueryRow(&number)
	return number
}
