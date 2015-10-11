package harvest

// HarvestingRecords is the operation which collects the records from a remote repository
type HarvestingRecords struct {
	*Operation
}

// Start gets the collections
func (operation *HarvestingRecords) Start() {
	operation.Operation.Start()
	operation.tryToGet()
}

func (operation *HarvestingRecords) tryToGet() {

	// rem := operation.process.GetRemote()

	// create a task to request the server
	// task := newGetRequestTask(operation, rem)

	// content, err := task.RequestRecords()
	//
	// if err != nil {
	// 	operation.harvestFailed()
	// } else {
	// 	operation.tryToParse(content)
	// }
}

func (operation *HarvestingRecords) tryToParse(page []byte) {
	task := newParseRequestTask(operation)
	collections, err := task.GetRecords(page)
	if err != nil {
		operation.harvestFailed()
	} else {
		operation.insertRecords(collections)
	}
}

func (operation *HarvestingRecords) insertRecords(collections []Recorder) {
	repository := operation.Operation.GetRepository()
	task := newInsertRecordsTask(operation, repository)
	err := task.Insert(collections)
	if err != nil {
		operation.harvestFailed()
	} else {
		operation.harvestSuccess()
	}
}

func (operation *HarvestingRecords) harvestFailed() {
	operation.ChangeResult("danger")
	operation.Finish()
}

func (operation *HarvestingRecords) harvestSuccess() {
	operation.Finish()
}

// constructor
func newHarvestingRecords(harvestProcess *Process) Operationer {
	return &HarvestingRecords{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "Harvest Records"),
		},
	}
}
