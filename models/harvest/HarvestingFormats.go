package harvest

// HarvestingFormats is the operation which collects the formats from a remote repository
type HarvestingFormats struct {
	*Operation
}

// Start marks the local repository status of the local repository as updating and gets the formats
func (operation *HarvestingFormats) Start() {
	operation.Operation.Start()
	operation.ChangeRepositoryStatus("updating")
	operation.tryToGet()
}

func (operation *HarvestingFormats) tryToGet() {

	// remote := operation.Process.GetRemote()
	repository := operation.process.GetRepository()

	// create a task to request the server
	task := newGetRequestTask(operation, repository)

	content, err := task.RequestFormats()

	if err != nil {
		operation.harvestFailed()
	} else {
		operation.tryToParse(content)
	}
}

func (operation *HarvestingFormats) tryToParse(page []byte) {
	task := newParseRequestTask(operation)
	_, err := task.GetFormats(page)
	if err != nil {
		operation.harvestFailed()
	} else {
		operation.harvestSuccess()
	}
}

func (operation *HarvestingFormats) harvestFailed() {
	operation.ChangeResult("danger")
	operation.Finish()
}

func (operation *HarvestingFormats) harvestSuccess() {
	operation.Finish()
}

// constructor
func newHarvestingFormats(harvestProcess *Process) Operationer {
	return &HarvestingFormats{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "Harvesting Formats"),
		},
	}
}
