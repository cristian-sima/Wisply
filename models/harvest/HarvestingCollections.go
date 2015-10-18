package harvest

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/wisply"
)

// HarvestingCollections is the operation which collects the collections from a remote repository
type HarvestingCollections struct {
	*HarvestingOperation
}

// Start gets the collections
func (operation *HarvestingCollections) Start() {
	operation.Operation.Start()
	operation.multiRequest()
}

func (operation *HarvestingCollections) multiRequest() {
	var (
		hasMoreCollections bool
		err                error
	)
	token := operation.getCurrentToken()
	operation.process.SaveToken("collections", token)
	hasMoreCollections = true // in order to enter in the loop
	initNumberOfCollections := operation.process.Collections
	for hasMoreCollections && (err == nil) {
		err = operation.tryToGet(token)
		if err == nil {
			token, hasMoreCollections = operation.GetRemote().GetNextPage()
		}
		if hasMoreCollections {
			operation.process.SaveToken("collections", token)
		} else if operation.process.Collections != initNumberOfCollections {
			lastToken := operation.GetRemote().GetFinishToken()
			fmt.Println("finishing token: " + lastToken)
			operation.process.SaveToken("collections", lastToken)
		}
	}
	if err != nil {
		operation.failed()
	} else {
		operation.succedded()
	}
}

func (operation *HarvestingCollections) tryToGet(token string) error {
	rem := operation.process.GetRemoteServer()
	task := newGetTask(operation, rem)
	content, err := task.GetCollections(token)
	if err != nil {
		return err
	}
	return operation.tryToParse(content)

}

func (operation *HarvestingCollections) tryToParse(page []byte) error {
	rem := operation.process.GetRemoteServer()
	task := newParseTask(operation, rem)
	collections, err := task.GetCollections(page)
	if err != nil {
		return err
	}
	return operation.insertCollections(collections)

}

func (operation *HarvestingCollections) insertCollections(collections []wisply.Collectioner) error {
	repository := operation.Operation.GetRepository()
	task := newInsertCollectionsTask(operation, repository)
	err := task.Insert(collections)
	if err != nil {
		return err
	}
	oldCollections := operation.process.Collections
	newCollections := len(collections)
	difference := newCollections - oldCollections
	if difference != 0 {
		err = operation.process.updateCollections(difference)
	}
	return err
}

// constructor
func newHarvestingCollections(harvestProcess *Process) Operationer {
	return &HarvestingCollections{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Collections"),
	}
}
