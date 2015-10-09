package harvest

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
	repository := operation.process.GetRepository()

	// create a task to request the server
	task := newGetRequestTask(operation, repository)

	content, err := task.RequestCollections()

	if err != nil {
		operation.harvestFailed()
	} else {
		operation.tryToParse(content)
	}
}

func (operation *HarvestingCollections) tryToParse(page []byte) {
	task := newParseRequestTask(operation)
	_, err := task.GetCollections(page)
	if err != nil {
		operation.harvestFailed()
	} else {
		operation.harvestFailed()
		//operation.insertFormats(collections)
	}
}

func (operation *HarvestingCollections) insertFormats(formats []Formater) {
	repository := operation.Operation.GetRepository()
	task := newInsertFormatsTask(operation, repository)
	err := task.Insert(formats)
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
