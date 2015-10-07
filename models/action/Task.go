package action

import (
	"fmt"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/database"
)

// Task executes a single instruction. It can insert to database or requests an URL
// It is coordonated by an operation
// It has a type which denotes the state of the task
type Task struct {
	*Action
	OperationID int
	status      string // it can be: error, warning, success, normal
}

// ChangeStatus checks if the status is valid and it changes it
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

// Change modifies the status and the content of the task
func (task *Task) Change(status, content string) {
	task.changeStatus(status)
	task.Content = content
}

// Finish records in the database that the task is finished.
func (task *Task) Finish(status, content string) {

	task.Change(status, content)
	task.End = getCurrentTimestamp()

	stmt, err := wisply.Database.Prepare("UPDATE `task` SET end=?,content=?,status=? WHERE id=?")
	if err != nil {
		fmt.Println("Error 1 when updating the task: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(task.End, content, status, strconv.Itoa(task.ID))
	if err != nil {
		fmt.Println("Error 2 when updating the task: ")
		fmt.Println(err)
	}
}
