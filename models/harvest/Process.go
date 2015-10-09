package harvest

import (
	"fmt"
	"strconv"

	action "github.com/cristian-sima/Wisply/models/action"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Process is a link between controller and repository
type Process struct {
	*action.Process
	remote         RemoteRepositoryInterface
	db             *databaseManager
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
	fmt.Println("RUN.....")
	for {
		select {
		case message := <-process.Process.GetOperationConduit():
			fmt.Println("Got a message from the operation " + message.GetOperation().Content + ": ")
			switch message.GetOperation().Content {
			case "Verification":
				if message.GetValue() == "normal" {
					go process.identify()
				} else {
					process.ChangeResult("danger")
					process.Finish()
				}
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

// --- end activity

// GetRepository returns the wisply repository
func (process *Process) GetRepository() *repository.Repository {
	return process.Repository
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

// SaveFormats retrives a format and saves it
func (process *Process) SaveFormats(result FormatResulter) {
	formats := result.GetData()
	process.record("Format received")
	process.updateAction(len(formats), "formats")
	process.db.InsertFormats(formats)
}

// EndFormats it notifies the client the formats are finished
func (process *Process) EndFormats() {
	process.endAction("formats")
	process.startHarvestingCollections()
}

// COLLECTIONS

func (process *Process) startHarvestingCollections() {
	process.record("I am harvasting collections")
	process.setCurrentAction("harvesting")
	process.createAction("collections")
	process.db.ClearCollections()
	process.remote.HarvestCollections()

}

// SaveCollections retrives the collections and stores them
func (process *Process) SaveCollections(result CollectionResult) {
	collections := result.GetData()
	numberOfCollections := len(collections)
	process.record(strconv.Itoa(numberOfCollections) + " collections received")
	process.updateAction(numberOfCollections, "collections")
	process.db.InsertCollections(collections)
}

// EndCollections it notifies the client the collections are finished
func (process *Process) EndCollections() {
	process.record("The collections harvesting is finished")
	process.endAction("collections")
	process.harvestRecords()
}

// Records

func (process *Process) harvestRecords() {
	process.record("I am harvasting records")
	process.setCurrentAction("harvesting")
	process.createAction("records")
	process.db.ClearRecords()
	process.remote.HarvestRecords()

}

// SaveRecords retrives the records and stores them
func (process *Process) SaveRecords(result RecordResult) {
	records := result.GetData()
	numberOfRecords := len(records)
	process.record(strconv.Itoa(numberOfRecords) + " records received. From [" + records[0].GetIdentifier() + "] until [" + records[numberOfRecords-1].GetIdentifier() + "].")
	process.updateAction(numberOfRecords, "records")
	process.db.InsertRecords(records)
}

// EndRecords it notifies the client the records are finished
func (process *Process) EndRecords() {
	process.record("The record harvesting is finished")
	process.endAction("records")
	process.End()
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
	// TODO notify the controller
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
	process.Repository.ModifyStatus(newStatus)
	process.notifyController(&Message{
		Name:  "status-changed",
		Value: newStatus,
	})
}

func (process *Process) notifyController(message *Message) {
	message.Repository = process.Repository.ID
	process.Controller.Notify(message)
}

func (process *Process) record(message string) {
	process.notifyController(&Message{
		Value: message,
		Name:  "event-notice",
	})
	//process.operation.record(message, process.Local.ID)
}

// CreateProcess creates a new harvest process
func CreateProcess(ID string, controller WisplyController) *Process {
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
		Process:    &*action.CreateProcess(ID, "harvesting"),
		remote:     remote,
		Controller: controller,
		db:         db,
		Actions:    make(map[string]*Action2),
	}
	// process.SetName("Harvest")
	db.SetManager(process)
	remote.SetManager(process)

	return process
}
