package harvest

import action "github.com/cristian-sima/Wisply/models/action"

// VerificationOperation encapsulates the methods for validating the repository
type VerificationOperation struct {
	Operationer
	*Operation
}

// Start starts the operation
func (operation *VerificationOperation) Start() {
	operation.Operation.Start()
	operation.ChangeRepositoryStatus("verifying")
}

// Activity creates a task to verify if the url is good
func (operation *VerificationOperation) Activity() {
	// create a task to request the server

	// create a task to check the result
}

func (operation *VerificationOperation) verificationFailed() {
	operation.ChangeRepositoryStatus("verification-failed")
	operation.ChangeResult("danger")
	operation.Finish()
}

func (operation *VerificationOperation) verificationSucceded() {
	operation.ChangeRepositoryStatus("verified")
	operation.Finish()
}

//  process.record("The validation passed")
//  process.ChangeLocalStatus("verified")
//  process.startIdentification()

// GetOperation returns the operation
func (operation *VerificationOperation) GetOperation() *action.Operation {
	return operation.Operation.Operation
}

func newVerificationOperation(harvestProcess *Process) Operationer {
	return &VerificationOperation{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "verification"),
		},
	}
}
