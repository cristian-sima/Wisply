package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// RepositoriesSearch searches for repositories
type RepositoriesSearch struct {
	*search
}

// Perform gets the results
func (search *RepositoriesSearch) Perform() Result {
	searchQuery := Result{
		Results: []ResultItem{},
	}
	repositoryObjects := search.getFromDB()
	for _, repositoryObject := range repositoryObjects {
		result := ResultItem{
			Title:       repositoryObject.Name,
			URL:         search.getURL(repositoryObject.ID),
			Description: repositoryObject.Description,
			Icon:        "/static/img/public/repository/repository.png",
			Category:    "Repository",
		}
		searchQuery.Results = append(searchQuery.Results, result)
	}
	return searchQuery
}

// gets the repositoryObjects
func (search *RepositoriesSearch) getFromDB() []repository.Repository {
	var list []repository.Repository
	fieldsList := "`id`, `name`, `description`"
	whereClause := "WHERE `name` LIKE ? OR `description` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM `repository` " + whereClause
	rows, err := database.Connection.Query(sql, search.likeText(), search.likeText())
	fmt.Println(sql)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		fmt.Println("rep")
		repositoryObject := repository.Repository{}
		rows.Scan(&repositoryObject.ID, &repositoryObject.Name, &repositoryObject.Description)
		list = append(list, repositoryObject)
	}
	return list
}

func (search *RepositoriesSearch) getURL(repositoryObjectID int) string {
	path := "/repository/"
	action := path + strconv.Itoa(repositoryObjectID)
	return action
}

// NewRepositoriesSearch creates a new search object for finding repositoryObjects
func NewRepositoriesSearch(text string) RepositoriesSearch {
	return RepositoriesSearch{
		search: &search{
			text:     text,
			category: "Repository",
		},
	}
}
