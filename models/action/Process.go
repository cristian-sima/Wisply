package action

import (
	"fmt"
	"strconv"

	database "github.com/cristian-sima/Wisply/models/database"
)

// Process is a top level action which coordonates many operations and communicates with the controller
type Process struct {
	*Action
	ID               int
	isSuspended      bool
	currentOperation *Operation
	operationConduit chan OperationMessager
}

// Finish records in the database that the process is finished.
func (process *Process) Finish() {
	process.Action.Finish()
	process.updateInDatabase()
}

// GetOperationConduit returns the channel for communicating with operations
func (process *Process) GetOperationConduit() chan OperationMessager {
	return process.operationConduit
}

func (process *Process) updateInDatabase() {
	stmt, err := database.Connection.Prepare("UPDATE `process` SET is_suspended=?, end=?, is_running=?, current_operation=?, result=? WHERE id=?")
	if err != nil {
		fmt.Println("Error 1 when finishing the process: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(strconv.FormatBool(process.isSuspended), process.End, strconv.FormatBool(process.IsRunning), process.currentOperation.ID, process.result, strconv.Itoa(process.ID))
	if err != nil {
		fmt.Println("Error 2 when finishing the process: ")
		fmt.Println(err)
	}
}

// ChangeCurrentOperation sets the current operation
func (process *Process) ChangeCurrentOperation(operation *Operation) {
	process.currentOperation = operation
	process.updateInDatabase()
}

// Suspend sets the process as suspended
func (process *Process) Suspend() {
	process.isSuspended = true
	process.updateInDatabase()
}

// IsSuspended verifiies if the process is suspended
func (process *Process) IsSuspended() bool {
	return process.isSuspended
}

// GetCurrentOperation returns the current operation
func (process *Process) GetCurrentOperation() *Operation {
	return process.currentOperation
}

// CreateOperation creates a new operation
func (process *Process) CreateOperation(content string) *Operation {
	operation := &Operation{
		Action:  NewAction(true, content),
		Process: process,
		conduit: make(chan Messager, 100),
	}
	columns := "(`start`, `process`, `content`)"
	values := "(?, ?, ?)"
	sql := "INSERT INTO `operation` " + columns + " VALUES " + values

	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the operation:")
		fmt.Println(err)
	}

	query.Exec(operation.Start, process.ID, operation.Content)

	// find its ID
	sql = "SELECT `id` FROM `operation` WHERE start=? AND process=? AND is_running=?"
	query, err = database.Connection.Prepare(sql)
	query.QueryRow(operation.Start, operation.Process.ID, strconv.FormatBool(operation.IsRunning)).Scan(&operation.ID)

	if err != nil {
		fmt.Println("Error when selecting the operation id:")
		fmt.Println(err)
	}

	return operation
}

// GetOperations it returns the list of operations
func (process *Process) GetOperations() []*Operation {

	// fields
	fieldList := "operation.id, operation.content, operation.start, operation.end, operation.is_running, operation.current_task, operation.result"

	// the query
	sql := "SELECT " + fieldList + " FROM `operation` AS operation WHERE process=? ORDER BY operation.id DESC"

	rows, err := database.Connection.Query(sql, process.Action.ID)
	if err != nil {
		fmt.Println("Problem when getting all the operations of the process: ")
		fmt.Println(err)
	}

	var (
		list                             []*Operation
		ID, currentTaskID                int
		start, end                       int64
		isRunning                        bool
		content, isRunningString, result string
		task                             *Task
	)

	for rows.Next() {
		rows.Scan(&ID, &content, &start, &end, &isRunningString, &currentTaskID, &result)

		isRunning, err = strconv.ParseBool(isRunningString)

		if err != nil {
			fmt.Println(err)
		}

		if isRunning {
			task = NewTask(currentTaskID)
		}

		list = append(list, &Operation{
			CurrentTask: task,
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

// Delete deletes the process along with the tasks and operations
func (process *Process) Delete() {

	sql := "DELETE FROM `process` WHERE id=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println("Delete 1 error for process:")
		fmt.Println(err)
	}
	query.Exec(process.Action.ID)

	// operations

	sql = "DELETE FROM `operation` WHERE process=?"
	query, err = database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println("Delete 2 error for process:")
		fmt.Println(err)
	}
	query.Exec(process.Action.ID)

	// tasks

	sql = "DELETE FROM `task` WHERE process=?"
	query, err = database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println("Delete 3 error for process:")
		fmt.Println(err)
	}
	query.Exec(process.Action.ID)

}
