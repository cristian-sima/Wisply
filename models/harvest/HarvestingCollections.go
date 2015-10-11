package harvest

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

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
	//
	// rem := operation.process.GetRemoteServer()
	//
	// task := newGetTask(operation, rem)
	//
	// content, err := task.GetCollections()
	//
	// if err != nil {
	// 	operation.failed()
	// } else {
	// 	operation.tryToParse(content)
	// }
}

func (operation *HarvestingCollections) tryToParse(page []byte) {
	// task := newParseRequestTask(operation)
	// collections, err := task.GetCollections(page)
	// if err != nil {
	// 	operation.harvestFailed()
	// } else {
	// 	operation.insertFormats(collections)
	// }
}

func (operation *HarvestingCollections) insertFormats(collections []wisply.Collectioner) {
	repository := operation.Operation.GetRepository()
	task := newInsertCollectionsTask(operation, repository)
	err := task.Insert(collections)
	if err != nil {
		operation.failed()
	} else {
		operation.succeeded()
	}
}

// constructor
func newHarvestingCollections(harvestProcess *Process) Operationer {
	return &HarvestingCollections{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "collections"),
	}
}
