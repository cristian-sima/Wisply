package harvest

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

// HarvestingRecords is the operation which collects the records from a remote repository
type HarvestingRecords struct {
	*HarvestingOperation
}

// Start gets the collections
func (operation *HarvestingRecords) Start() {
	operation.Operation.Start()
	operation.tryToGet()
}

func (operation *HarvestingRecords) tryToGet() {

	rem := operation.GetRemote()

	task := newGetTask(operation, rem)

	content, err := task.GetRecords()

	if err != nil {
		operation.failed()
	} else {
		operation.tryToParse(content)
	}
}

func (operation *HarvestingRecords) tryToParse(page []byte) {
	rem := operation.GetRemote()
	task := newParseTask(operation, rem)
	records, err := task.GetRecords(page)
	if err != nil {
		operation.failed()
	} else {
		operation.insertRecords(records)
	}
}

func (operation *HarvestingRecords) insertRecords(records []wisply.Recorder) {
	repository := operation.Operation.GetRepository()
	task := newInsertRecordsTask(operation, repository)
	err := task.Insert(records)
	if err != nil {
		operation.failed()
	} else {
		operation.succeeded()
	}
}

// constructor
func newHarvestingRecords(harvestProcess *Process) Operationer {
	return &HarvestingRecords{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Records"),
	}
}
