package harvest

import (
	"fmt"

	wisply "github.com/cristian-sima/Wisply/models/database"
)

// Actioner ... defines the set of methods to be implemented by an action
type Actioner interface {
	Go()
	Finish()
}

// Action is the most basic type. It has a starting and ending timestamps and a content
type Action struct {
	Actioner
	ID      int
	Start   string
	End     string
	Content string
}

// Task executes a single instruction. It can insert to database or requests an URL
// It is coordonated by an operation
// It has a type which denotes the state of the task
type Task struct {
	Action
	Operation Operation
	status    string // it can be: error, warning, success, normal
}

// changeStatus checks if the status is valid and it changes it
func (task *Task) changeStatus(status string) {
	if status != "error" &&
		status != "warning" &&
		status != "success" &&
		status != "normal" {
		fmt.Println("Task change status error. This status is not valid: " + status)
	}
	task.status = status
}

// GetStatus returns the status of the task
func (task *Task) GetStatus() string {
	return task.status
}

// update the task in the database
func (task *Task) update() {

}

func (action *Action) update() {
}

// Operation coordonates a number of many tasks.
// An example of operation may be inserting records into database. It coordonates the task which creates the buffer and the one which inserts the file.
// It is coordonated by a process
type Operation struct {
	Action
	Manager Process
}

// - constructors

func newTask(operation Operation) *Task {
	task := &Task{
		Operation: operation,
		status:    "normal",
	}
	columns := "(`id`, `start`, `operation`, `status`)"
	values := "(?, ?, ?, ?)"
	sql := "INSERT INTO `task` " + columns + " VALUES " + values

	query, err := wisply.Database.Prepare(sql)

	if err != nil {
		fmt.Println("Hmmm problems when inserting this task:")
		fmt.Println(task)
	}

	query.Exec(task.ID, task.Start, operation.ID, task.GetStatus())

	return task
}
