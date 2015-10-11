package harvest

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

// HarvestingIdentifiers is the operation which collects the identifiers from a remote repository
type HarvestingIdentifiers struct {
	*HarvestingOperation
}

// Start gets the identifiers
func (operation *HarvestingIdentifiers) Start() {
	operation.Operation.Start()
	operation.tryToGet()
}

func (operation *HarvestingIdentifiers) tryToGet() {

	rem := operation.GetRemote()

	task := newGetTask(operation, rem)

	content, err := task.GetIdentifiers()

	if err != nil {
		operation.failed()
	} else {
		operation.tryToParse(content)
	}
}

func (operation *HarvestingIdentifiers) tryToParse(page []byte) {
	rem := operation.GetRemote()
	task := newParseTask(operation, rem)
	identifiers, err := task.GetIdentifiers(page)
	if err != nil {
		operation.failed()
	} else {
		operation.insertIdentifiers(identifiers)
	}
}

func (operation *HarvestingIdentifiers) insertIdentifiers(identifiers []wisply.Identifier) {
	repository := operation.Operation.GetRepository()
	task := newInsertIdentifiersTask(operation, repository)
	err := task.Insert(identifiers)
	if err != nil {
		operation.failed()
	} else {
		operation.succeeded()
	}
}

// constructor
func newHarvestingIdentifiers(harvestProcess *Process) Operationer {
	return &HarvestingIdentifiers{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Identifiers"),
	}
}
