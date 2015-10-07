package action

// GetAllProcesses returns a list with all available processes
func GetAllProcesses() []string {
	var list []string
	return list
	// fieldList := "event.id as eventID, event.timestamp, repository.name, event.content, event.operation_name, event.operation_type, event.duration, repository.id"
	// sql := "SELECT " + fieldList + " FROM `history_event` as `event` JOIN `repository` as `repository` ON event.repository = repository.id ORDER by eventID DESC"
	// fmt.Println(sql)
	// rows, err := wisply.Database.Query(sql)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for rows.Next() {
	// 	var (
	// 		repository, timestamp, content, operationName, operationType string
	// 		id, repositoryID                                             int
	// 		duration                                                     float32
	// 	)
	// 	rows.Scan(&id, &timestamp, &repository, &content, &operationName, &operationType, &duration, &repositoryID)
	// 	list = append(list, GUIEvent{
	// 		Event: Event{
	// 			ID:            id,
	// 			Timestamp:     timestamp,
	// 			Content:       content,
	// 			OperationName: operationName,
	// 			OperationType: operationType,
	// 			Duration:      duration,
	// 			Repository:    repositoryID,
	// 		},
	// 		RepositoryName: repository,
	// 	})
	// }
	// return list
}
