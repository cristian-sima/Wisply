package action

import (
	"fmt"
	"strconv"

	wisply "github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Process is a top level action which coordonates many operations and communicates with the controller
type Process struct {
	*Action
	ID         int
	Repository *repository.Repository
}

// Finish records in the database that the process is finished.
func (process *Process) Finish() {

	process.End = getCurrentTimestamp()
	process.IsRunning = false

	stmt, err := wisply.Database.Prepare("UPDATE `process` SET end=?, is_running=? WHERE id=?")
	if err != nil {
		fmt.Println("Error 1 when finishing the process: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(process.End, strconv.FormatBool(process.IsRunning), strconv.Itoa(process.ID))
	if err != nil {
		fmt.Println("Error 2 when finishing the process: ")
		fmt.Println(err)
	}
}
