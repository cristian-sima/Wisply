package action

import (
	"fmt"
	"strconv"
	"time"

	database "github.com/cristian-sima/Wisply/models/database"
)

// CreateProcess creates a new harvest process
func CreateProcess(content string) *Process {

	process := &Process{
		Action:           NewAction(true, content),
		operationConduit: make(chan OperationMessager, 100),
	}
	// insert

	columns := "(`content`, `start`, `is_running`)"
	values := "(?, ?, ?)"
	sql := "INSERT INTO `process` " + columns + " VALUES " + values

	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the process:")
		fmt.Println(sql)
		fmt.Println(err)
	}

	query.Exec(process.Content, process.Start, strconv.FormatBool(process.IsRunning))

	// find its ID
	sql = "SELECT `id` FROM `process` WHERE content=? AND start=? AND is_running=?"
	query, err = database.Connection.Prepare(sql)
	query.QueryRow(process.Content, process.Start, strconv.FormatBool(process.IsRunning)).Scan(&process.ID)

	if err != nil {
		fmt.Println("Error when selecting the ID of process:")
		fmt.Println(err)
	}

	return process
}

func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}
