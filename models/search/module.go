package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// ModuleSearch searches for modules
type ModuleSearch struct {
	*search
}

// Perform gets the results
func (search ModuleSearch) Perform() {
	modules := search.getFromDB()
	for _, module := range modules {
		institutionID := module.GetInstitution()
		institution, _ := repository.NewInstitution(strconv.Itoa(institutionID))
		result := &Result{
			Title:       module.GetTitle(),
			URL:         search.getURL(module),
			Description: institution.Name,
			Icon:        institution.LogoURL,
			Category:    "Module",
			IsVisible:   true,
		}
		search.response.AppendItem(result)
	}
}

// gets the modules
func (search ModuleSearch) getFromDB() []repository.Module {
	var list []repository.Module

	limitClause := search.options.GetLimit()
	fieldsList := "`id`"
	whereClause := "WHERE `title` LIKE ? OR `code` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM `institution_module` " + whereClause + space + limitClause

	rows, err := database.Connection.Query(sql, search.likeQuery(), search.likeQuery())

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		moduleID := ""
		rows.Scan(&moduleID)
		item, _ := repository.NewModule(moduleID)
		list = append(list, item)
	}
	return list
}

func (search ModuleSearch) getURL(module repository.Module) string {
	path := "/institutions/" + strconv.Itoa(module.GetInstitution()) + "/module/" + strconv.Itoa(module.GetID())
	return path
}

// NewModulesSearch creates a new search object for finding modules
func NewModulesSearch(search *search) ModuleSearch {
	return ModuleSearch{
		search: search,
	}
}
