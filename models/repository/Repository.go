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
	result := hasValidModificationDetails(repositoryDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("It does not have valid details")
	}
	err := repository.updateDatabase(repositoryDetails)
	return problem, err
}

// ModifyURL changes the URL
func (repository *Repository) ModifyURL(URL string) error {

	result := isValidURL(URL)
	if !result.IsValid {
		return errors.New("It does not have valid URL")
	}
	id := strconv.Itoa(repository.ID)

	sql := "UPDATE `repository` SET URL=? WHERE id=?"
	query, _ := database.Database.Prepare(sql)
	_, err := query.Exec(URL, id)
	return err
}

// ModifyStatus changes the status
func (repository *Repository) ModifyStatus(newStatus string) error {

	result := isValidStatus(newStatus)
	if !result {
		return errors.New("It does not have valid status")
	}
	id := strconv.Itoa(repository.ID)

	sql := "UPDATE `repository` SET status=? WHERE id=?"
	query, _ := database.Database.Prepare(sql)
	_, err := query.Exec(newStatus, id)
	return err
}

func (repository *Repository) updateDatabase(repositoryDetails map[string]interface{}) error {
	name := repositoryDetails["name"].(string)
	description := repositoryDetails["description"].(string)
	id := strconv.Itoa(repository.ID)

	sql := "UPDATE `repository` SET name=?, description=? WHERE id=?"
	query, _ := database.Database.Prepare(sql)
	_, err := query.Exec(name, description, id)
	return err
}
