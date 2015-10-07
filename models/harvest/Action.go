package harvest

import (
	"fmt"
	"strconv"
	"time"

	wisply "github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
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
	start   string
	end     string
	Content string
}

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
	task.end = getCurrentTimestamp()

	stmt, err := wisply.Database.Prepare("UPDATE `task` SET end=?,content=?,status=? WHERE id=?")
	if err != nil {
		fmt.Println("Error 1 when updating the task: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(task.end, content, status, strconv.Itoa(task.ID))
	if err != nil {
		fmt.Println("Error 2 when updating the task: ")
		fmt.Println(err)
	}
}

// Operation coordonates a number of many tasks.
// An example of operation may be inserting records into database. It coordonates the task which creates the buffer and the one which inserts the file.
// It is coordonated by a process
type Operation struct {
	*Action
	ProcessID   int
	CurrentTask Task
}

// Finish records in the database that the task is finished.
func (operation *Operation) Finish(content string) {

	operation.end = getCurrentTimestamp()
	operation.Content = content

	stmt, err := wisply.Database.Prepare("UPDATE `operation` SET end=?,content=? WHERE id=?")
	if err != nil {
		fmt.Println("Error 1 when updating the operation: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(operation.end, content, strconv.Itoa(operation.ID))
	if err != nil {
		fmt.Println("Error 2 when updating the operation: ")
		fmt.Println(err)
	}
}

// - constructors

func newAction() *Action {
	return &Action{
		start: getCurrentTimestamp(),
	}
}

func newTask(operation Operation) *Task {
	task := &Task{
		OperationID: operation.ID,
		status:      "normal",
		Action:      newAction(),
	}
	columns := "(`start`, `operation`, `status`)"
	values := "(?, ?, ?)"
	sql := "INSERT INTO `task` " + columns + " VALUES " + values
	query, err := wisply.Database.Prepare(sql)
	if err != nil {
		fmt.Println("Error when creating the task:")
		fmt.Println(err)
	}

	query.Exec(task.start, operation.ID, task.GetStatus())

	// find its id

	sql = "SELECT `id` FROM `task` WHERE start=? AND operation=? AND status=?"
	query, err = wisply.Database.Prepare(sql)
	query.QueryRow(task.start, task.OperationID, task.status).Scan(&task.ID)

	if err != nil {
		fmt.Println("Error when selecting the task id:")
		fmt.Println(err)
	}

	return task
}

func newOperation(process Process) *Operation {
	operation := &Operation{
		Action:    newAction(),
		ProcessID: process.ID,
	}
	columns := "(start`, `process`)"
	values := "(?, ?, ?)"
	sql := "INSERT INTO `operation` " + columns + " VALUES " + values

	query, err := wisply.Database.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the operation:")
		fmt.Println(err)
	}

	query.Exec(operation.start, 0)

	// find its ID
	sql = "SELECT `id` FROM `operation` WHERE start=? AND process=?"
	query, err = wisply.Database.Prepare(sql)
	query.QueryRow(operation.start, operation.ProcessID).Scan(&operation.ID)

	if err != nil {
		fmt.Println("Error when selecting the operation id:")
		fmt.Println(err)
	}

	return operation
}

// NewProcess creates a new harvest process
func NewProcess(ID string, controller WisplyController) *Process {
	var remote RemoteRepositoryInterface
	local, _ := repository.NewRepository(ID)

	switch local.Category {
	case "EPrints":
		{
			remote = &EPrintsRepository{
				URL: local.URL,
			}
		}
	}
	db := &databaseManager{}

	process := &Process{
		Action:     newAction(),
		local:      local,
		remote:     remote,
		Controller: controller,
		db:         db,
		Actions:    make(map[string]*Action2),
	}
	// process.SetName("Harvest")
	db.SetManager(process)
	remote.SetManager(process)

	// insert

	columns := "(content`, `repository`, `start`, `is_running`)"
	values := "(?, ?, ?, ?)"
	sql := "INSERT INTO `process` " + columns + " VALUES " + values

	query, err := wisply.Database.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the process:")
		fmt.Println(err)
	}

	query.Exec("harvesting", local.ID, process.start, "1")

	// find its ID
	sql = "SELECT `id` FROM `process` WHERE start=? AND repository=? AND start=? AND is_running=?"
	query, err = wisply.Database.Prepare(sql)
	query.QueryRow("harvesting", local.ID, process.start, "1").Scan(&process.ID)

	if err != nil {
		fmt.Println("Error when selecting the ID of process:")
		fmt.Println(err)
	}

	return process
}

func getCurrentTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
