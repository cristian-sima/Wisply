package harvest

// database "github.com/cristian-sima/Wisply/models/database"

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
)

// Manager is a link between controller and repository
type databaseManager struct {
	manager *Process
}

// IDENTITY

// StartProcess starts the process
func (db *databaseManager) InsertIdentity(identification *Identificationer) {

	ID := strconv.Itoa(db.manager.GetRepository().ID)

	db.log("I insert the identification for repository " + ID + " in database ... ")
	db.clearIdentification()
	db.insertIdentificationDetails(identification)
	db.insertEmails(identification)
}

func (db *databaseManager) clearIdentification() {
	var sql string
	sql = "DELETE from `repository_identification` WHERE repository=?"
	query, _ := database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(db.manager.GetRepository().ID))

	sql = "DELETE from `repository_email` WHERE repository=?"
	query, _ = database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(db.manager.GetRepository().ID))
}

func (db *databaseManager) insertIdentificationDetails(identification *Identificationer) {

	var sql string

	ID := strconv.Itoa(db.manager.GetRepository().ID)

	sqlColumns := "(`repository`, `protocol_version`, `earliest_datestamp`, `delete_policy`, `granularity`)"
	sqlValues := "(?, ?, ?, ?, ?)"
	sql = "INSERT INTO `repository_identification` " + sqlColumns + " VALUES " + sqlValues

	query, err := database.Database.Prepare(sql)
	query.Exec(ID, (*identification).GetProtocol(), (*identification).GetEarliestDatestamp(), (*identification).GetDeletedRecord(), (*identification).GetGranularity())

	if err != nil {
		db.log("Hmmm problems when inserting details identification")
	}
}

func (db *databaseManager) insertEmails(identification *Identificationer) {

	ID := strconv.Itoa(db.manager.GetRepository().ID)

	emails := (*identification).GetAdminEmails()

	for _, email := range emails {

		sqlColumns := "(`repository`, `email`)"
		sqlValues := "(?, ?)"
		sql := "INSERT INTO `repository_email` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Database.Prepare(sql)
		query.Exec(ID, email)

		if err != nil {
			db.log("Hmmm problems when inserting emails identification")
		}
	}

}

//  FORMATS

// InsertFormats inserts the formats in the database
func (db databaseManager) InsertFormats(formats []Formater) {
	db.clearFormats()
	db.insertFormats(formats)
}

func (db databaseManager) insertFormats(formats []Formater) {
	var sql string
	ID := strconv.Itoa(db.manager.GetRepository().ID)
	for _, format := range formats {
		sqlColumns := "(`repository`, `md_schema`, `namespace`, `prefix`)"
		sqlValues := "(?, ?, ?, ?)"
		sql = "INSERT INTO `repository_format` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Database.Prepare(sql)
		query.Exec(ID, format.GetSchema(), format.GetNamespace(), format.GetPrefix())

		if err != nil {
			db.log("Hmmm problems when inserting formats")
		}
	}
}

func (db *databaseManager) clearFormats() {
	var sql string
	sql = "DELETE from `repository_format` WHERE repository=?"
	query, _ := database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(db.manager.GetRepository().ID))
}

// COLLECTIONS

func (db databaseManager) InsertCollections(collections []Collection) {
	var sql string
	ID := strconv.Itoa(db.manager.GetRepository().ID)
	for _, collection := range collections {
		sqlColumns := "(`repository`, `name`, `spec`)"
		sqlValues := "(?, ?, ?)"
		sql = "INSERT INTO `repository_collection` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Database.Prepare(sql)
		query.Exec(ID, collection.GetName(), collection.GetSpec())

		if err != nil {
			db.log("Hmmm problems when inserting collections")
		}
	}
}

func (db *databaseManager) ClearCollections() {
	var sql string
	sql = "DELETE from `repository_collection` WHERE repository=?"
	query, _ := database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(db.manager.GetRepository().ID))
}

// Records

func (db databaseManager) InsertRecords(records []Record) {
	var sql string
	ID := strconv.Itoa(db.manager.GetRepository().ID)
	for _, record := range records {

		sqlColumns := "(`repository`, `identifier`, `datestamp`)"
		sqlValues := "(?, ?, ?)"
		sql = "INSERT INTO `repository_resource` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Database.Prepare(sql)
		exec, err := query.Exec(ID, record.GetIdentifier(), record.GetDatestamp())

		if err != nil {
			db.log("Hmmm problems when inserting records")
		}
		ID, _ := exec.LastInsertId()
		record.SetID(ID)

		db.saveKeys(&record)
	}
}

func (db *databaseManager) saveKeys(record *Record) {
	db.insertTitles(record)
}

func (db *databaseManager) insertTitles(record *Record) {
	var keys = (*record).GetKeys()
	db.insertKeys(record, keys.Titles, "title")
	db.insertKeys(record, keys.Creators, "creator")
	db.insertKeys(record, keys.Subjects, "subject")
	db.insertKeys(record, keys.Descriptions, "description")
	db.insertKeys(record, keys.Publishers, "publisher")
	db.insertKeys(record, keys.Contributors, "contributor")
	db.insertKeys(record, keys.Dates, "date")
	db.insertKeys(record, keys.Types, "type")
	db.insertKeys(record, keys.Formats, "format")
	db.insertKeys(record, keys.Identifiers, "identifier")
	db.insertKeys(record, keys.Sources, "source")
	db.insertKeys(record, keys.Languages, "language")
	db.insertKeys(record, keys.Relations, "relation")
	db.insertKeys(record, keys.Coverages, "coverage")
	db.insertKeys(record, keys.Rights, "right")
}

func (db *databaseManager) insertKeys(record *Record, keys []string, name string) {

	var sql string
	ID := strconv.Itoa(db.manager.GetRepository().ID)

	for _, value := range keys {
		sqlColumns := "(`repository`, `resource`, `value`, `resource_key`)"
		sqlValues := "(?, ?, ?, ?)"
		sql = "INSERT INTO `resource_key` " + sqlColumns + " VALUES " + sqlValues

		query, err := database.Database.Prepare(sql)
		query.Exec(ID, (*record).GetID(), value, name)

		if err != nil {
			db.log("Hmmm problems when inserting records")
		}
	}

}

func (db *databaseManager) ClearRecords() {
	var sql string
	sql = "DELETE from `repository_resource` WHERE repository=?"
	query, _ := database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(db.manager.GetRepository().ID))

	// clear keys

	sql = "DELETE from `resource_key` WHERE repository=?"
	query, _ = database.Database.Prepare(sql)
	query.Exec(strconv.Itoa(db.manager.GetRepository().ID))
}

// ---

func (db *databaseManager) SetManager(manager *Process) {
	db.manager = manager
}

func (db *databaseManager) log(message string) {
	fmt.Println("<--> DatabaseManager: " + message)
}
