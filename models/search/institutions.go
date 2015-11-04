package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InstitutionSearch searches for institutions
type InstitutionSearch struct {
	*search
}

// Perform gets the results
func (search InstitutionSearch) Perform() {
	institutions := search.getFromDB()
	for _, institution := range institutions {
		result := &Result{
			Title:       institution.Name,
			URL:         search.getURL(institution.ID),
			Description: institution.Description,
			Icon:        institution.LogoURL,
			Category:    "Institution",
			IsVisible:   true,
		}
		search.response.AppendItem(result)
	}
}

// gets the institutions
func (search InstitutionSearch) getFromDB() []repository.Institution {
	var list []repository.Institution

	limitClause := search.options.GetLimit()
	fieldsList := "`id`, `name`, `url`, `description`, `logo_URL`, `wikiURL`, `wikiID`"
	whereClause := "WHERE `name` LIKE ? OR `url` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM `institution` " + whereClause + space + limitClause

	rows, err := database.Connection.Query(sql, search.likeQuery(), search.likeQuery())

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		institution := repository.Institution{}
		rows.Scan(&institution.ID, &institution.Name, &institution.URL, &institution.Description, &institution.LogoURL, &institution.WikiURL, &institution.WikiID)
		list = append(list, institution)
	}
	return list
}

func (search InstitutionSearch) getURL(institutionID int) string {
	path := "/institutions/"
	action := path + strconv.Itoa(institutionID)
	return action
}

// NewInstitutionsSearch creates a new search object for finding institutions
func NewInstitutionsSearch(search *search) InstitutionSearch {
	return InstitutionSearch{
		search: search,
	}
}
