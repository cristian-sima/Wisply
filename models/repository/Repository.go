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
	Institution int
	Category    string
	PublicURL   string
	LastProcess int
}

// HasBeenProcessed checks if the repository has any process
func (repository *Repository) HasBeenProcessed() bool {
	return (repository.LastProcess != 0)
}

// SetLastProcess changes the ID of last harvesting process
func (repository *Repository) SetLastProcess(processID int) error {
	sql := "UPDATE `repository` SET lastProcess=? WHERE id=?"
	query, _ := database.Connection.Prepare(sql)
	_, err := query.Exec(processID, repository.ID)
	return err
}

// Delete removes the repository from database
func (repository *Repository) Delete() error {

	// delete all things

	// delete repository
	sql := "DELETE from `repository` WHERE id=?"
	query, err := database.Connection.Prepare(sql)
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

// GetInstitution returns a reference to the institution which holds the repository
func (repository *Repository) GetInstitution() *Institution {
	institution, _ := NewInstitution(strconv.Itoa(repository.Institution))
	return institution
}

func (repository *Repository) updateDatabase(repositoryDetails map[string]interface{}) error {
	name := repositoryDetails["name"].(string)
	description := repositoryDetails["description"].(string)
	url := repositoryDetails["url"].(string)
	id := strconv.Itoa(repository.ID)

	sql := "UPDATE `repository` SET name=?, description=?, url=? WHERE id=?"
	repository.Name = name
	repository.Description = description
	query, _ := database.Connection.Prepare(sql)
	_, err := query.Exec(name, description, url, id)
	return err
}

// ModifyURL changes the URL
func (repository *Repository) ModifyURL(URL string) error {

	result := isValidURL(URL)
	if !result.IsValid {
		return errors.New("It does not have valid URL")
	}
	id := strconv.Itoa(repository.ID)

	repository.URL = URL

	sql := "UPDATE `repository` SET URL=? WHERE id=?"
	query, _ := database.Connection.Prepare(sql)
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
	query, _ := database.Connection.Prepare(sql)
	_, err := query.Exec(newStatus, id)

	repository.Status = newStatus

	return err
}

// GetIdentification returns the identification
func (repository *Repository) GetIdentification() *Identification {

	identification := &Identification{}

	sql := "SELECT id, repository, protocol_version, earliest_datestamp, delete_policy, granularity FROM repository_identification WHERE repository = ?"
	query, _ := database.Connection.Prepare(sql)
	query.QueryRow(repository.ID).Scan(&identification.ID, &identification.Repository, &identification.Protocol, &identification.EarliestDatestamp, &identification.RecordPolicy, &identification.Granularity)

	emailsSQL := "SELECT `email` FROM `repository_email` WHERE `repository` = ?"
	smt, _ := database.Connection.Prepare(emailsSQL)
	rows, _ := smt.Query(repository.ID)

	for rows.Next() {
		email := ""
		rows.Scan(&email)
		identification.AdminEmails = append(identification.AdminEmails, email)
	}
	return identification
}
