package harvest

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/database"
)

// HistoryManager manages the operation for storing events
type HistoryManager struct {
}

// Record stores an event
func (history *HistoryManager) Record(msg *Message) {

	sqlColumns := "(`repository`, `content`)"
	sqlValues := "(?, ?)"
	sql := "INSERT INTO `history_event` " + sqlColumns + " VALUES " + sqlValues

	query, err := database.Database.Prepare(sql)
	query.Exec(msg.Repository, msg.Content)

	if err != nil {
		fmt.Println("Hmmm problems when inserting events")
		fmt.Println(err)
	}
}
