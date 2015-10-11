package harvest

// HarvestingIdentifiers is the operation which collects the identifiers from a remote repository
type HarvestingIdentifiers struct {
	*Operation
}

// Start gets the identifiers
func (operation *HarvestingIdentifiers) Start() {
	operation.Operation.Start()
	operation.tryToGet()
}

func (operation *HarvestingIdentifiers) tryToGet() {

	// rem := operation.process.GetRemote()

	// create a task to request the server
	// task := newGetRequestTask(operation, rem)

	// content, err := task.RequestIdentifiers()
	//
	// if err != nil {
	// 	operation.harvestFailed()
	// } else {
	// 	operation.tryToParse(content)
	// }
}

func (operation *HarvestingIdentifiers) tryToParse(page []byte) {
	// task := newParseRequestTask(operation)
	// collections, err := task.GetIdentifiers(page)
	// if err != nil {
	// 	operation.harvestFailed()
	// } else {
	// 	operation.insertRecords(collections)
	// }
}

func (operation *HarvestingIdentifiers) insertRecords(collections []Identifier) {
	repository := operation.Operation.GetRepository()
	task := newInsertIdentifiersTask(operation, repository)
	err := task.Insert(collections)
	if err != nil {
		operation.harvestFailed()
	} else {
		operation.harvestSuccess()
	}
}

func (operation *HarvestingIdentifiers) harvestFailed() {
	operation.ChangeResult("danger")
	operation.Finish()
}

func (operation *HarvestingIdentifiers) harvestSuccess() {
	operation.Finish()
}

// constructor
func newHarvestingIdentifiers(harvestProcess *Process) Operationer {
	return &HarvestingIdentifiers{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "Harvest Identifiers"),
		},
	}
}
