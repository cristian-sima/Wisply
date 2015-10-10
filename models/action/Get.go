package action

import (
	"fmt"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/database"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// DeleteProcessesOfRepository removes all the processes for a repository
func DeleteProcessesOfRepository(ID int) error {

	processes := GetProcessesByRepository(ID)

	for _, process := range processes {
		process.Delete()
	}
	return nil
}

// DeleteEntireLog deletes all the processes, operations and tasks
func DeleteEntireLog() {
	processes := GetAllProcesses()
	for _, process := range processes {
		process.Delete()
	}
}

// GetAllProcesses returns a list with all available processes
func GetAllProcesses() []*Process {
	return getProcesses("ALL")
}

// GetProcessesByRepository returns the list of processes for a specific repository
func GetProcessesByRepository(ID int) []*Process {
	return getProcesses(strconv.Itoa(ID))
}

// "ALL" for receiving the entire set
func getProcesses(repositoryID string) []*Process {
	var (
		list        []*Process
		whereClause string
	)
	// fields
	processFields := "process.id, process.result, process.content, process.start, process.end, process.is_running, process.current_operation"
	repositoryFields := "repository.id, repository.name"

	fieldList := processFields + ", " + repositoryFields

	// joins
	joinRepository := "INNER JOIN `repository` AS `repository` ON repository.id = process.repository"
	joins := joinRepository

	if repositoryID == "'ALL'" {
		whereClause = " WHERE process.repository = " + repositoryID + " "
	}

	// the query
	sql := "SELECT " + fieldList + " FROM `process` AS process " + joins + whereClause + " ORDER BY process.id DESC"

	fmt.Println(sql)
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
			operation = NewOperation(currentOperationID)
		}

		list = append(list, &Process{
			Repository:       rep,
			currentOperation: operation,
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
