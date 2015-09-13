package sources

import (
	"errors"
	. "github.com/cristian-sima/Wisply/models/adapter"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"strconv"
)

type Source struct {
	Id          int
	Name        string
	Url         string
	Description string
}

func (source *Source) Delete() error {
	elememts := []string{
		strconv.Itoa(source.Id),
	}
	_, err := Database.Raw("DELETE from `source` WHERE id=?", elememts).Exec()
	return err
}

func (source *Source) Modify(sourceDetails map[string]interface{}) (WisplyError, error) {

	var problem = WisplyError{}

	result := HasValidDetails(sourceDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("It does not have valid details")
	}

	err := source.update(sourceDetails)

	return problem, err
}

func (source *Source) update(sourceDetails map[string]interface{}) error {
	stringElements := []string{sourceDetails["name"].(string),
		sourceDetails["description"].(string),
		sourceDetails["url"].(string),
		strconv.Itoa(source.Id),
	}
	_, err := Database.Raw("UPDATE `source` SET name=?, description=?, url=? WHERE id=?", stringElements).Exec()
	return err
}
