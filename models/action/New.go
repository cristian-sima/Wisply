package action

import (
	"fmt"
	"strconv"
	"time"

	wisply "github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// - constructors

// NewAction creates a new action
func NewAction(isRunning bool) *Action {
	return &Action{
		IsRunning: isRunning,
		Start:     getCurrentTimestamp(),
	}
}

// NewTask creates a new operation
func NewTask(operation Operation) *Task {
	task := &Task{
		OperationID: operation.ID,
		status:      "normal",
		Action:      NewAction(true),
	}
	columns := "(`start`, `operation`, `status`)"
	values := "(?, ?, ?)"
	sql := "INSERT INTO `task` " + columns + " VALUES " + values
	query, err := wisply.Database.Prepare(sql)
	if err != nil {
		fmt.Println("Error when creating the task:")
		fmt.Println(err)
	}

	query.Exec(task.Start, operation.ID, task.GetStatus())

	// find its id

	sql = "SELECT `id` FROM `task` WHERE start=? AND operation=? AND status=? AND is_running=?"
	query, err = wisply.Database.Prepare(sql)
	query.QueryRow(task.Start, task.OperationID, task.status, strconv.FormatBool(task.IsRunning)).Scan(&task.ID)

	if err != nil {
		fmt.Println("Error when selecting the task id:")
		fmt.Println(err)
	}

	return task
}

// NewOperation creates a new operation
func NewOperation(process Process) *Operation {
	operation := &Operation{
		Action:    NewAction(true),
		ProcessID: process.ID,
	}
	columns := "(`start`, `process`)"
	values := "(?, ?, ?)"
	sql := "INSERT INTO `operation` " + columns + " VALUES " + values

	query, err := wisply.Database.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the operation:")
		fmt.Println(err)
	}

	query.Exec(operation.Start, 0)

	// find its ID
	sql = "SELECT `id` FROM `operation` WHERE start=? AND process=? AND is_running=?"
	query, err = wisply.Database.Prepare(sql)
	query.QueryRow(operation.Start, operation.ProcessID, strconv.FormatBool(operation.IsRunning)).Scan(&operation.ID)

	if err != nil {
		fmt.Println("Error when selecting the operation id:")
		fmt.Println(err)
	}

	return operation
}

// NewProcess creates a new harvest process
func NewProcess(repositoryID, content string) *Process {

	local, _ := repository.NewRepository(repositoryID)

	process := &Process{
		Action:     NewAction(true),
		Repository: local,
	}
	// insert

	columns := "(`content`, `start`, `repository`, `is_running`)"
	values := "(?, ?, ?, ?)"
	sql := "INSERT INTO `process` " + columns + " VALUES " + values

	query, err := wisply.Database.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the process:")
		fmt.Println(sql)
		fmt.Println(err)
	}

	query.Exec(content, process.Start, local.ID, strconv.FormatBool(process.IsRunning))

	// find its ID
	sql = "SELECT `id` FROM `process` WHERE content=? AND start=? AND repository=? AND is_running=?"
	query, err = wisply.Database.Prepare(sql)
	query.QueryRow("harvesting", process.Start, local.ID, strconv.FormatBool(process.IsRunning)).Scan(&process.ID)

	if err != nil {
		fmt.Println("Error when selecting the ID of process:")
		fmt.Println(err)
	}

	fmt.Println(process.ID)
	return process
}

func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}
