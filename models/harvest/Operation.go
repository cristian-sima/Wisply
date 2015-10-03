package harvest

import "fmt"

import (
	history "github.com/cristian-sima/Wisply/models/history"
	"github.com/cristian-sima/Wisply/models/repository"
)

// WisplyController represents a controller
type WisplyController interface {
	Notify(*Message)
}

// WisplyProcessInterface ... defines the methods which must be implemented by a process
type WisplyProcessInterface interface {
	Start()
	ManagerFinished()
}

// Operation represents the basic operation
type operation struct {
	history       *history.Manager
	operationName string
	operationType string
	local         *repository.Repository
}

func (operation *operation) record(message string, repository int) {
	event := history.Event{
		Content:       message,
		Repository:    repository,
		OperationType: operation.operationType,
		OperationName: operation.operationName,
	}
	operation.history.Record(&event)
}

// WisplyProcess is a basic process. A process does a series of actions using managers
type WisplyProcess struct {
	operation
	controller *WisplyController
}

func (process *WisplyProcess) log(message string) {
	fmt.Println(process.operation.operationType + " " + process.operation.operationName + ": " + message)
}

// GetController returns the reference to the controller which manages the process
func (process *WisplyProcess) GetController() *WisplyController {
	return process.controller
}

// SetName sets the name of a process
func (process *WisplyProcess) SetName(name string) {
	process.operationType = "Process"
	process.operationName = name
}

// ManagerInterface ... defines the set of methods which must be implemented by a harvest manager
type ManagerInterface interface {
	Start()
	End()
	Save()
}

// WisplyManager represents a general manager
type WisplyManager struct {
	process *WisplyProcess
}

// GetProcess returns the proces of the manager
func (manager *WisplyManager) GetProcess() *WisplyProcess {
	return manager.process
}
