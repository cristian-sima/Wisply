package wisply

import "github.com/cristian-sima/Wisply/models/database"

// GetCollections returns the collections
func GetCollections(repositoryID int) []*Collection {
	var list []*Collection
	sql := "SELECT `id`, `spec`, `name`, `description` FROM `repository_collection` WHERE `repository` = ?"
	rows, _ := database.Connection.Query(sql, repositoryID)
	for rows.Next() {
		collection := &Collection{
			Repository: repositoryID,
		}
		rows.Scan(&collection.ID, &collection.Spec, &collection.Name, &collection.Description)
		list = append(list, collection)
	}
	return list
}

// GetRecords returns all the records
func GetRecords(repositoryID int, options database.SQLOptions) []*Record {
	var list []*Record
	sql := "SELECT record.`id`, record.`identifier`, record.`datestamp` FROM `repository_resource` AS record WHERE record.`repository`=? ORDER by record.id DESC " + options.GetLimit()
	rows, _ := database.Connection.Query(sql, repositoryID)
	for rows.Next() {
		record := &Record{}
		rows.Scan(&record.ID, &record.identifier, &record.timestamp)
		sql2 := "SELECT `resource_key`, `value` FROM `resource_key` WHERE `resource`=? "
		rows2, _ := database.Connection.Query(sql2, record.identifier)
		keys := &RecordKeys{}
		for rows2.Next() {
			var name, value string
			rows2.Scan(&name, &value)
			keys.Add(name, value)
		}
		record.Keys = keys
		list = append(list, record)
	}
	return list
}
