package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/education"
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
	_, err = query.Exec(institution.ID)

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
	fieldList := "`id`, `name`, `url`, `description`, `status`, `institution`, `category`, `public_url`, `last_process`"
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

	setClause := "SET name=?, description=?, logo_URL=?, wiki_URL=?, wiki_ID=?"
	sql := "UPDATE `institution` " + setClause + " WHERE id=?"
	institution.Name = name
	institution.Description = description
	query, _ := database.Connection.Prepare(sql)
	_, err := query.Exec(name, description, logoURL, wikiURL, wikiID, id)
	return err
}

// GetEducationSubjects returns the distinct Subjects
func (institution Institution) GetEducationSubjects() []*education.Subject {
	list := []*education.Subject{}
	fieldList := "`subject`"
	whereClause := "WHERE `institution` = ?"
	sql := "SELECT DISTINCT " + fieldList + " FROM `institution_program` " + whereClause
	rows, _ := database.Connection.Query(sql, strconv.Itoa(institution.ID))
	for rows.Next() {
		ID := ""
		rows.Scan(&ID)
		item, _ := education.NewSubject(ID)
		list = append(list, item)
	}
	return list
}

// GetProgramsBySubjectID returns the programs of study for a particular subject
func (institution Institution) GetProgramsBySubjectID(subjectID int) []Program {
	list := []Program{}
	allPrograms := institution.GetPrograms()
	for _, program := range allPrograms {
		if program.GetSubject().GetID() == subjectID {
			list = append(list, program)
		}
	}
	return list
}

// GetPrograms returns programs of study for the institution
func (institution Institution) GetPrograms() []Program {
	var list []Program
	fieldList := "`id`, `institution`, `title`, `content`, `code`, `year`, `ucas_code`, `level`, `subject`"
	orderClause := "ORDER BY `year` DESC"
	whereClause := "WHERE `institution` = ?"
	sql := "SELECT " + fieldList + " FROM `institution_program` " + whereClause + " " + orderClause
	rows, err := database.Connection.Query(sql, strconv.Itoa(institution.ID))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		item := Program{}
		rows.Scan(&item.id, &item.institution, &item.title, &item.content, &item.code, &item.year, &item.ucasCode, &item.level, &item.subject)
		list = append(list, item)
	}
	return list
}

// GetModules returns the list of the modules for the institution
func (institution Institution) GetModules() []Module {
	list := []Module{}

	fieldList := "`id`, `title`, `content`, `code`, `credits`, `year`"
	sql := "SELECT " + fieldList + " FROM `institution_module` WHERE institution=? "
	rows, err := database.Connection.Query(sql, strconv.Itoa(institution.ID))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		module := Module{}
		rows.Scan(&module.id, &module.title, &module.content, &module.code, &module.credits, &module.year)
		list = append(list, module)
	}
	return list
}
