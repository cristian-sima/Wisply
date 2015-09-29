package repository

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

// GetRepositories returns the list of repositories
func (institution *Institution) GetRepositories() []Repository {
	var list []Repository
	sql := "SELECT id, name, url, description, status, institution FROM repository WHERE institution = ?"
	rows, _ := database.Database.Query(sql, institution.ID)
	for rows.Next() {
		repository := Repository{}
		rows.Scan(&repository.ID, &repository.Name, &repository.URL, &repository.Description, &repository.Status, &repository.Institution)
		list = append(list, repository)
	}
	return list
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
