// Package wisply represents a Wisply repository
// It contains the metadata from the remote repository and the processed data
package wisply

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/cristian-sima/Wisply/models/database"
)

// GetRecords returns all the records
func GetRecords(repositoryID int, options database.SQLOptions) []*Record {
	start := time.Now()
	var list []*Record

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

		// this takes too long
		//sql := "SELECT " + fieldList + " FROM `repository_resource` AS record INNER JOIN `identifier_set` ON `record`.`identifier` = `identifier_set`.`identifier` WHERE `identifier_set`.`setSpec` = ? ORDER BY record.id DESC " + options.GetLimit()

		sql := "SELECT DISTINCT `identifier_set`.`identifier`, 0, 0 FROM `identifier_set` WHERE `identifier_set`.`setSpec` = ?  ORDER BY identifier DESC " + options.GetLimit()

		rows, err = database.Connection.Query(sql, options.Where["collection"])

		fmt.Println(sql)
		fmt.Println()
	}

	if err != nil {
		fmt.Println("Error #1 with Wisply repository records")
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

// GetRecordByIdentifier finds the record in the database by identifier
func GetRecordByIdentifier(identifier string) Record {
	record := Record{}

	record.identifier = identifier

	sql := "SELECT `id`, `datestamp` FROM `repository_resource` WHERE identifier = ?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println(err)
	}

	query.QueryRow(record.identifier).Scan(&record.ID, &record.timestamp)

	sql2 := "SELECT `resource_key`, `value` FROM `resource_key` WHERE `resource`=? "
	rows2, _ := database.Connection.Query(sql2, record.identifier)

	keys := &RecordKeys{}
	for rows2.Next() {
		var name, value string
		rows2.Scan(&name, &value)
		keys.Add(name, value)
	}
	record.Keys = keys
	return record
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
	setClause := "SET `lastProcess`=0, `status`='unverified'"
	sql := "UPDATE `repository` " + setClause + " WHERE `id`=?"
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
