package action

import (
	"fmt"
	"strconv"

	database "github.com/cristian-sima/Wisply/models/database"
)

// Task executes a single instruction. It can insert to database or requests an URL
// It is coordonated by an operation
// It has a type which denotes the state of the task
type Task struct {
	*Action
	Explication string
	Operation   *Operation
}

// Finish finishes in a normal state the request
// It requires an explication (the error, a success message)
func (task *Task) Finish(explication string) {
	task.Action.Finish()
	task.Explication = explication
	task.updateInDatabase()
}

// CustomFinish finishes the task with a custom result
func (task *Task) CustomFinish(result, explication string) {
	task.Action.ChangeResult(result)
	task.Finish(explication)
}

func (task *Task) updateInDatabase() {
	setClause := "SET end=?, content=?, result=?, is_running=?, explication=?"
	sql := "UPDATE `task` " + setClause + " WHERE id=?"
	stmt, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println("Error #1 when updating the task: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(task.End, task.Content, task.result, strconv.FormatBool(task.IsRunning), task.Explication, strconv.Itoa(task.ID))
	if err != nil {
		fmt.Println("Error #2 when updating the task: ")
		fmt.Println(err)
	}
}
