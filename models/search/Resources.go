package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// ResourcesSearch searches for resources
type ResourcesSearch struct {
	*search
}

// Perform gets the results
func (search ResourcesSearch) Perform() {

	collections := search.getNonEmptyFromDB()

	for _, collection := range collections {
		result := &Result{
			Title:       collection.Name,
			URL:         search.getURL(collection),
			Description: collection.Description + " -  " + strconv.Itoa(collection.NumberOfResources) + " resources",
			Icon:        "/static/img/public/repository/resource.png",
			Category:    "Collection",
		}
		search.response.AppendItem(result)
	}
}

// gets the repositoryObjects
func (search ResourcesSearch) getNonEmptyFromDB() []wisply.Collection {
	var list []wisply.Collection
	fieldsList := "`resource`"
	limitClause := search.options.GetLimit()
	whereClause := "WHERE `value` LIKE ? "
	sql := "SELECT DISTINCT " + fieldsList + " FROM `resource_key` " + whereClause + space + limitClause
	fmt.Println(sql)
	rows, err := database.Connection.Query(sql, search.likeQuery(), search.likeQuery())
	if err != nil {
		fmt.Println(sql)
		fmt.Println(err)
	}
	for rows.Next() {
		collection := wisply.Collection{}
		rows.Scan(&collection.ID, &collection.Name, &collection.Description, &collection.Spec, &collection.Path, &collection.NumberOfResources, &collection.Repository)
		list = append(list, collection)
	}
	return list
}

func (search ResourcesSearch) getURL(collection wisply.Collection) string {
	path := "/repository/" + strconv.Itoa(collection.Repository) + "#list|0-15*collection|" + strconv.Itoa(collection.ID)
	return path
}

// NewResourcesSearch creates a new search object for finding repositoryObjects
func NewResourcesSearch(search *search) ResourcesSearch {
	return ResourcesSearch{
		search: search,
	}
}
