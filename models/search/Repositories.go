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
func (search RepositoriesSearch) Perform() {

	repositoryObjects := search.getFromDB()
	for _, repositoryObject := range repositoryObjects {
		result := &Result{
			Title:       repositoryObject.Name,
			URL:         search.getURL(repositoryObject.ID),
			Description: repositoryObject.Description,
			Icon:        "/static/img/public/repository/repository.png",
			Category:    "Repository",
		}
		search.response.AppendItem(result)
	}
}

// gets the repositoryObjects
func (search RepositoriesSearch) getFromDB() []repository.Repository {
	var list []repository.Repository

	fieldsList := "`id`, `name`, `description`"
	limitClause := search.options.GetLimit()
	whereClause := "WHERE `name` LIKE ? OR `description` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM `repository` " + whereClause + space + limitClause
	rows, err := database.Connection.Query(sql, search.likeQuery(), search.likeQuery())

	if err != nil {
		fmt.Println("Erorr search repositories: ")
		fmt.Println(err)
	}
	for rows.Next() {
		repositoryObject := repository.Repository{}
		rows.Scan(&repositoryObject.ID, &repositoryObject.Name, &repositoryObject.Description)
		list = append(list, repositoryObject)
	}
	return list
}

func (search RepositoriesSearch) getURL(repositoryObjectID int) string {
	path := "/repository/"
	action := path + strconv.Itoa(repositoryObjectID)
	return action
}

// NewRepositoriesSearch creates a new search object for finding repositoryObjects
func NewRepositoriesSearch(search *search) RepositoriesSearch {
	return RepositoriesSearch{
		search: search,
	}
}
