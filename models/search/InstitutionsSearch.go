package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InstitutionSearch represents a search object for the institutions
type InstitutionSearch struct {
	*search
}

// Perform gets institutions
func (search *InstitutionSearch) Perform() Result {
	searchQuery := Result{
		Results: []ResultItem{},
	}
	institutions := search.getFromDB()
	for _, institution := range institutions {
		result := ResultItem{
			Title:       institution.Name,
			URL:         search.getURL(institution.ID),
			Description: institution.Description,
			Icon:        institution.LogoURL,
		}
		searchQuery.Results = append(searchQuery.Results, result)
	}
	return searchQuery
}

func (search *InstitutionSearch) getFromDB() []repository.Institution {
	var list []repository.Institution
	fieldsList := "`id`, `name`, `url`, `description`, `logoURL`, `wikiURL`, `wikiID`"
	whereClause := "WHERE `name` LIKE ? OR `url` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM institution " + whereClause
	rows, err := database.Connection.Query(sql, search.likeText(), search.likeText())
	fmt.Println(sql)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		fmt.Println("da")
		institution := repository.Institution{}
		rows.Scan(&institution.ID, &institution.Name, &institution.URL, &institution.Description, &institution.LogoURL, &institution.WikiURL, &institution.WikiID)
		list = append(list, institution)
	}
	return list
}

func (search *InstitutionSearch) getURL(institutionID int) string {
	path := "/institutions/"
	action := path + strconv.Itoa(institutionID)
	return action
}

// NewInstitutionsSearch creates a new search object for finding institutions
func NewInstitutionsSearch(text string) InstitutionSearch {
	return InstitutionSearch{
		search: &search{
			text:     text,
			category: "institutions",
		},
	}
}
