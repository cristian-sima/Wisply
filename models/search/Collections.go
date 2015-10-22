package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// CollectionsSearch searches for collections
type CollectionsSearch struct {
	*search
}

// Perform gets the results
func (search CollectionsSearch) Perform() {

	collections := search.getNonEmptyFromDB()

	for _, collection := range collections {
		result := &Result{
			Title:       collection.Name,
			URL:         search.getURL(collection),
			Description: collection.Description + "<br /> " + strconv.Itoa(collection.NumberOfResources) + " resources",
			Icon:        "/static/img/public/repository/repository.png",
			Category:    "Collection",
		}
		search.response.AppendItem(result)
	}
}

// gets the repositoryObjects
func (search CollectionsSearch) getNonEmptyFromDB() []wisply.Collection {
	var list []wisply.Collection
	fieldsList := "`id`, `name`, `description`, `spec`, `path`, `numberOfRecords`, `repository`"
	limitClause := search.options.GetLimit()
	whereClause := "WHERE `name` LIKE ? OR `description` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM `repository_collection` " + whereClause + space + limitClause
	rows, err := database.Connection.Query(sql, search.likeQuery(), search.likeQuery())
	if err != nil {
		fmt.Println(sql)
		fmt.Println(err)
	}
	for rows.Next() {
		collection := wisply.Collection{}
		rows.Scan(&collection.ID, &collection.Name, &collection.Description, &collection.Spec, &collection.Path, &collection.NumberOfResources, &collection.Repository)
		if collection.NumberOfResources != 0 {
			list = append(list, collection)
		}
	}
	return list
}

func (search CollectionsSearch) getURL(collection wisply.Collection) string {
	path := "/repository/" + strconv.Itoa(collection.Repository) + "#list|0-15*collection|" + strconv.Itoa(collection.ID)
	return path
}

// NewCollectionsSearch creates a new search object for finding repositoryObjects
func NewCollectionsSearch(search *search) CollectionsSearch {
	return CollectionsSearch{
		search: search,
	}
}
