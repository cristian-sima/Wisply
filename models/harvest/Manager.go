package harvest

import (
	"fmt"

	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Manager is a link between controller and repository
type Manager struct {
	Remote        *RemoteRepositoryInterface
	Local         *repository.Repository
	CurrentAction int                `json:"CurrentAction"`
	Actions       map[string]*Action `json:"Actions"`
	Controller    Controller
}

// StartProcess starts the process
func (manager *Manager) StartProcess() {
	fmt.Println("Started")
}

// Notify is called by a harvest repository with a message
func (manager *Manager) Notify(message *Message) {

	fmt.Println("--> Harvest Manager: The manager has received this message:")
	fmt.Println(message)

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

	manager := &Manager{
		Local:      local,
		Controller: controller,
		Remote:     &remote,
	}

	remote.SetManager(manager)

	return manager
}
