package harvest

import (
	action "github.com/cristian-sima/Wisply/models/action"
	"github.com/cristian-sima/Wisply/models/harvest/remote"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Operationer ... defines the set of methods which should be implemented by the harvest operations
type Operationer interface {
	Start()
	GetOperation() *action.Operation
}

// Operation represents an operation
type Operation struct {
	Operationer
	*action.Operation
	process *Process // it is the harvest process
}

// Start notifies the controller that it is starting
func (operation *Operation) Start() {
	// TODO notify controller
}

// Finish calls its finish parents method
func (operation *Operation) Finish() {
	operation.Operation.Finish()

	msg := action.Message{
		Name:  "Finish",
		Value: operation.GetResult(),
	}

	operation.TellProcess(msg)
}

// ChangeRepositoryStatus tells the controller to change the status of the local repository
func (operation *Operation) ChangeRepositoryStatus(status string) {
	operation.process.ChangeRepositoryStatus(status)
}

// GetRepository returns a reference to the local repository
func (operation *Operation) GetRepository() *repository.Repository {
	return operation.process.GetRepository()
}

// GetRemote returns a reference to the remote repository
func (operation *Operation) GetRemote() remote.RepositoryInterface {
	return operation.process.GetRemoteServer()
}

// GetOperation returns the operation
func (operation *Operation) GetOperation() *action.Operation {
	return operation.Operation
}

func newOperation(process *action.Process, content string) *action.Operation {
	return &*process.CreateOperation(content)
}

// HarvestingOperation is the operation which harvest something
type HarvestingOperation struct {
	*Operation
}

func (operation *HarvestingOperation) failed() {
	operation.ChangeResult("danger")
	operation.Finish()
}

func (operation *HarvestingOperation) succeeded() {
	operation.Finish()
}

// constructor
func newHarvestingOperation(harvestProcess *Process, name string) *HarvestingOperation {
	return &HarvestingOperation{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "Harvest "+name),
		},
	}
}
