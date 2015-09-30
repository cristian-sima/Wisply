package harvest

import (
	"fmt"

	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Manager is a link between controller and repository
type Manager struct {
	remote         RemoteRepositoryInterface
	local          *repository.Repository
	db             *databaseManager
	CurrentAction  int                `json:"CurrentAction"`
	Actions        map[string]*Action `json:"Actions"`
	Controller     Controller         `json:"-"`
	Identification *Identificationer  `json:"Identification"`
}

// StartProcess starts the process
func (manager *Manager) StartProcess() {
	manager.log("I start the process for repository at " + manager.local.URL + "... ")
	manager.changeLocalStatus("verifying")
	manager.setCurrentAction("verifying")
	manager.remote.Validate()
}

// Notify is called by a harvest repository with a message
func (manager *Manager) Notify(message *Message) {
	manager.log("The manager has received this message:")
	fmt.Println(message)
	switch message.Name {
	case "verification-finished":
		{
			if message.Value == "failed" {
				manager.changeLocalStatus("verification-failed")
				manager.notifyController(message)
			} else {
				manager.log("The validation passed")
				manager.changeLocalStatus("verified")
				manager.harvestIdentification()
			}
			break
		}
	}
}

func (manager *Manager) harvestIdentification() {
	manager.log("I harvest the identification")
	manager.changeLocalStatus("initializing")
	manager.setCurrentAction("initializing")
	manager.remote.HarvestIdentification()
}

// SaveIdentification receives the identification result and saves it in the local repository
func (manager *Manager) SaveIdentification(result IdentificationResulter) {
	manager.log("I received the identification.")
	if !result.IsOk() {
		manager.changeLocalStatus("verification-failed")
		manager.log("Problems with identification")
	} else {
		manager.notifyController(&Message{
			Name:  "identification-details",
			Value: result.GetData(),
		})
		manager.log("The identification is ok")
		manager.changeLocalStatus("ok")
		manager.db.InsertIdentity(result.GetData())
		manager.Identification = result.GetData()
		manager.harvestFormarts()
	}
}

func (manager *Manager) harvestFormarts() {
	manager.changeLocalStatus("updating")
	manager.setCurrentAction("harvesting")
	manager.createAction("formats")
	manager.remote.HarvestFormats()

}

// SaveFormats retrives a format and saves it
func (manager *Manager) SaveFormats(result FormatResulter) {
	formats := result.GetData()
	manager.log("Format received")
	manager.updateAction(len(formats), "formats")
	manager.db.InsertFormats(formats)
}

// EndFormats it notifies the client the formats are finished
func (manager *Manager) EndFormats() {
	manager.endAction("formats")
	manager.harvestCollections()
}

// COLLECTIONS

func (manager *Manager) harvestCollections() {
	manager.log("I am harvasting collections")
	manager.setCurrentAction("harvesting")
	manager.createAction("collections")
	manager.db.ClearCollections()
	manager.remote.HarvestCollections()

}

// SaveCollections retrives the collections and stores them
func (manager *Manager) SaveCollections(result CollectionResult) {
	collections := result.GetData()
	manager.log("Collections received")
	manager.updateAction(len(collections), "collections")
	fmt.Println("insert")
	manager.db.InsertCollections(collections)
}

// EndCollections it notifies the client the collections are finished
func (manager *Manager) EndCollections() {
	manager.endAction("collections")
	manager.harvestRecords()
}

// Records

func (manager *Manager) harvestRecords() {
	manager.log("I am harvasting records")
	manager.setCurrentAction("harvesting")
	manager.createAction("records")
	manager.db.ClearRecords()
	manager.remote.HarvestRecords()

}

// SaveRecords retrives the records and stores them
func (manager *Manager) SaveRecords(result RecordResult) {
	records := result.GetData()
	manager.log("Records received")
	manager.updateAction(len(records), "records")
	manager.db.InsertRecords(records)
}

// EndRecords it notifies the client the records are finished
func (manager *Manager) EndRecords() {
	manager.endAction("records")
	manager.End()
}

// ---

func (manager *Manager) endAction(name string) {
	manager.Actions[name].Finish()
	manager.notifyAction(manager.Actions[name], "finish")
}

func (manager *Manager) setCurrentAction(actionName string) {
	manager.CurrentAction = Actions[actionName]
}

func (manager *Manager) createAction(name string) {
	manager.Actions[name] = &Action{
		Type:      name,
		IsCurrent: true,
	}
	manager.notifyAction(manager.Actions[name], "start")
}

func (manager *Manager) updateAction(newCount int, name string) {
	action := manager.Actions[name]
	action.Update(newCount)
	manager.notifyAction(action, "update")
}

func (manager *Manager) notifyAction(action *Action, operation string) {

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

	manager.notifyController(&Message{
		Name:  "harvesting",
		Value: content,
	})
}

// End receives the identification result and saves it in the local repository
func (manager *Manager) End() {
	manager.changeLocalStatus("unverified")
	manager.notifyController(&Message{
		Name: "delete-process",
	})
}

func (manager *Manager) changeLocalStatus(newStatus string) {
	manager.local.ModifyStatus(newStatus)
	manager.notifyController(&Message{
		Name:  "status-changed",
		Value: newStatus,
	})
}

func (manager *Manager) log(message string) {
	fmt.Println("<-->  Harvest Manager: " + message)
}

// GetRepository returns the wisply repository
func (manager *Manager) GetRepository() *repository.Repository {
	return manager.local
}

func (manager *Manager) notifyController(message *Message) {
	message.Repository = manager.local.ID
	manager.Controller.Notify(message)
}

// NewManager creates a new mananger
func NewManager(ID string, controller Controller) *Manager {
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

	manager := &Manager{
		local:      local,
		remote:     remote,
		Controller: controller,
		db:         db,
		Actions:    make(map[string]*Action),
	}
	db.SetManager(manager)
	remote.SetManager(manager)
	return manager
}
