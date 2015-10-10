package harvest

import (
	"fmt"
	"strconv"

	action "github.com/cristian-sima/Wisply/models/action"
	database "github.com/cristian-sima/Wisply/models/database"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Process is a link between controller and repository
type Process struct {
	*action.Process
	repository     *repository.Repository
	remote         RemoteRepositoryInterface
	CurrentAction  int                 `json:"CurrentAction"`
	Actions        map[string]*Action2 `json:"Actions"`
	Controller     WisplyController    `json:"-"`
	Identification *Identificationer   `json:"Identification"`
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
			switch message.GetOperation().Content {
			case "Verification":
				if message.GetValue() == "normal" {
					go process.identify()
				} else {
					process.ChangeResult("danger")
					process.Finish()
				}
				break
			case "Identifying":
				if message.GetValue() == "normal" {
					go process.harvest()
				} else {
					process.ChangeResult("danger")
					process.Finish()
				}
				break
			case "Harvest Formats":
				if message.GetValue() == "normal" {
					go process.harvestCollections()
				} else {
					process.ChangeResult("danger")
					process.Finish()
				}
				break
			case "Harvest Collections":
				if message.GetValue() == "normal" {
					go process.harvestRecords()
				} else {
					process.ChangeResult("danger")
					process.Finish()
				}
				break
			case "Harvest Records":
				if message.GetValue() == "normal" {
					go process.harvestIdentifiers()
				} else {
					process.ChangeResult("danger")
					process.Finish()
				}
			case "Harvest Identifiers":
				if message.GetValue() != "normal" {
					process.ChangeResult("danger")
				}
				process.ChangeRepositoryStatus("ok")
				process.Finish()
				break
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

// GetRepository returns the wisply repository
func (process *Process) GetRepository() *repository.Repository {
	return process.repository
}

// GetRemote returns the interface of a remote repository
func (process *Process) GetRemote() RemoteRepositoryInterface {
	return process.remote
}

func (process *Process) harvestFormarts() {
	process.ChangeRepositoryStatus("updating")
	process.setCurrentAction("harvesting")
	process.createAction("formats")
	process.remote.HarvestFormats()

}

// ---

func (process *Process) endAction(name string) {
	process.Actions[name].Finish()
	process.notifyAction(process.Actions[name], "finish")
}

func (process *Process) setCurrentAction(actionName string) {
	process.CurrentAction = Actions[actionName]
}

func (process *Process) createAction(name string) {
	process.Actions[name] = &Action2{
		Type:      name,
		IsCurrent: true,
	}
	process.notifyAction(process.Actions[name], "start")
}

func (process *Process) updateAction(newCount int, name string) {
	action := process.Actions[name]
	action.Update(newCount)
	process.notifyAction(action, "update")
}

// ChangeCurrentOperation informs the controller about the change and it calls its father
func (process *Process) ChangeCurrentOperation(operation Operationer) {
	process.Process.ChangeCurrentOperation(operation.GetOperation())
}

func (process *Process) notifyAction(action *Action2, operation string) {

	type Content struct {
		Operation string `json:"Operation"`
		Type      string `json:"Type"`
		Count     int    `json:"Count"`
	}

	content := Content{
		Operation: operation,
		Type:      action.Type,
		Count:     action.Count,
	}

	process.notifyController(&Message{
		Name:  "harvesting",
		Value: content,
	})
}

// End receives the identification result and saves it in the local repository
func (process *Process) End() {
	process.record("The process is stopped")
	process.notifyController(&Message{
		Name: "delete-process",
	})
}

// ChangeRepositoryStatus changes the status of local repository
func (process *Process) ChangeRepositoryStatus(newStatus string) {
	process.repository.ModifyStatus(newStatus)
	process.notifyController(&Message{
		Name:  "status-changed",
		Value: newStatus,
	})
}

func (process *Process) notifyController(message *Message) {
	message.Repository = process.repository.ID
	process.Controller.Notify(message)
}

func (process *Process) record(message string) {
	process.notifyController(&Message{
		Value: message,
		Name:  "event-notice",
	})
	//process.operation.record(message, process.Local.ID)
}

// Delete deletes the harvest process and calls its parent method
func (process *Process) Delete() {
	DeleteProcess(process.ID)
}

// CreateProcess creates a new harvest process
func CreateProcess(ID string, controller WisplyController) *Process {
	// var remote RemoteRepositoryInterface

	local, _ := repository.NewRepository(ID)

	// switch local.Category {
	// case "EPrints":
	// 	{
	// 		remote = &EPrintsRepository{
	// 			URL: local.URL,
	// 		}
	// 	}
	// }

	process := &Process{
		Process: &*action.CreateProcess("Harvest"),
		// remote:     remote,
		Controller: controller,
		repository: local,
		Actions:    make(map[string]*Action2),
	}

	insertHarvestProcess(process)

	return process
}

func insertHarvestProcess(process *Process) {
	columns := "(`process`, `repository`)"
	values := "(?, ?)"
	sql := "INSERT INTO `process_harvest` " + columns + " VALUES " + values

	query, err := database.Database.Prepare(sql)

	if err != nil {
		fmt.Println("Error when creating the harvest process:")
		fmt.Println(sql)
		fmt.Println(err)
	}
	query.Exec(process.ID, process.GetRepository().ID)
}

// NewProcess selects from database and creates a harvest.Process by ID
// NOTE! It returns only the Repository
func NewProcess(processID int) *Process {

	var (
		repID int
		local *repository.Repository
	)

	sql := "SELECT `repository` FROM `process_harvest` WHERE process=?"
	query, err := database.Database.Prepare(sql)

	if err != nil {
		fmt.Println("Error when selecting the ID of repository from harvest process:")
		fmt.Println(err)
	}
	query.QueryRow(processID).Scan(&repID)

	local, err2 := repository.NewRepository(strconv.Itoa(repID))

	if err2 != nil {
		fmt.Println(err2)
	}

	return &Process{
		repository: local,
	}
}

// GetProcessesByRepository returns the processes of for the repository
func GetProcessesByRepository(repositoryID int) []*Process {

	var (
		list      []*Process
		processID int
		repID     string
	)

	repID = strconv.Itoa(repositoryID)

	sql := "SELECT `process` FROM `process_harvest` WHERE `repository` = ? ORDER BY process DESC"
	rows, err := database.Database.Query(sql, repositoryID)

	if err != nil {
		fmt.Println("Error while selecting the processes by repository: ")
		fmt.Println(repositoryID)
	}

	for rows.Next() {
		rows.Scan(&processID)
		rep, _ := repository.NewRepository(repID)
		process := Process{
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
	query, err := database.Database.Prepare(sql)

	fmt.Println(processID)

	if err != nil {
		fmt.Println("Error when deleting the harvest process:")
		fmt.Println(err)
	}
	query.Exec(processID)
}
