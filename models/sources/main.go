package sources

import (
	"errors"
	. "github.com/cristian-sima/Wisply/models/adapter"
	. "github.com/cristian-sima/Wisply/models/wisply"
)

type Model struct {
}

func (model *Model) GetAll() []Source {
	var list []Source
	Database.Raw("SELECT id, name, url, description FROM source").QueryRows(&list)
	return list
}

func (model *Model) NewSource(rawIndex string) (*Source, error) {

	source := new(Source)
	isValid := IsValidId(rawIndex)
	if !isValid.IsValid {
		return source, errors.New("Validation invalid")
	}
	err := Database.Raw("SELECT id, name, url, description FROM source WHERE id = ?", rawIndex).QueryRow(&source)
	if err != nil {
		return source, errors.New("No source like that")
	}
	return source, nil
}

func (model *Model) InsertNewSource(sourceDetails map[string]interface{}) (WisplyError, error) {

	var problem = WisplyError{}

	result := HasValidDetails(sourceDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("Error")
	}

	stringElements := []string{sourceDetails["name"].(string),
		sourceDetails["description"].(string),
		sourceDetails["url"].(string)}
	_, err := Database.Raw("INSERT INTO `source` (`name`, `description`, `url`) VALUES (?, ?, ?)", stringElements).Exec()

	if err != nil {
		problem.Message = "No source like that"
		return problem, errors.New("Error")
	}

	return problem, nil
}

func CountSources() int {
	var number int
	Database.Raw("SELECT count(*) FROM source").QueryRow(&number)
	return number
}
