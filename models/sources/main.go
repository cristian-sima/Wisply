package sources

import (
	"errors"
	. "github.com/cristian-sima/Wisply/models/wisply"
)

type Model struct {
}

func (model *Model) GetAll() []Source {
	var list []Source
	Database.Raw("SELECT id, name, url, description FROM source").QueryRows(&list)
	return list
}

func (model *Model) GetSourceById(rawIndex string) (*Source, error) {
	var isValid bool
	source := new(Source)
	isValid = ValidateIndex(rawIndex)
	if !isValid {
		return source, errors.New("Validation invalid")
	}
	error := Database.Raw("SELECT name, url, description FROM source WHERE id = ?", rawIndex).QueryRow(&source)
	return source, error
}

func (model *Model) ValidateSource(rawData map[string]interface{}) (map[string][]string, error) {
	validationResult := ValidateSourceDetails(rawData)
	if !validationResult.IsValid {
		return validationResult.Errors, errors.New("Validation invalid")
	}
	return nil, nil
}

func (model *Model) UpdateSourceById(sourceId string, rawData map[string]interface{}) error {
	stringElements := []string{rawData["name"].(string),
		rawData["description"].(string),
		rawData["url"].(string),
		sourceId}
	_, err := Database.Raw("UPDATE `source` SET name=?, description=?, url=? WHERE id=?", stringElements).Exec()
	return err
}

func (model *Model) DeleteSourceById(id string) error {
	elememts := []string{id}
	_, err := Database.Raw("DELETE from `source` WHERE id=?", elememts).Exec()
	return err
}

func (model *Model) InsertNewSource(rawData map[string]interface{}) error {
	stringElements := []string{rawData["name"].(string),
		rawData["description"].(string),
		rawData["url"].(string)}
	_, err := Database.Raw("INSERT INTO `source` (`name`, `description`, `url`) VALUES (?, ?, ?)", stringElements).Exec()
	return err
}

func Count() int {
	var number int
	Database.Raw("SELECT count(*) FROM source").QueryRow(&number)
	return number
}
