package action

import (
	"fmt"
	"strconv"

	database "github.com/cristian-sima/Wisply/models/database"
)

// GetProcessesByType returns all the process of a type
func GetProcessesByType(processType string) []*Process {
	return getProcesses(processType)
}

// GetAllProcesses returns an array with all the processes
func GetAllProcesses() []*Process {
	return getProcesses("ALL")
}

func getProcesses(processType string) []*Process {
	var (
		list        []*Process
		whereClause string
	)
	// fields
	processFields := "process.id, process.is_suspended, process.result, process.content, process.start, process.end, process.is_running, process.current_operation"

	fieldList := processFields

	if processType != "ALL" {
		whereClause = " WHERE process.type ='" + processType + "' "
	}

	// the query
	sql := "SELECT " + fieldList + " FROM `process` AS process " + whereClause + " ORDER BY process.id DESC"

	rows, err := database.Connection.Query(sql)
	if err != nil {
		fmt.Println("Problem when getting all the processes: ")
		fmt.Println(err)
	}

	var (
		ID, currentOperationID           int
		start, end                       int64
		isRunning, isSuspended           bool
		content, isRunningString, result string
		operation                        *Operation
	)

	for rows.Next() {
		rows.Scan(&ID, &isSuspended, &result, &content, &start, &end, &isRunningString, &currentOperationID)

		isRunning, err = strconv.ParseBool(isRunningString)

		if err != nil {
			fmt.Println(err)
		}

		if isRunning {
			operation = NewOperation(currentOperationID)
		}

		list = append(list, &Process{
			currentOperation: operation,
			isSuspended:      isSuspended,
			Action: &Action{
				ID:        ID,
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
