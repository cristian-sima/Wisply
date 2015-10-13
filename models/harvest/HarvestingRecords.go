package harvest

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/harvest/wisply"
)

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
		hasMoreRecords bool
		err            error
	)
	token := operation.getCurrentToken()
	hasMoreRecords = true // in order to enter in the loop
	for hasMoreRecords && (err == nil) {
		err = operation.tryToGet(token)
		if err == nil {
			token, hasMoreRecords = operation.GetRemote().GetNextPage()
		}
		if hasMoreRecords {
			operation.process.SaveToken("records", token)
		} else if operation.process.Records != 0 {
			lastToken := operation.GetRemote().GetFinishToken()
			fmt.Println("finishing token: " + lastToken)
			operation.process.SaveToken("records", lastToken)
		} else if operation.process.Records == 0 {
			operation.process.SaveToken("records", token)
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
	if len(records) != 0 {
		return operation.insertRecords(records)
	}
	return nil
}

func (operation *HarvestingRecords) insertRecords(records []wisply.Recorder) error {
	repository := operation.Operation.GetRepository()
	task := newInsertRecordsTask(operation, repository)
	err := task.Insert(records)
	if err != nil {
		return err
	}
	err = operation.process.updateRecords(len(records))
	return err

}

// constructor
func newHarvestingRecords(harvestProcess *Process) Operationer {
	return &HarvestingRecords{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Records"),
	}
}
