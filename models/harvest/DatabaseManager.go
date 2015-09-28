package harvest

// database "github.com/cristian-sima/Wisply/models/database"

import (
	"fmt"
	"strconv"
)

// Manager is a link between controller and repository
type databaseManager struct {
	manager *Manager
}

// StartProcess starts the process
func (db *databaseManager) InsertIdentity(identification *Identificationer) {
	rep := db.manager.GetRepository()
	fmt.Println(rep)
	ID := strconv.Itoa(rep.ID)
	db.log("I insert the identification for repository " + ID + " in database ... ")
	db.log("Identity inserted for " + ID)
}

func (db *databaseManager) SetManager(manager *Manager) {
	db.manager = manager
}

func (db *databaseManager) log(message string) {
	fmt.Println("<--> DatabaseManager: " + message)
}
