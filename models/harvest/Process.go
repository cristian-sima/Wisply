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
	harvestID  int
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

			fmt.Println("Process: I got: ")
			fmt.Println(message)

			switch message.GetOperation().Content {
			case "Verification":
				if message.GetValue() == "normal" {
					go process.identify()
				} else {
					process.ChangeResult("danger")
					process.Finish()
					process.ChangeRepositoryStatus("verification-failed")
				}
				break
			case "Identifying":
				if message.GetValue() == "normal" {
					go process.harvest()
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
					process.ChangeResult("danger")
				}
				process.ChangeRepositoryStatus("ok")
				process.Finish()
				break
			default:
				fmt.Println("No such operation: " + message.GetOperation().Content)
			}
		}
	}
}

// Stage 1

func (process *Process) verify() {
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
	process.harvestFormats()
}

// FORMATS

func (process *Process) harvestFormats() {
	harvestingFormats := newHarvestingFormats(process)
	process.ChangeCurrentOperation(harvestingFormats)
	harvestingFormats.Start()
}

// COLLECTIONS

func (process *Process) harvestCollections() {
	harvestingCollections := newHarvestingCollections(process)
	process.ChangeCurrentOperation(harvestingCollections)
	harvestingCollections.Start()
}

// RECORDS

func (process *Process) harvestRecords() {
	harvestingRecords := newHarvestingRecords(process)
	process.ChangeCurrentOperation(harvestingRecords)
	harvestingRecords.Start()
}

// IDENTIFIERS

func (process *Process) harvestIdentifiers() {
	identifiers := newHarvestingIdentifiers(process)
	process.ChangeCurrentOperation(identifiers)
	identifiers.Start()
}

// --- end activity

func (process *Process) processFails() {
	process.ChangeResult("danger")
	process.Finish()
	process.ChangeRepositoryStatus("problems")
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

// SaveToken saves the token for a particular name
func (process *Process) SaveToken(name string, token string) {
	stmt, err := database.Connection.Prepare("UPDATE `process_harvest` SET `token_" + name + "`=? WHERE id=?")
	if err != nil {
		fmt.Println("Error 1 when updating the token for " + name + " inside harvesting process: ")
		fmt.Println(err)
	}
	_, err = stmt.Exec(token, process.harvestID)
	if err != nil {
		fmt.Println("Error 2 when updating the token for " + name + " inside harvesting process: ")
		fmt.Println(err)
	}
}

func (process *Process) tellController(simple *Message) {
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
	channel := process.controller.GetConduit()

	channel <- msg
}

// Delete deletes the harvest process and calls its parent method
func (process *Process) Delete() {
	DeleteProcess(process.ID)
}

// CreateProcess creates a new harvest process
func CreateProcess(ID string, controller Controller) *Process {

	var (
		rem       remote.RepositoryInterface
		harvestID int
	)

	local, _ := repository.NewRepository(ID)

	switch local.Category {
	case "EPrints":
		{
			rem = remote.NewEPrints(local)
		}
	}

	process := &Process{
		Process:    &*action.CreateProcess("Harvest"),
		remote:     rem,
		controller: controller,
		repository: local,
	}

	harvestID = insertHarvestProcess(process)

	process.harvestID = harvestID

	return process
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
		harvestID:  harvestID,
		repository: local,
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
			harvestID:  harvestID,
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
