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
func NewAction(isRunning bool, content string) *Action {
	return &Action{
		IsRunning: isRunning,
		Start:     getCurrentTimestamp(),
		Content:   content,
	}
}

// CreateProcess creates a new harvest process
func CreateProcess(repositoryID, content string) *Process {

	local, _ := repository.NewRepository(repositoryID)

	process := &Process{
		Action:     NewAction(true, content),
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

	query.Exec(process.Content, process.Start, local.ID, strconv.FormatBool(process.IsRunning))

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
