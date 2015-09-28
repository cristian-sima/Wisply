package harvest

// database "github.com/cristian-sima/Wisply/models/database"

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
)

// Manager is a link between controller and repository
type databaseManager struct {
	manager *Manager
}

// StartProcess starts the process
func (db *databaseManager) InsertIdentity(identification *Identificationer) {

	ID := strconv.Itoa(db.manager.GetRepository().ID)

	db.log("I insert the identification for repository " + ID + " in database ... ")
	db.clearIdentification()
	db.insertIdentificationDetails(identification)
	db.insertEmails(identification)
	db.log("Identity inserted for " + ID)
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
		fmt.Println("Hmmm problems when inserting details identification")
	}
}

func (db *databaseManager) insertEmails(identification *Identificationer) {

	ID := strconv.Itoa(db.manager.GetRepository().ID)

	db.log("I insert the emails for repository " + ID + " in database ... ")

	emails := (*identification).GetAdminEmails()

	for _, email := range emails {

		sqlColumns := "(`repository`, `email`)"
		sqlValues := "(?, ?)"
		sql := "INSERT INTO `repository_email` " + sqlColumns + " VALUES " + sqlValues

		fmt.Println(sql)
		query, err := database.Database.Prepare(sql)
		query.Exec(ID, email)

		if err != nil {
			fmt.Println("Hmmm problems when inserting emails identification")
		}
	}

}

func (db *databaseManager) SetManager(manager *Manager) {
	db.manager = manager
}

func (db *databaseManager) log(message string) {
	fmt.Println("<--> DatabaseManager: " + message)
}
