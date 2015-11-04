package wisply

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
)

type collections struct {
}

// GetCollections returns the collections
func (collections *collections) Get(repositoryID int) []*Collection {
	var (
		list []*Collection
		name string
	)
	fieldSet := "`id`, `spec`, `name`, `description`, `number_of_records`"
	sql := "SELECT " + fieldSet + " FROM `repository_collection` WHERE `repository` = ? ORDER BY `number_of_records` DESC"
	rows, _ := database.Connection.Query(sql, repositoryID)
	for rows.Next() {
		collection := &Collection{
			Repository: repositoryID,
		}
		rows.Scan(&collection.ID, &collection.Spec, &name, &collection.Description, &collection.NumberOfResources)
		elements := strings.Split(name, ":")
		collection.Name = elements[len(elements)-1]
		collection.Name = strings.Replace(collection.Name, "=", "-", -1)
		list = append(list, collection)
	}
	return list
}
