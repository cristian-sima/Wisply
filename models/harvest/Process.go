package harvest

import (
	"fmt"

	action "github.com/cristian-sima/Wisply/models/action"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Process is a link between controller and repository
type Process struct {
	*action.Process
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
	repository.NewRepository(ID)

	// switch local.Category {
	// case "EPrints":
	// 	{
	// 		remote = &EPrintsRepository{
	// 			URL: local.URL,
	// 		}
	// 	}
	// }

	process := &Process{
		Process:    &*action.CreateProcess(ID, "harvesting"),
		remote:     remote,
		Controller: controller,
		Actions:    make(map[string]*Action2),
	}
	// process.SetName("Harvest")
	// db.SetManager(process)
	// remote.SetManager(process)

	return process
}
