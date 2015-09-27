// Package institution manages the institutions
package institution

import (
	"errors"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
)

// Model contains the main operations for institutions
type Model struct {
}

// GetAll returns an array of Institution with all institutions
func (model *Model) GetAll() []Institution {
	var list []Institution
	sql := "SELECT id, name, url, description FROM institution"
	rows, _ := database.Database.Query(sql)
	for rows.Next() {
		institution := Institution{}
		rows.Scan(&institution.ID, &institution.Name, &institution.URL, &institution.Description)
		list = append(list, institution)
	}
	return list
}

// NewInstitution creates a new institution using the ID
func (model *Model) NewInstitution(ID string) (*Institution, error) {

	institution := new(Institution)
	isValid := isValidID(ID)
	if !isValid.IsValid {
		return institution, errors.New("Validation invalid")
	}
	sql := "SELECT id, name, url, description FROM institution WHERE id = ?"
	query, err := database.Database.Prepare(sql)
	query.QueryRow(ID).Scan(&institution.ID, &institution.Name, &institution.URL, &institution.Description)
	if err != nil {
		return institution, errors.New("No institution like that")
	}
	return institution, nil
}

// InsertNewInstitution tries to create a new institution
func (model *Model) InsertNewInstitution(institutionDetails map[string]interface{}) (adapter.WisplyError, error) {

	problem := adapter.WisplyError{}

	result := hasValidInsertDetails(institutionDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("Error")
	}

	name := institutionDetails["name"].(string)
	description := institutionDetails["description"].(string)
	url := institutionDetails["url"].(string)
	sql := "INSERT INTO `institution` (`name`, `description`, `url`) VALUES (?, ?, ?)"
	query, err := database.Database.Prepare(sql)
	query.Exec(name, description, url)
	if err != nil {
		problem.Message = "No institution like that"
		return problem, errors.New("Error")
	}

	return problem, nil
}

// CountInstitutions returns the number of institutions
func CountInstitutions() int {
	var number int
	query, _ := database.Database.Prepare("SELECT count(*) FROM institution")
	query.QueryRow().Scan(&number)
	return number
}
