// Package wisply contains the metadata from the remote repository and the
// processed data
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
	fieldList := "record.`identifier`, record.`id`, record.`datestamp`, record.`repository`, record.`is_visible`"

	// If no collection has been chosen
	if options.Where["collection"] == "" {
		sql := "SELECT " + fieldList + " FROM `repository_resource` AS record WHERE record.`repository`=? ORDER BY record.id DESC " + options.GetLimit()
		rows, err = database.Connection.Query(sql, repositoryID)
	} else {
		sql := "SELECT DISTINCT `identifier_set`.`identifier`, 0, 0, 0, 0 FROM `identifier_set` WHERE `identifier_set`.`set_spec` = ?  ORDER BY identifier DESC " + options.GetLimit()
		rows, err = database.Connection.Query(sql, options.Where["collection"])
	}

	if err != nil {
		fmt.Println("Error #1 with Wisply repository records")
		fmt.Println(err)
	}

	counter := 0
	for rows.Next() {
		var isVisibleInt int

		counter++
		record := &Record{}

		rows.Scan(&record.Identifier, &record.ID, &record.Timestamp, &record.Repository, &isVisibleInt)
		record.IsVisible = database.GetBoolFromInt(isVisibleInt)

		if record.ID == 0 {
			sql := "SELECT `id`, `datestamp`, `repository`, `is_visible` FROM `repository_resource` WHERE identifier = ?"
			query, err := database.Connection.Prepare(sql)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(record.Identifier)
			err = query.QueryRow(record.Identifier).Scan(&record.ID, &record.Timestamp, &record.Repository, &isVisibleInt)
			record.IsVisible = database.GetBoolFromInt(isVisibleInt)
			fmt.Println(err)
		}
		sql2 := "SELECT `resource_key`, `value` FROM `resource_key` WHERE `resource`=? "
		rows2, _ := database.Connection.Query(sql2, record.Identifier)

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

// GetRecordByID finds the record in the database by wisply ID
func GetRecordByID(value string) (Record, error) {
	return getRecord("`id`", value)
}

// GetRecordByIdentifier finds the record in the database by identifier
func GetRecordByIdentifier(value string) Record {
	record, _ := getRecord("`identifier`", value)
	return record
}

func getRecord(byFild, value string) (Record, error) {
	var isVisibleInt int
	record := Record{}

	sql := "SELECT `id`, `datestamp`, `identifier`, `repository`, `is_visible` FROM `repository_resource` WHERE " + byFild + " = ?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		return record, err
	}

	query.QueryRow(value).Scan(&record.ID, &record.Timestamp, &record.Identifier, &record.Repository, &isVisibleInt)

	record.IsVisible = database.GetBoolFromInt(isVisibleInt)

	sql2 := "SELECT `resource_key`, `value` FROM `resource_key` WHERE `resource`=? "
	rows2, keyErr := database.Connection.Query(sql2, record.Identifier)

	if keyErr != nil {
		return record, keyErr
	}

	keys := &RecordKeys{}
	for rows2.Next() {
		var name, value string
		rows2.Scan(&name, &value)
		keys.Add(name, value)
	}
	record.Keys = keys
	return record, nil
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
	setClause := "SET `last_process`=0, `status`='unverified'"
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
