package history

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/database"
)

// Manager manages the operation for storing events
type Manager struct {
}

// Record stores an event
func (history *Manager) Record(event *Event) {

	sqlColumns := "(`repository`, `content`, `operation_type`, `operation_name`)"
	sqlValues := "(?, ?, ?, ?)"
	sql := "INSERT INTO `history_event` " + sqlColumns + " VALUES " + sqlValues

	query, err := database.Database.Prepare(sql)
	query.Exec(event.Repository, event.Content, event.OperationType, event.OperationName)

	if err != nil {
		fmt.Println("Hmmm problems when inserting events")
		fmt.Println(err)
	}
}
