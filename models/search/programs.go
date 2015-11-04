package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// ProgramSearch searches for programs
type ProgramSearch struct {
	*search
}

// Perform gets the results
func (search ProgramSearch) Perform() {
	programs := search.getFromDB()
	for _, program := range programs {
		institutionID := program.GetInstitution()
		institution, _ := repository.NewInstitution(strconv.Itoa(institutionID))
		result := &Result{
			Title:       program.GetTitle(),
			URL:         search.getURL(program),
			Description: institution.Name,
			Icon:        institution.LogoURL,
			Category:    "Program of study",
			IsVisible:   true,
		}
		search.response.AppendItem(result)
	}
}

// gets the programs
func (search ProgramSearch) getFromDB() []repository.Program {
	var list []repository.Program

	limitClause := search.options.GetLimit()
	fieldsList := "`id`"
	whereClause := "WHERE `title` LIKE ? OR `code` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM `institution_program` " + whereClause + space + limitClause

	rows, err := database.Connection.Query(sql, search.likeQuery(), search.likeQuery())

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		programID := ""
		rows.Scan(&programID)
		item, _ := repository.NewProgram(programID)
		list = append(list, item)
	}
	return list
}

func (search ProgramSearch) getURL(program repository.Program) string {
	path := "/institutions/" + strconv.Itoa(program.GetInstitution()) + "/program/" + strconv.Itoa(program.GetID())
	return path
}

// NewProgramsSearch creates a new search object for finding programs
func NewProgramsSearch(search *search) ProgramSearch {
	return ProgramSearch{
		search: search,
	}
}
