package wisply

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Storage represents an overview of the remote Repository
// It contains information about remote repository, its data and the courses
type Storage struct {
	Remote *repository.Repository
}

// GetCollections returns a list of the collections
func (storage *Storage) GetCollections() []*Collection {

	var (
		list []*Collection
		name string
	)
	fieldSet := "`id`, `spec`, `name`, `description`, `numberOfRecords`"
	sql := "SELECT " + fieldSet + " FROM `repository_collection` WHERE `repository` = ? ORDER BY `numberOfRecords` DESC"
	rows, _ := database.Connection.Query(sql, storage.Remote.ID)
	for rows.Next() {
		collection := &Collection{
			Repository: storage.Remote.ID,
		}
		rows.Scan(&collection.ID, &collection.Spec, &name, &collection.Description, &collection.NumberOfResources)
		elements := strings.Split(name, ":")
		collection.Name = elements[len(elements)-1]
		collection.Name = strings.Replace(collection.Name, "=", "-", -1)
		list = append(list, collection)
	}
	return list
}

// NewStorage loads and creates a storage based on the remote repository
func NewStorage(remote *repository.Repository) *Storage {
	return &Storage{
		Remote: remote,
	}
}
