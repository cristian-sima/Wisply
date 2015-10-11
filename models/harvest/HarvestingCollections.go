package harvest

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

// HarvestingCollections is the operation which collects the collections from a remote repository
type HarvestingCollections struct {
	*Operation
}

// Start gets the collections
func (operation *HarvestingCollections) Start() {
	operation.Operation.Start()
	operation.tryToGet()
}

func (operation *HarvestingCollections) tryToGet() {

	// remote := operation.Process.GetRemote()
	// rem := operation.process.GetRemote()

	// create a task to request the server
	// task := newGetRequestTask(operation, rem)

	// content, err := task.RequestCollections()
	//
	// if err != nil {
	// 	operation.harvestFailed()
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
		operation.harvestFailed()
	} else {
		operation.harvestSuccess()
	}
}

func (operation *HarvestingCollections) harvestFailed() {
	operation.ChangeResult("danger")
	operation.Finish()
}

func (operation *HarvestingCollections) harvestSuccess() {
	operation.Finish()
}

// constructor
func newHarvestingCollections(harvestProcess *Process) Operationer {
	return &HarvestingCollections{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "Harvest Collections"),
		},
	}
}
