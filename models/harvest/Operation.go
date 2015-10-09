package harvest

import action "github.com/cristian-sima/Wisply/models/action"

// Operationer ... defines the set of methods which should be implemented by the harvest operations
type Operationer interface {
	Start()
	Activity()
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
	// TODO notify the controller it is finishing
}

// ChangeRepositoryStatus tells the controller to change the status of the local repository
func (operation *Operation) ChangeRepositoryStatus(status string) {
	operation.process.ChangeRepositoryStatus(status)
}

func newOperation(process *action.Process, content string) *action.Operation {
	return &*process.CreateOperation(content)
}
