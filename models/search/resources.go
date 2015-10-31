package search

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// ResourcesSearch searches for resources
type ResourcesSearch struct {
	*search
}

// Perform gets the results
func (search ResourcesSearch) Perform() {

	resources := search.getNonEmptyFromDB()

	for _, resource := range resources {
		result := &Result{
			Title:       resource.Keys.GetTitle(),
			URL:         resource.GetWisplyURL(),
			Description: search.getDescription(resource),
			Icon:        "/static/img/public/repository/resource.png",
			Category:    search.getCategory(resource),
			IsVisible:   search.getVisibility(resource),
		}
		search.response.AppendItem(result)
	}
}

// gets the repositoryObjects
func (search ResourcesSearch) getNonEmptyFromDB() []wisply.Record {
	var list []wisply.Record

	fieldsList := "`resource`"
	limitClause := search.options.GetLimit()
	whereClause := "WHERE `value` LIKE ? "
	sql := "SELECT DISTINCT " + fieldsList + " FROM `resource_key` " + whereClause + space + limitClause
	fmt.Println(sql)

	rows, err := database.Connection.Query(sql, search.likeQuery())
	if err != nil {
		fmt.Println(sql)
		fmt.Println(err)
	}

	for rows.Next() {
		var identifier string
		rows.Scan(&identifier)
		record := wisply.GetRecordByIdentifier(identifier)
		list = append(list, record)
	}

	return list
}

func (search ResourcesSearch) getName(resource wisply.Record) string {
	return "Name"
}

func (search ResourcesSearch) getCategory(resource wisply.Record) string {
	return "Resource"
}

func (search ResourcesSearch) getVisibility(resource wisply.Record) bool {
	return resource.IsVisible
}

func (search ResourcesSearch) getDescription(resource wisply.Record) string {
	return ""
}

// NewResourcesSearch creates a new search object for finding repositoryObjects
func NewResourcesSearch(search *search) ResourcesSearch {
	return ResourcesSearch{
		search: search,
	}
}
