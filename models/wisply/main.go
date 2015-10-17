package wisply

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
)

// GetCollections returns the collections
func GetCollections(repositoryID int) []*Collection {
	var (
		list []*Collection
		name string
	)

	sql := "SELECT `id`, `spec`, `name`, `description` FROM `repository_collection` WHERE `repository` = ?"
	rows, _ := database.Connection.Query(sql, repositoryID)
	for rows.Next() {
		collection := &Collection{
			Repository: repositoryID,
		}

		rows.Scan(&collection.ID, &collection.Spec, &name, &collection.Description)

		elements := strings.Split(name, ":")

		collection.Name = elements[len(elements)-1]

		collection.Name = strings.Replace(collection.Name, "=", "-", -1)

		list = append(list, collection)
	}
	return list
}

// GetRecords returns all the records
func GetRecords(repositoryID int, options database.SQLOptions) []*Record {
	var list []*Record

	// select from identifier
	// left join identifier_set ON identifier
	//

	var (
		rows *sql.Rows
		err  error
	)
	fieldList := "record.`id`, record.`identifier`, record.`datestamp` FROM `repository_resource`"

	fmt.Println("Collection is ")

	fmt.Println("[" + options.Where["collection"] + "]")

	// If no collection has been chosen
	if options.Where["collection"] == "" {
		sql := "SELECT record.`id`, record.`identifier`, record.`datestamp` FROM `repository_resource` AS record WHERE record.`repository`=? ORDER by record.id DESC " + options.GetLimit()
		rows, err = database.Connection.Query(sql, repositoryID)
		fmt.Println("all")
	} else {
		sql := "SELECT " + fieldList + " AS record INNER JOIN `identifier_set` ON record.identifier = identifier_set.identifier WHERE `identifier_set`.setSpec LIKE ? ORDER by record.id DESC " + options.GetLimit()
		rows, err = database.Connection.Query(sql, "%"+options.Where["collection"]+"%")
		fmt.Println(sql)
		fmt.Println("collection")
	}
	if err != nil {
		fmt.Println("error with the records sql")
		fmt.Println(err)
	}

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
	return list
}

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
	sql = "DELETE FROM `resource_key` WHERE `repository` = ?"
	query, _ = database.Connection.Prepare(sql)
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
	sql = "DELETE FROM `identifier_set` WHERE `repository` = ?"
	query, _ = database.Connection.Prepare(sql)
	query.Exec(repositoryID)
}
