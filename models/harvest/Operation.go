package harvest

import action "github.com/cristian-sima/Wisply/models/action"

// Operationer ... defines the set of methods which should be implemented by the harvest operations
type Operationer interface {
	Start()
	GetOperation() *action.Operation
}

// Operation represents an operation
type Operation struct {
	*action.Operation
}

// VerificationOperation encapsulates the methods for validating the repository
type VerificationOperation struct {
	Operationer
	*Operation
	process *Process
}

// Start starts the operation
func (operation *VerificationOperation) Start() {
	operation.process.ChangeLocalStatus("verifying")
}

// GetOperation returns the operation
func (operation *VerificationOperation) GetOperation() *action.Operation {
	return operation.Operation.Operation
}

func newVerificationOperation(process *Process) Operationer {
	return &VerificationOperation{
		process: process,
		Operation: &Operation{
			newOperation(process.Process, "verification"),
		},
	}
}

func newOperation(process *action.Process, content string) *action.Operation {
	return &*process.CreateOperation("Verification")
}
