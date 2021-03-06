package action

import (
	"fmt"
	"strconv"

	database "github.com/cristian-sima/Wisply/models/database"
)

// Operation coordonates a number of many tasks.
// An example of operation may be inserting records into database.
// It coordonates the task which creates the buffer and the one which inserts the file.
// It is coordonated by a process
type Operation struct {
	*Action
	Process     *Process
	CurrentTask *Task
	conduit     chan Messager
}

// Finish finishes in a normal state the request
func (operation *Operation) Finish() {
	operation.Action.Finish()
	operation.updateInDatabase()
}

// TellProcess tells to the process a message
func (operation *Operation) TellProcess(message Message) {
	var channel = operation.Process.GetOperationConduit()

	if channel == nil {
		panic("No channel for process")
	} else {
		msg := OperationMessage{
			Message:   &message,
			Operation: operation,
		}
		channel <- &msg
	}
}

// GetTaskConduit returns the channel for communicating with tasks
func (operation *Operation) GetTaskConduit() chan Messager {
	return operation.conduit
}

func (operation *Operation) updateInDatabase() {
	setClause := "SET end=?, content=?, is_running=?, result=?"
	sql := "UPDATE `operation` " + setClause + " WHERE id=?"
	stmt, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println("Error #1 when updating the operation: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(operation.End, operation.Content, strconv.FormatBool(operation.IsRunning), operation.result, strconv.Itoa(operation.ID))
	if err != nil {
		fmt.Println("Error #2 when updating the operation: ")
		fmt.Println(err)
	}
}

// CreateTask creates a new operation
func (operation *Operation) CreateTask(content string) *Task {
	task := &Task{
		Operation: operation,
		Action:    NewAction(true, content),
	}
	columns := "(`process`, `start`, `operation`, `result`, `content`)"
	values := "(?, ?, ?, ?, ?)"
	sql := "INSERT INTO `task` " + columns + " VALUES " + values
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println("Error #1 when creating the task:")
		fmt.Println(err)
	}

	_, err = query.Exec(operation.Process.ID, task.Start, operation.ID, task.GetResult(), task.Content)
	if err != nil {
		fmt.Println("Error #2 when creating the task:")
		fmt.Println(err)
	}

	// find its id
	whereClause := "start=? AND operation=? AND result=? AND is_running=? AND content=?"
	sql = "SELECT `id` FROM `task` WHERE " + whereClause
	query, err = database.Connection.Prepare(sql)
	query.QueryRow(task.Start, task.Operation.ID, task.result, strconv.FormatBool(task.IsRunning), task.Content).Scan(&task.ID)

	if err != nil {
		fmt.Println("Error when selecting the task id:")
		fmt.Println(err)
	}

	return task
}

// GetTasks returns the list of tasks
func (operation *Operation) GetTasks() []*Task {

	// fields
	fieldList := "task.id, task.content, task.start, task.end, task.is_running, task.result, task.explication"

	// the query
	sql := "SELECT " + fieldList + " FROM `task` WHERE operation=? ORDER BY task.id DESC"

	rows, err := database.Connection.Query(sql, operation.Action.ID)
	if err != nil {
		fmt.Println("Problem when getting all the tasks of the operation (Error #5): ")
		fmt.Println(err)
	}

	var (
		list                                          []*Task
		ID                                            int
		start, end                                    int64
		isRunning                                     bool
		content, isRunningString, result, explication string
	)

	for rows.Next() {
		rows.Scan(&ID, &content, &start, &end, &isRunningString, &result, &explication)

		isRunning, err = strconv.ParseBool(isRunningString)

		if err != nil {
			fmt.Println(err)
		}

		list = append(list, &Task{
			Explication: explication,
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
