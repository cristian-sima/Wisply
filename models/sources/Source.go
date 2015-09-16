package sources

import (
	"errors"
	"strconv"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	wisply "github.com/cristian-sima/Wisply/models/wisply"
)

// Source represents a source for resources
type Source struct {
	ID          int
	Name        string
	URL         string
	Description string
}

// Delete removes the source from database
func (source *Source) Delete() error {
	elememts := []string{
		strconv.Itoa(source.ID),
	}
	_, err := wisply.Database.Raw("DELETE from `source` WHERE id=?", elememts).Exec()
	return err
}

// Modify changes the details of the source
func (source *Source) Modify(sourceDetails map[string]interface{}) (adapter.WisplyError, error) {

	var problem = adapter.WisplyError{}

	result := hasValidDetails(sourceDetails)
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
		strconv.Itoa(source.ID),
	}
	_, err := wisply.Database.Raw("UPDATE `source` SET name=?, description=?, url=? WHERE id=?", stringElements).Exec()
	return err
}
