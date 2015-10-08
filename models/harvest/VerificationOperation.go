package harvest

import action "github.com/cristian-sima/Wisply/models/action"

// VerificationOperation encapsulates the methods for validating the repository
type VerificationOperation struct {
	Operationer
	*Operation
	process *Process
}

// Start starts the operation
func (operation *VerificationOperation) Start() {
	operation.process.ChangeLocalStatus("verifying")
	operation.verificationFailed()
}

func (operation *VerificationOperation) verificationFailed() {
	operation.ChangeResult("danger")
	operation.Finish()
}

//  process.record("The validation passed")
//  process.ChangeLocalStatus("verified")
//  process.startIdentification()

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
