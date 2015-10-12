package harvest

import (
	"fmt"
	"strconv"

	action "github.com/cristian-sima/Wisply/models/action"
	database "github.com/cristian-sima/Wisply/models/database"
	remote "github.com/cristian-sima/Wisply/models/harvest/remote"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Process is a link between controller and repository
type Process struct {
	*action.Process
	HarvestID  int
	repository *repository.Repository
	remote     remote.RepositoryInterface
	controller Controller
}

// Start starts the process
func (process *Process) Start() {
	go process.run()
	go process.verify()
}

func (process *Process) run() {
	fmt.Println("Run process...")
	for {
		select {
		case message := <-process.Process.GetOperationConduit():
			fmt.Println("The process has received: ")
			fmt.Println(message)
			if message.GetName() == "operation-failed" {
				fmt.Println("Ii spune")
				process.processFails()
			} else {
				process.chooseAction(message)
			}
		}
	}
}

func (process *Process) chooseAction(message action.OperationMessager) {
	fmt.Println("Process: I got: ")
	fmt.Println(message)

	switch message.GetOperation().Content {
	case "Harvest Verification":
		if message.GetValue() == "normal" {
			go process.identify()
		} else {
			process.processFails()
		}
		break
	case "Harvest Identifying":
		if message.GetValue() == "normal" {
			go process.harvestFormats()
		} else {
			process.processFails()
		}
		break
	case "Harvest Formats":
		if message.GetValue() == "normal" {
			go process.harvestCollections()
		} else {
			process.processFails()
		}
		break
	case "Harvest Collections":
		if message.GetValue() == "normal" {
			go process.harvestRecords()
		} else {
			process.processFails()
		}
		break
	case "Harvest Records":
		if message.GetValue() == "normal" {
			go process.harvestIdentifiers()
		} else {
			process.processFails()
		}
	case "Harvest Identifiers":
		if message.GetValue() != "normal" {
			process.processFails()
		}
		process.ChangeRepositoryStatus("ok")
		process.succeeded()
		break
	default:
		fmt.Println("No such operation: " + message.GetOperation().Content)
	}
}

// It starts the process from a certain stage
func (process *Process) startFrom(name string) {
	fmt.Println("I am starting from: " + name)
	switch name {
	case "Harvest Verification":
		go process.verify()
		break
	case "Harvest Identifying":
		go process.identify()
		break
	case "Harvest Formats":
		go process.harvestFormats()
		break
	case "Harvest Identifiers":
		go process.harvestIdentifiers()
		break
	case "Harvest Records":
		go process.harvestRecords()
		break
	case "Harvest Collections":
		go process.harvestCollections()
		break
	}
}

// Stage 1

func (process *Process) verify() {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	verification := newVerificationOperation(process)
	process.ChangeCurrentOperation(verification)
	verification.Start()
}

// Stage 2

func (process *Process) identify() {
	identification := newIdentificationOperation(process)
	process.ChangeCurrentOperation(identification)
	identification.Start()
}

// Stage 3

func (process *Process) harvest() {
	process.ChangeRepositoryStatus("updating")
}

// FORMATS

func (process *Process) harvestFormats() {
	process.harvest()
	harvestingFormats := newHarvestingFormats(process)
	process.ChangeCurrentOperation(harvestingFormats)
	harvestingFormats.Start()
}

// COLLECTIONS

func (process *Process) harvestCollections() {
	process.harvest()
	harvestingCollections := newHarvestingCollections(process)
	process.ChangeCurrentOperation(harvestingCollections)
	harvestingCollections.Start()
}

// RECORDS

func (process *Process) harvestRecords() {
	process.harvest()
	harvestingRecords := newHarvestingRecords(process)
	process.ChangeCurrentOperation(harvestingRecords)
	harvestingRecords.Start()
}

// IDENTIFIERS

func (process *Process) harvestIdentifiers() {
	process.harvest()
	identifiers := newHarvestingIdentifiers(process)
	process.ChangeCurrentOperation(identifiers)
	identifiers.Start()
}

// --- end activity

// SuccessFinish tells the constructor that the process is finished
func (process *Process) succeeded() {
	process.delete()
	process.ChangeResult("success")
	process.Process.Finish()
}

func (process *Process) delete() {
	process.tellController(&Message{
		Name: "process-finished",
	})
}

func (process *Process) processFails() {
	process.ChangeResult("warning")
	process.delete()
	process.Suspend()
	process.ChangeRepositoryStatus("problems")
}

// Suspend stops the process and waits for user
func (process *Process) Suspend() {
	process.Process.Suspend()
}

// GetRepository returns the wisply repository
func (process *Process) GetRepository() *repository.Repository {
	return process.repository
}

// GetRemoteServer returns the interface of a remote repository
func (process *Process) GetRemoteServer() remote.RepositoryInterface {
	return process.remote
}

// ChangeCurrentOperation informs the controller about the change and it calls its father
func (process *Process) ChangeCurrentOperation(operation Operationer) {
	process.Process.ChangeCurrentOperation(operation.GetOperation())
}

// ChangeRepositoryStatus changes the status of local repository
func (process *Process) ChangeRepositoryStatus(newStatus string) {
	fmt.Println("I am telling the controller")
	process.repository.ModifyStatus(newStatus)
	process.tellController(&Message{
		Name:  "repository-status-changed",
		Value: newStatus,
	})
}

// GetToken returns the token for a particular name
func (process *Process) GetToken(name string) string {
	var token string
	// find its ID
	sql := "SELECT `token_" + name + "` FROM `process_harvest` WHERE id=? LIMIT 0,1"
	query, err := database.Connection.Prepare(sql)
	query.QueryRow().Scan(&token)

	if err != nil {
		fmt.Println("Error when selecting the token for " + name + ":")
		fmt.Println(err)
	}
	return token
}

// SaveToken saves the token for a particular name
func (process *Process) SaveToken(name string, token string) {
	stmt, err := database.Connection.Prepare("UPDATE `process_harvest` SET `token_" + name + "`=? WHERE id=?")
	if err != nil {
		fmt.Println("Error 1 when updating the token for " + name + " inside harvesting process: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(token, process.HarvestID)
	if err != nil {
		fmt.Println("Error 2 when updating the token for " + name + " inside harvesting process: ")
		fmt.Println(err)
	}
}

func (process *Process) tellController(simple *Message) {

	if process.controller != nil {

		channel := process.controller.GetConduit()

		msg := &ProcessMessage{
			Repository: process.GetRepository().ID,
			ProcessMessage: action.ProcessMessage{
				Process: process,
				Message: &action.Message{
					Name:  simple.Name,
					Value: simple.Value,
				},
			},
		}
		channel <- msg
	}
}

// Delete deletes the harvest process and calls its parent method
func (process *Process) Delete() {
	DeleteProcess(process.ID)
}

// Recover loads the process in memory and starts executing the last stage
func (process *Process) Recover() {
	var name string
	lastOperation := process.Process.GetCurrentOperation().ID

	// gets the name of the current stage
	sql := "SELECT `content` FROM `operation` WHERE id=? LIMIT 0,1"
	query, err := database.Connection.Prepare(sql)
	query.QueryRow(lastOperation).Scan(&name)

	if err != nil {
		fmt.Println("Error when selecting the harvest id:")
		fmt.Println(err)
	}

	process.Process.Recover()

	// start channel
	go process.run()

	// execute stage
	fmt.Println("Recover")
	process.startFrom(name)
}

// -------------------------------------------- Functions ---

// RecoverProcess gets the information about the current process and recovers it
func RecoverProcess(process *Process, controller Controller) *Process {
	process.controller = controller
	return process
}

// CreateProcess creates a new harvest process
func CreateProcess(ID string, controller Controller) *Process {

	process := &*action.CreateProcess("Harvest")

	harvestProcess := buildProcess(ID, controller, process)

	harvestID := insertHarvestProcess(harvestProcess)

	harvestProcess.HarvestID = harvestID

	return harvestProcess
}

func buildProcess(ID string, controller Controller, process *action.Process) *Process {

	local, _ := repository.NewRepository(ID)

	harvestProcess := &Process{
		Process:    process,
		remote:     getRemoteServer(local),
		controller: controller,
		repository: local,
	}

	return harvestProcess
}

func getRemoteServer(local *repository.Repository) remote.RepositoryInterface {

	var rem remote.RepositoryInterface

	switch local.Category {
	case "EPrints":
		{
			rem = remote.NewEPrints(local)
		}
	}

	return rem
}

func insertHarvestProcess(process *Process) int {

	var harvestID int

	columns := "(`process`, `repository`)"
	values := "(?, ?)"
	sql := "INSERT INTO `process_harvest` " + columns + " VALUES " + values

	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the harvest process:")
		fmt.Println(sql)
		fmt.Println(err)
	}
	query.Exec(process.ID, process.GetRepository().ID)

	// find its ID
	sql = "SELECT `id` FROM `process_harvest` WHERE process=? AND repository=? ORDER by id LIMIT 0,1"
	query, err = database.Connection.Prepare(sql)
	query.QueryRow(process.ID, process.GetRepository().ID).Scan(&harvestID)

	if err != nil {
		fmt.Println("Error when selecting the harvest id:")
		fmt.Println(err)
	}

	return harvestID
}

// NewProcessByID loads the harvest process from database by its own ID
func NewProcessByID(harvestProcessID int) *Process {
	var processID int
	sql := "SELECT `process` FROM `process_harvest` WHERE id=?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error 123 when selecting the ID of repository from harvest process:")
		fmt.Println(err)
	}
	query.QueryRow(harvestProcessID).Scan(&processID)

	return NewProcess(processID)
}

// NewProcess selects from database and creates a harvest.Process by ID
// NOTE! It returns only the Repository
func NewProcess(processID int) *Process {

	var (
		repID, harvestID int
		local            *repository.Repository
	)

	sql := "SELECT `id`, `repository` FROM `process_harvest` WHERE process=?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when selecting the ID of repository from harvest process:")
		fmt.Println(err)
	}
	query.QueryRow(processID).Scan(&harvestID, &repID)

	local, err2 := repository.NewRepository(strconv.Itoa(repID))

	if err2 != nil {
		fmt.Println(err2)
	}

	return &Process{
		HarvestID:  harvestID,
		repository: local,
		remote:     getRemoteServer(local),
		Process:    &*action.NewProcess(processID),
	}
}

// GetProcessesByRepository returns the processes of for the repository
func GetProcessesByRepository(repositoryID int) []*Process {

	var (
		list                 []*Process
		processID, harvestID int
		repID                string
	)

	repID = strconv.Itoa(repositoryID)

	sql := "SELECT `id`, `process` FROM `process_harvest` WHERE `repository` = ? ORDER BY process DESC"
	rows, err := database.Connection.Query(sql, repositoryID)

	if err != nil {
		fmt.Println("Error while selecting the processes by repository: ")
		fmt.Println(repositoryID)
	}

	for rows.Next() {
		rows.Scan(&harvestID, &processID)
		rep, _ := repository.NewRepository(repID)
		process := Process{
			HarvestID:  harvestID,
			repository: rep,
			Process:    action.NewProcess(processID),
		}
		list = append(list, &process)
	}
	return list
}

// DeleteProcess deletes a process
func DeleteProcess(processID int) {
	sql := "DELETE FROM `process_harvest` WHERE process=?"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		fmt.Println("Error when deleting the harvest process:")
		fmt.Println(err)
	}
	query.Exec(processID)
}
