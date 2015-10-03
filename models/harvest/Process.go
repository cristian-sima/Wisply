package harvest

import (
	"strconv"

	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Process is a link between controller and repository
type Process struct {
	WisplyProcess
	remote         RemoteRepositoryInterface
	local          *repository.Repository
	db             *databaseManager
	CurrentAction  int                `json:"CurrentAction"`
	Actions        map[string]*Action `json:"Actions"`
	Controller     WisplyController   `json:"-"`
	Identification *Identificationer  `json:"Identification"`
	managers       []*WisplyManager
	currentManager *WisplyManager
}

// Start starts the process
func (process *Process) Start() {
	process.record("The process has started")
	process.startVerification()
}

func (process *Process) startVerification() {
	process.record("I start the harvesting process")
	process.changeLocalStatus("verifying")
	process.setCurrentAction("verifying")
	process.remote.Validate()
}

// Notify is called by a harvest repository with a message
func (process *Process) Notify(message *Message) {
	switch message.Name {
	case "verification-finished":
		{
			if message.Value == "failed" {
				process.record("The verification failed")
				process.changeLocalStatus("verification-failed")
				process.notifyController(message)
				process.End()
			} else {
				process.record("The validation passed")
				process.changeLocalStatus("verified")
				process.startIdentification()
			}
			break
		}
	}
}

func (process *Process) startIdentification() {
	process.record("I harvest the identification")
	process.changeLocalStatus("initializing")
	process.setCurrentAction("initializing")
	process.remote.HarvestIdentification()
}

// SaveIdentification receives the identification result and saves it in the local repository
func (process *Process) SaveIdentification(result IdentificationResulter) {
	process.record("I received the identification.")
	if !result.IsOk() {
		process.changeLocalStatus("verification-failed")
		process.record("Problems with identification")
	} else {
		process.notifyController(&Message{
			Name:  "identification-details",
			Value: result.GetData(),
		})
		process.record("The identification is ok")
		process.changeLocalStatus("ok")
		process.db.InsertIdentity(result.GetData())
		process.Identification = result.GetData()
		process.harvestFormarts()
	}
}

func (process *Process) harvestFormarts() {
	process.changeLocalStatus("updating")
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
	process.Actions[name] = &Action{
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

func (process *Process) notifyAction(action *Action, operation string) {

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

func (process *Process) changeLocalStatus(newStatus string) {
	process.local.ModifyStatus(newStatus)
	process.notifyController(&Message{
		Name:  "status-changed",
		Value: newStatus,
	})
}

// GetRepository returns the wisply repository
func (process *Process) GetRepository() *repository.Repository {
	return process.local
}

func (process *Process) notifyController(message *Message) {
	message.Repository = process.local.ID
	process.Controller.Notify(message)
}

func (process *Process) record(message string) {
	process.notifyController(&Message{
		Value: message,
		Name:  "event-notice",
	})
	process.operation.record(message, process.local.ID)
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
		local:      local,
		remote:     remote,
		Controller: controller,
		db:         db,
		Actions:    make(map[string]*Action),
	}
	process.SetName("Harvest")
	db.SetManager(process)
	remote.SetManager(process)
	return process
}
