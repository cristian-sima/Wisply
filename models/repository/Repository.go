package repository

import (
	"errors"
	"strconv"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	database "github.com/cristian-sima/Wisply/models/database"
)

// Repository represents a repository for rerepositorys
type Repository struct {
	ID          int
	Name        string
	URL         string
	Description string
	Status      string
}

// Delete removes the repository from database
func (repository *Repository) Delete() error {
	sql := "DELETE from `repository` WHERE id=?"
	query, err := database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(repository.ID))
	return err
}

// Modify changes the details of the repository
func (repository *Repository) Modify(repositoryDetails map[string]interface{}) (adapter.WisplyError, error) {
	var problem = adapter.WisplyError{}
	result := hasValidDetails(repositoryDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("It does not have valid details")
	}
	err := repository.updateDatabase(repositoryDetails)
	return problem, err
}

func (repository *Repository) updateDatabase(repositoryDetails map[string]interface{}) error {
	name := repositoryDetails["name"].(string)
	description := repositoryDetails["description"].(string)
	url := repositoryDetails["url"].(string)
	id := strconv.Itoa(repository.ID)

	sql := "UPDATE `repository` SET name=?, description=?, url=? WHERE id=?"
	query, _ := database.Database.Prepare(sql)
	_, err := query.Exec(name, description, url, id)
	return err
}
