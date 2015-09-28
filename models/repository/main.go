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

// GetAllInstitutions returns an array of Institution with all institutions
func (model *Model) GetAllInstitutions() []Institution {
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

// GetAllRepositories returns an array of Repository with all repositories
func (model *Model) GetAllRepositories() []Repository {
	var list []Repository
	sql := "SELECT id, name, url, description, status, institution FROM repository"
	rows, _ := database.Database.Query(sql)
	for rows.Next() {
		repository := Repository{}
		rows.Scan(&repository.ID, &repository.Name, &repository.URL, &repository.Description, &repository.Status, &repository.Institution)
		list = append(list, repository)
	}
	return list
}

// GetAllStatus returns an array of Repository with all repositories
func (model *Model) GetAllStatus() []Repository {
	var list []Repository

	sql := "SELECT id, status FROM repository"
	rows, _ := database.Database.Query(sql)
	for rows.Next() {
		repository := Repository{}
		rows.Scan(&repository.ID, &repository.Status)
		list = append(list, repository)
	}
	return list
}

// NewInstitution creates a new institution using the ID
func NewInstitution(ID string) (*Institution, error) {

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

// NewRepository creates a new repository using the ID
func NewRepository(ID string) (*Repository, error) {

	repository := new(Repository)
	isValid := isValidID(ID)
	if !isValid.IsValid {
		return repository, errors.New("Validation invalid")
	}
	sql := "SELECT id, name, url, description, status, institution FROM repository WHERE id = ?"
	query, err := database.Database.Prepare(sql)
	query.QueryRow(ID).Scan(&repository.ID, &repository.Name, &repository.URL, &repository.Description, &repository.Status, &repository.Institution)
	if err != nil {
		return repository, errors.New("No repository like that")
	}
	return repository, nil
}

// InsertNewRepository tries to create a new repository
func (model *Model) InsertNewRepository(repositoryDetails map[string]interface{}) (adapter.WisplyError, error) {

	problem := adapter.WisplyError{}

	// check Institution

	_, errInst := NewInstitution(repositoryDetails["institution"].(string))

	if errInst != nil {
		problem.Message = "This institution does not exist"
		return problem, errors.New("Error")
	}

	result := hasValidInsertDetails(repositoryDetails)
	if !result.IsValid {
		problem.Data = result.Errors
		return problem, errors.New("Error")
	}

	name := repositoryDetails["name"].(string)
	description := repositoryDetails["description"].(string)
	url := repositoryDetails["url"].(string)
	institutionID := repositoryDetails["institution"].(string)
	sql := "INSERT INTO `repository` (`name`, `description`, `url`, `institution`) VALUES (?, ?, ?, ?)"
	query, err := database.Database.Prepare(sql)
	query.Exec(name, description, url, institutionID)
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
