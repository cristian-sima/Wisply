package action

import (
	"fmt"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/database"
)

// Operation coordonates a number of many tasks.
// An example of operation may be inserting records into database. It coordonates the task which creates the buffer and the one which inserts the file.
// It is coordonated by a process
type Operation struct {
	*Action
	Process     *Process
	CurrentTask *Task
}

// Finish records in the database that the task is finished.
func (operation *Operation) Finish(content string) {

	operation.End = getCurrentTimestamp()
	operation.Content = content
	operation.IsRunning = false

	stmt, err := wisply.Database.Prepare("UPDATE `operation` SET end=?, content=?, is_running=? WHERE id=?")
	if err != nil {
		fmt.Println("Error 1 when updating the operation: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(operation.End, content, strconv.FormatBool(operation.IsRunning), strconv.Itoa(operation.ID))
	if err != nil {
		fmt.Println("Error 2 when updating the operation: ")
		fmt.Println(err)
	}
}
