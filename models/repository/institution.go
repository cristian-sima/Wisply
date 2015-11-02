package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
)

// Institution represents a institution for reinstitutions
type Institution struct {
	ID          int
	Name        string
	URL         string
	Description string
	LogoURL     string
	WikiID      string
	WikiURL     string
}

// Delete removes the institution from database
func (institution *Institution) Delete() error {

	// delete institution
	sql := "DELETE FROM `institution` WHERE id=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	fmt.Println(institution)
	query.Exec(institution.ID)

	return err
}

// Modify changes the details of the institution
func (institution *Institution) Modify(institutionDetails map[string]interface{}) (adapter.WisplyError, error) {
	var problem = adapter.WisplyError{}
	result := hasValidInstitutionModifyDetails(institutionDetails)
	if !result.IsValid {
		problem.Data = result
		return problem, errors.New("It does not have valid details")
	}
	err := institution.updateInstitutionInDatabase(institutionDetails)
	return problem, err
}

// GetRepositories returns the list of repositories
func (institution *Institution) GetRepositories() []Repository {
	var list []Repository
	fieldList := "`id`, `name`, `url`, `description`, `status`, `institution`, `category`, `public_url`, `lastProcess`"
	sql := "SELECT " + fieldList + " FROM repository WHERE institution = ?"
	rows, _ := database.Connection.Query(sql, institution.ID)
	for rows.Next() {
		repository := Repository{}
		rows.Scan(&repository.ID, &repository.Name, &repository.URL, &repository.Description, &repository.Status, &repository.Institution, &repository.Category, &repository.PublicURL, &repository.LastProcess)
		list = append(list, repository)
	}
	return list
}

func (institution *Institution) updateInstitutionInDatabase(institutionDetails map[string]interface{}) error {
	name := institutionDetails["name"].(string)
	description := institutionDetails["description"].(string)
	id := strconv.Itoa(institution.ID)
	logoURL := institutionDetails["logoURL"].(string)
	wikiURL := institutionDetails["wikiURL"].(string)
	wikiID := institutionDetails["wikiID"].(string)

	setClause := "SET name=?, description=?, logoURL=?, wikiURL=?, wikiID=?"
	sql := "UPDATE `institution` " + setClause + " WHERE id=?"
	institution.Name = name
	institution.Description = description
	query, _ := database.Connection.Prepare(sql)
	_, err := query.Exec(name, description, logoURL, wikiURL, wikiID, id)
	return err
}
