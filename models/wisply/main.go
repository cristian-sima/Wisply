package wisply

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cristian-sima/Wisply/models/database"
)

// GetCollections returns the collections
func GetCollections(repositoryID int) []*Collection {
	var (
		list []*Collection
		name string
	)

	sql := "SELECT `id`, `spec`, `name`, `description`, `numberOfRecords` FROM `repository_collection` WHERE `repository` = ? ORDER BY `numberOfRecords` DESC"
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

// GetRecords returns all the records
func GetRecords(repositoryID int, options database.SQLOptions) []*Record {
	start := time.Now()
	var list []*Record

	// select from identifier
	// left join identifier_set ON identifier
	//

	var (
		rows *sql.Rows
		err  error
	)
	fieldList := "record.`identifier`, record.`id`, record.`datestamp`"

	// If no collection has been chosen
	if options.Where["collection"] == "" {
		sql := "SELECT " + fieldList + " FROM `repository_resource` AS record WHERE record.`repository`=? ORDER BY record.id DESC " + options.GetLimit()
		fmt.Println(sql)
		rows, err = database.Connection.Query(sql, repositoryID)
	} else {

		//sql := "SELECT " + fieldList + " FROM `repository_resource` AS record INNER JOIN `identifier_set` ON `record`.`identifier` = `identifier_set`.`identifier` WHERE `identifier_set`.`setSpec` = ? ORDER BY record.id DESC " + options.GetLimit()

		sql := "SELECT DISTINCT `identifier_set`.`identifier`, 0, 0 FROM `identifier_set` WHERE `identifier_set`.`setSpec` = ?  ORDER BY identifier DESC " + options.GetLimit()

		// sql := "SELECT DISTINCT `identifier_set`.`identifier`, `repository_resource`.`id`, `repository_resource`.`datestamp` FROM `identifier_set` INNER JOIN `repository_resource` ON `repository_resource`.`identifier` = `identifier_set`.`identifier` WHERE `identifier_set`.`setSpec` = ?  ORDER BY identifier DESC " + options.GetLimit()

		rows, err = database.Connection.Query(sql, options.Where["collection"])

		fmt.Println(sql)
		fmt.Println()
	}

	if err != nil {
		fmt.Println("error with the records sql")
		fmt.Println(err)
	}

	counter := 0
	for rows.Next() {
		counter++
		record := &Record{}
		rows.Scan(&record.identifier, &record.ID, &record.timestamp)

		if record.ID == 0 {

			sql := "SELECT `id`, `datestamp` FROM `repository_resource` WHERE identifier = ?"
			query, err := database.Connection.Prepare(sql)

			if err != nil {
				fmt.Println(err)
			}

			query.QueryRow(record.identifier).Scan(&record.ID, &record.timestamp)
		}
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
	elapsed := time.Since(start)
	log.Printf("Has taken %s", elapsed)
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
