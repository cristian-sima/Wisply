package harvest

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

// HarvestingRecords is the operation which collects the records from a remote repository
type HarvestingRecords struct {
	*HarvestingOperation
}

// Start gets the collections
func (operation *HarvestingRecords) Start() {
	operation.Operation.Start()
	operation.start()
}

func (operation *HarvestingRecords) start() {
	err := operation.clear()
	if err != nil {
		operation.failed()
	} else {
		operation.multiRequest()
	}
}

func (operation *HarvestingRecords) clear() error {
	rem := operation.Operation.GetRepository()
	task := newInsertRecordsTask(operation, rem)
	return task.Clear()
}

func (operation *HarvestingRecords) multiRequest() {
	var (
		hasMoreRecords  bool
		resumptionToken string
		err             error
	)
	hasMoreRecords = true // in order to enter in the loop
	for hasMoreRecords && (err == nil) {
		err = operation.tryToGet(resumptionToken)
		if err == nil {
			resumptionToken, hasMoreRecords = operation.GetRemote().GetNextPage()
		}
	}
	if err != nil {
		operation.failed()
	} else {
		operation.succeeded()
	}
}

func (operation *HarvestingRecords) tryToGet(token string) error {

	rem := operation.GetRemote()

	task := newGetTask(operation, rem)

	content, err := task.GetRecords(token)

	if err != nil {
		return err
	}
	return operation.tryToParse(content)

}

func (operation *HarvestingRecords) tryToParse(page []byte) error {
	rem := operation.GetRemote()
	task := newParseTask(operation, rem)
	records, err := task.GetRecords(page)
	if err != nil {
		return err
	}
	return operation.insertRecords(records)

}

func (operation *HarvestingRecords) insertRecords(records []wisply.Recorder) error {
	repository := operation.Operation.GetRepository()
	task := newInsertRecordsTask(operation, repository)
	err := task.Insert(records)
	if err != nil {
		return err
	}
	return nil
}

// constructor
func newHarvestingRecords(harvestProcess *Process) Operationer {
	return &HarvestingRecords{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Records"),
	}
}
