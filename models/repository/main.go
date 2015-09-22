// Package repository manages the repositories
package repository

import (
	"errors"

	adapter "github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
)

// Model contains the main operations for repositories
type Model struct {
}

// GetAll returns an array of Repository with all repositories
func (model *Model) GetAll() []Repository {
	var list []Repository
	sql := "SELECT id, name, url, description FROM repository"
	rows, _ := database.Database.Query(sql)
	for rows.Next() {
		repository := Repository{}
		rows.Scan(&repository.ID, &repository.Name, &repository.URL, &repository.Description)
		list = append(list, repository)
	}

	return list
}

// NewRepository creates a new repository using the ID
func (model *Model) NewRepository(ID string) (*Repository, error) {

	repository := new(Repository)
	isValid := isValidID(ID)
	if !isValid.IsValid {
		return repository, errors.New("Validation invalid")
	}
	sql := "SELECT id, name, url, description FROM repository WHERE id = ?"
	query, err := database.Database.Prepare(sql)
	query.QueryRow(ID).Scan(&repository.ID, &repository.Name, &repository.URL, &repository.Description)
	if err != nil {
		return repository, errors.New("No repository like that")
	}
	return repository, nil
}

// InsertNewRepository tries to create a new repository
func (model *Model) InsertNewRepository(repositoryDetails map[string]interface{}) (adapter.WisplyError, error) {

	problem := adapter.WisplyError{}

	result := hasValidDetails(repositoryDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("Error")
	}

	name := repositoryDetails["name"].(string)
	description := repositoryDetails["description"].(string)
	url := repositoryDetails["url"].(string)
	sql := "INSERT INTO `repository` (`name`, `description`, `url`) VALUES (?, ?, ?)"
	query, err := database.Database.Prepare(sql)
	query.Exec(name, description, url)
	if err != nil {
		problem.Message = "No repository like that"
		return problem, errors.New("Error")
	}

	return problem, nil
}

// CountRepositories returns the number of repositories
func CountRepositories() int {
	var number int
	query, _ := database.Database.Prepare("SELECT count(*) FROM repository")
	query.QueryRow().Scan(&number)
	return number
}
