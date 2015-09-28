package harvest

import (
	"fmt"

	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Manager is a link between controller and repository
type Manager struct {
	remote        RemoteRepositoryInterface
	local         *repository.Repository
	db            *databaseManager
	CurrentAction int                `json:"CurrentAction"`
	Actions       map[string]*Action `json:"Actions"`
	Controller    Controller
}

// StartProcess starts the process
func (manager *Manager) StartProcess() {
	manager.log("I start the process for repository at " + manager.local.URL + "... ")
	manager.changeLocalStatus("verifying")
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
				manager.log("I harvest the identification")
				manager.changeLocalStatus("verified")
				manager.changeLocalStatus("initializing")
				manager.remote.HarvestIdentification()
			}
			break
		}
	}
}

// SaveIdentification receives the identification result and saves it in the local repository
func (manager *Manager) SaveIdentification(result IdentificationResulter) {
	manager.log("I received the identification.")
	if !result.IsOk() {
		manager.changeLocalStatus("problems")
		manager.log("Problems with identification")
	} else {
		manager.log("The identification is ok")
		manager.changeLocalStatus("ok")
		manager.db.InsertIdentity(result.GetData())
		manager.End()
	}
}

// End receives the identification result and saves it in the local repository
func (manager *Manager) End() {
	manager.changeLocalStatus("unverified")
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
	}
	db.SetManager(manager)
	remote.SetManager(manager)
	return manager
}
