package wisply

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/database"
)

// ClearRepository deletes all the resources, formats, collections, identifiers, emails and the identificaiton details
func ClearRepository(repositoryID int) {
	deleteResources(repositoryID)
	deleteFormats(repositoryID)
	deleteCollections(repositoryID)
	deleteEmails(repositoryID)
	deleteIdentification(repositoryID)
	deleteIdentifiers(repositoryID)
	updateRepository(repositoryID)
}

func updateRepository(repositoryID int) {
	sql := "UPDATE `repository` SET `lastProcess`=0, `status`='unverified' WHERE `id`=?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}

func deleteResources(repositoryID int) {
	sql := "DELETE FROM `repository_resource` WHERE `repository` = ?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}

func deleteFormats(repositoryID int) {
	sql := "DELETE FROM `repository_format` WHERE `repository` = ?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}

func deleteCollections(repositoryID int) {
	sql := "DELETE FROM `repository_collection` WHERE `repository` = ?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}

func deleteEmails(repositoryID int) {
	sql := "DELETE FROM `repository_email` WHERE `repository` = ?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}

func deleteIdentification(repositoryID int) {
	sql := "DELETE FROM `repository_identification` WHERE `repository` = ?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}

func deleteKeys(repositoryID int) {
	sql := "DELETE FROM `repository_key` WHERE `repository` = ?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}

func deleteIdentifiers(repositoryID int) {
	sql := "DELETE FROM `identifier` WHERE `repository` = ?"
	query, _ := database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}

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

	fmt.Println(sql)
	rows, _ := database.Connection.Query(sql, repositoryID)

	counter := 0
	for rows.Next() {
		counter++
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
	fmt.Println(counter)
	return list
}
