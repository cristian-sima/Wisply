package harvest

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

// IdentificationOperation encapsulates the methods for requesting information from the repository
type IdentificationOperation struct {
	*HarvestingOperation
}

// Start starts the action. Gets the page, get the content, clean database and store it
func (operation *IdentificationOperation) Start() {
	operation.Operation.Start()
	operation.ChangeRepositoryStatus("initializing")
	operation.tryToGet()
}

func (operation *IdentificationOperation) tryToGet() {

	rem := operation.process.GetRemoteServer()

	// create a task to request the server
	task := newGetTask(operation, rem)

	content, err := task.Identify()

	if err != nil {
		operation.failed()
	} else {
		operation.tryToParse(content)
	}
}

func (operation *IdentificationOperation) tryToParse(page []byte) {

	rem := operation.process.GetRemoteServer()

	task := newParseTask(operation, rem)

	response, err := task.GetIdentification(page)

	if err != nil {
		operation.failed()
	} else {
		operation.insertIdentification(response)
	}
}

func (operation *IdentificationOperation) insertIdentification(result *wisply.Identificationer) {
	task := newInsertIdentificationTask(operation, operation.GetRepository())
	err := task.Insert(result)
	if err != nil {
		operation.failed()
	} else {
		operation.succeeded()
	}
}

// constructor
func newIdentificationOperation(harvestProcess *Process) Operationer {
	return &IdentificationOperation{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Identifying"),
	}
}
