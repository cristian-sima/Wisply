package harvest

import action "github.com/cristian-sima/Wisply/models/action"

// VerificationOperation encapsulates the methods for validating the repository
type VerificationOperation struct {
	Operationer
	*Operation
}

// Start starts the action.Operation, change the status of repository to verifying and gets the page
func (operation *VerificationOperation) Start() {
	operation.Operation.Start()
	operation.ChangeRepositoryStatus("verifying")
	operation.tryToGet()
}

// Activity creates a task to verify if the URL is good
func (operation *VerificationOperation) tryToGet() {

	// remote := operation.Process.GetRemote()
	repository := operation.process.GetRepository()

	// create a task to request the server
	task := newGetRequestTask(operation)

	page, err := task.Identify(repository.URL)

	if err != nil {
		operation.verificationFailed()
	} else {
		operation.tryToParse(page)
	}
	// create a task to check the result
}

func (operation *VerificationOperation) tryToParse(page []byte) {
	task := newParseRequestTask(operation)
	_, err := task.Parse(page)
	if err != nil {
		operation.verificationFailed()
	} else {
		operation.verificationSucceded()
	}
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

// GetOperation returns the operation
func (operation *VerificationOperation) GetOperation() *action.Operation {
	return operation.Operation.Operation
}

func newVerificationOperation(harvestProcess *Process) Operationer {
	return &VerificationOperation{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "Verification"),
		},
	}
}
