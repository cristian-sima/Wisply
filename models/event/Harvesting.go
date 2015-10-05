package history

import (
	"fmt"

	wisply "github.com/cristian-sima/Wisply/models/database"
)

// Manager manages the operation for storing events
type Manager struct {
}

// Record stores an event
func (history *Manager) Record(event *Event) {

	sqlColumns := "(`repository`, `content`, `operation_type`, `operation_name`)"
	sqlValues := "(?, ?, ?, ?)"
	sql := "INSERT INTO `history_event` " + sqlColumns + " VALUES " + sqlValues

	query, err := wisply.Database.Prepare(sql)
	query.Exec(event.Repository, event.Content, event.OperationType, event.OperationName)

	if err != nil {
		fmt.Println("Hmmm problems when inserting events")
		fmt.Println(err)
	}
}

// GetLastEvents returns a list with the last events
func (history *Manager) GetLastEvents() []GUIEvent {
	var list []GUIEvent
	fieldList := "event.id, event.timestamp, repository.name, event.content, event.operation_name, event.operation_type, event.duration"
	sql := "SELECT " + fieldList + " FROM `history_event` as event JOIN `repository` as repository ON event.repository = repository.id ORDER by `id` DESC"
	fmt.Println(sql)
	rows, err := wisply.Database.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var (
			repository, timestamp, content, operationName, operationType string
			id                                                           int
			duration                                                     float32
		)
		rows.Scan(&id, &timestamp, &repository, &content, &operationName, &operationType, &duration)
		list = append(list, GUIEvent{
			Event: Event{
				ID:            id,
				Timestamp:     timestamp,
				Content:       content,
				OperationName: operationName,
				OperationType: operationType,
				Duration:      duration,
			},
			Repository: repository,
		})
	}
	return list
}
