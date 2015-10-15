package harvest

import "github.com/cristian-sima/Wisply/models/wisply"

// HarvestingCollections is the operation which collects the collections from a remote repository
type HarvestingCollections struct {
	*HarvestingOperation
}

// Start gets the collections
func (operation *HarvestingCollections) Start() {
	operation.Operation.Start()
	operation.tryToGet()
}

func (operation *HarvestingCollections) tryToGet() {
	rem := operation.process.GetRemoteServer()
	task := newGetTask(operation, rem)
	content, err := task.GetCollections()
	if err != nil {
		operation.failed()
	} else {
		operation.tryToParse(content)
	}
}

func (operation *HarvestingCollections) tryToParse(page []byte) {
	rem := operation.process.GetRemoteServer()
	task := newParseTask(operation, rem)
	collections, err := task.GetCollections(page)
	if err != nil {
		operation.failed()
	} else {
		operation.insertCollections(collections)
	}
}

func (operation *HarvestingCollections) insertCollections(collections []wisply.Collectioner) {
	repository := operation.Operation.GetRepository()
	task := newInsertCollectionsTask(operation, repository)
	err := task.Insert(collections)
	if err != nil {
		operation.failed()
	}
	oldCollections := operation.process.Collections
	newCollections := len(collections)
	difference := newCollections - oldCollections
	if difference != 0 {
		err = operation.process.updateCollections(difference)
	}
	if err != nil {
		operation.failed()
	} else {
		operation.succedded()
	}
}

// constructor
func newHarvestingCollections(harvestProcess *Process) Operationer {
	return &HarvestingCollections{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Collections"),
	}
}
