package sources

import (
	"errors"
	"strconv"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	database "github.com/cristian-sima/Wisply/models/database"
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
	sql := "DELETE from `source` WHERE id=?"
	query, err := database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(source.ID))
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
	err := source.updateDatabase(sourceDetails)
	return problem, err
}

func (source *Source) updateDatabase(sourceDetails map[string]interface{}) error {
	name := sourceDetails["name"].(string)
	description := sourceDetails["description"].(string)
	url := sourceDetails["url"].(string)
	id := strconv.Itoa(source.ID)

	sql := "UPDATE `source` SET name=?, description=?, url=? WHERE id=?"
	query, _ := database.Database.Prepare(sql)
	_, err := query.Exec(name, description, url, id)
	return err
}
