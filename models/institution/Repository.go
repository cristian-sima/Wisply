package institution

import (
	"errors"
	"strconv"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	database "github.com/cristian-sima/Wisply/models/database"
)

// Institution represents a institution for reinstitutions
type Institution struct {
	ID          int
	Name        string
	URL         string
	Description string
}

// Delete removes the institution from database
func (institution *Institution) Delete() error {
	sql := "DELETE from `institution` WHERE id=?"
	query, err := database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(institution.ID))
	return err
}

// Modify changes the details of the institution
func (institution *Institution) Modify(institutionDetails map[string]interface{}) (adapter.WisplyError, error) {
	var problem = adapter.WisplyError{}
	result := hasValidModificationDetails(institutionDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("It does not have valid details")
	}
	err := institution.updateDatabase(institutionDetails)
	return problem, err
}

func (institution *Institution) updateDatabase(institutionDetails map[string]interface{}) error {
	name := institutionDetails["name"].(string)
	description := institutionDetails["description"].(string)
	id := strconv.Itoa(institution.ID)

	sql := "UPDATE `institution` SET name=?, description=? WHERE id=?"
	institution.Name = name
	institution.Description = description
	query, _ := database.Database.Prepare(sql)
	_, err := query.Exec(name, description, id)
	return err
}

// ModifyURL changes the URL
func (institution *Institution) ModifyURL(URL string) error {

	result := isValidURL(URL)
	if !result.IsValid {
		return errors.New("It does not have valid URL")
	}
	id := strconv.Itoa(institution.ID)

	institution.URL = URL

	sql := "UPDATE `institution` SET URL=? WHERE id=?"
	query, _ := database.Database.Prepare(sql)
	_, err := query.Exec(URL, id)
	return err
}
