package harvest

import action "github.com/cristian-sima/Wisply/models/action"

// VerificationOperation encapsulates the methods for validating the repository
type VerificationOperation struct {
	*HarvestingOperation
}

// Start starts the action.Operation, change the status of repository to
// verifying and gets the page
func (operation *VerificationOperation) Start() {
	operation.Operation.Start()
	operation.ChangeRepositoryStatus("verifying")
	operation.tryToGet()
}

// Activity creates a task to verify if the URL is good
func (operation *VerificationOperation) tryToGet() {

	repository := operation.process.GetRemoteServer()

	task := newGetTask(operation, repository)

	page, err := task.Verify()

	if err != nil {
		operation.failed()
	} else {
		operation.tryToParse(page)
	}
	// create a task to check the result
}

func (operation *VerificationOperation) tryToParse(page []byte) {

	repository := operation.process.GetRemoteServer()

	task := newParseTask(operation, repository)
	err := task.Verify(page)
	if err != nil {
		operation.failed()
	} else {
		operation.succeeded()
	}
}

func (operation *VerificationOperation) failed() {
	operation.ChangeRepositoryStatus("verification-failed")
	operation.HarvestingOperation.failed()
}

func (operation *VerificationOperation) succeeded() {
	operation.HarvestingOperation.succeeded()
}

// GetOperation returns the operation
func (operation *VerificationOperation) GetOperation() *action.Operation {
	return operation.Operation.Operation
}

// constructor
func newVerificationOperation(harvestProcess *Process) Operationer {
	return &VerificationOperation{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Verification"),
	}
}
