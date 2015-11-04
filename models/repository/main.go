// Package repository contains the objects for encapsulating the information
// about the remote repositories and institutions
package repository

import (
	"errors"

	"github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
)

// ResetAllRepositories sets the default values for the repositories
func ResetAllRepositories() {
	sql := "UPDATE `repository` SET `last_process`=0, `status`='unverified'"
	query, _ := database.Connection.Prepare(sql)
	query.Exec()
}

// GetAllInstitutions returns an array of Institution with all institutions
func GetAllInstitutions() []Institution {
	var list []Institution
	sql := "SELECT `id`, `name`, `url`, `description`, `logoURL`, `wikiURL`, `wikiID` FROM `institution`"
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		institution := Institution{}
		rows.Scan(&institution.ID, &institution.Name, &institution.URL, &institution.Description, &institution.LogoURL, &institution.WikiURL, &institution.WikiID)
		list = append(list, institution)
	}
	return list
}

// InsertNewInstitution tries to create a new institution
func InsertNewInstitution(institutionDetails map[string]interface{}) (adapter.WisplyError, error) {

	problem := adapter.WisplyError{}

	result := hasValidInstitutionInsertDetails(institutionDetails)
	if !result.IsValid {
		problem.Data = result
		return problem, errors.New("Error")
	}

	name := institutionDetails["name"].(string)
	description := institutionDetails["description"].(string)
	url := institutionDetails["url"].(string)
	logoURL := institutionDetails["logoURL"].(string)
	wikiURL := institutionDetails["wikiURL"].(string)
	wikiID := institutionDetails["wikiID"].(string)

	sql := "INSERT INTO `institution` (`name`, `description`, `url`, `logoURL`, `wikiURL`, `wikiID`) VALUES (?, ?, ?, ?, ?, ?)"
	query, err := database.Connection.Prepare(sql)
	query.Exec(name, description, url, logoURL, wikiURL, wikiID)

	if err != nil {
		problem.Message = "No institution like that"
		return problem, errors.New("Error")
	}

	return problem, nil
}

// GetAllRepositories returns an array of Repository with all repositories
func GetAllRepositories() []Repository {
	var list []Repository
	sql := "SELECT `id`, `name`, `url`, `description`, `status`, `institution`, `category`, `public_url`, `last_process` FROM `repository`"
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		repository := Repository{}
		rows.Scan(&repository.ID, &repository.Name, &repository.URL, &repository.Description, &repository.Status, &repository.Institution, &repository.Category, &repository.PublicURL, &repository.LastProcess)
		list = append(list, repository)
	}
	return list
}

// GetAllStatus returns an array of Repository with all repositories
func GetAllStatus() []Repository {
	var list []Repository

	sql := "SELECT id, status FROM repository"
	rows, _ := database.Connection.Query(sql)
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

	fieldsList := "`id`, `name`, `url`, `description`, `logoURL`, `wikiURL`, `wikiID`"
	sql := "SELECT " + fieldsList + " FROM institution WHERE id = ?"
	query, err := database.Connection.Prepare(sql)

	query.QueryRow(ID).Scan(&institution.ID, &institution.Name, &institution.URL, &institution.Description, &institution.LogoURL, &institution.WikiURL, &institution.WikiID)

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
	sql := "SELECT id, `name`, `url`, `description`, `status`, `institution`, `category`, `public_url`, `last_process` FROM repository WHERE id = ?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return repository, errors.New("No repository like that")
	}
	query.QueryRow(ID).Scan(&repository.ID, &repository.Name, &repository.URL, &repository.Description, &repository.Status, &repository.Institution, &repository.Category, &repository.PublicURL, &repository.LastProcess)
	return repository, nil
}

// InsertNewRepository tries to create a new repository
func InsertNewRepository(repositoryDetails map[string]interface{}) (adapter.WisplyError, error) {

	problem := adapter.WisplyError{}

	// check Institution

	_, errInst := NewInstitution(repositoryDetails["institution"].(string))

	if errInst != nil {
		problem.Message = "This institution does not exist"
		return problem, errors.New("Error")
	}

	result := hasValidInsertDetails(repositoryDetails)
	if !result.IsValid {
		problem.Data = result
		return problem, errors.New("Error")
	}

	name := repositoryDetails["name"].(string)
	description := repositoryDetails["description"].(string)
	url := repositoryDetails["url"].(string)
	institutionID := repositoryDetails["institution"].(string)
	category := repositoryDetails["category"].(string)
	publicURL := repositoryDetails["public-url"].(string)
	sql := "INSERT INTO `repository` (`name`, `description`, `url`, `institution`, `category`, `public_url`) VALUES (?, ?, ?, ?, ?, ?)"
	query, err := database.Connection.Prepare(sql)
	query.Exec(name, description, url, institutionID, category, publicURL)
	if err != nil {
		problem.Message = "No repository like that"
		return problem, errors.New("Error")
	}

	return problem, nil
}

// CountRepositories returns the number of repositories
func CountRepositories() int {
	var number int
	query, _ := database.Connection.Prepare("SELECT count(*) FROM repository")
	query.QueryRow().Scan(&number)
	return number
}
