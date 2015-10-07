package action

import (
	"fmt"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/database"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// GetAllProcesses returns a list with all available processes
func GetAllProcesses() []*Process {
	var list []*Process
	// fields
	processFields := "process.id, process.result, process.content, process.start, process.end, process.is_running, process.current_operation"
	repositoryFields := "repository.id, repository.name"

	fieldList := processFields + ", " + repositoryFields

	// joins
	joinRepository := "INNER JOIN `repository` AS `repository` ON repository.id = process.repository"
	joins := joinRepository

	// the query
	sql := "SELECT " + fieldList + " FROM `process` AS process " + joins + " ORDER BY process.id DESC"

	rows, err := wisply.Database.Query(sql)
	if err != nil {
		fmt.Println("Problem when getting all the processes: ")
		fmt.Println(err)
	}

	var (
		ID, repID, currentOperationID             int
		start, end                                int64
		isRunning                                 bool
		content, repName, isRunningString, result string
		rep                                       *repository.Repository
		operation                                 *Operation
	)

	for rows.Next() {
		rows.Scan(&ID, &result, &content, &start, &end, &isRunningString, &currentOperationID, &repID, &repName)

		rep = &repository.Repository{
			ID:   repID,
			Name: repName,
		}

		isRunning, err = strconv.ParseBool(isRunningString)

		if err != nil {
			fmt.Println(err)
		}

		if isRunning {
			fmt.Println(repName)
			operation = NewOperation(currentOperationID)
		}

		list = append(list, &Process{
			ID:               ID,
			Repository:       rep,
			currentOperation: operation,
			Action: &Action{
				IsRunning: isRunning,
				Start:     start,
				End:       end,
				Content:   content,
				result:    result,
			},
		})
	}
	return list
}
