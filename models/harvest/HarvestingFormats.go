package harvest

// HarvestingFormats encapsulates the methods for harvesting the formats
type HarvestingFormats struct {
	*Operation
}

// Start starts the action. Gets the page, get the content, clean database and store it
func (operation *HarvestingFormats) Start() {
	operation.Operation.Start()
	operation.ChangeRepositoryStatus("updating")
	operation.tryToGet()
}

func (operation *HarvestingFormats) tryToGet() {

	// remote := operation.Process.GetRemote()
	// repository := operation.process.GetRepository()

	// create a task to request the server
	newGetRequestTask(operation)

	// page, err := task.Identify(repository.URL)
	//
	// if err != nil {
	// 	operation.identificationFailed()
	// } else {
	// 	operation.identificationSucceded(page)
	// }
	// create a task to check the result
}

// func (operation *HarvestingOperation) tryToParse(page []byte) {
// 	task := newParseRequestTask(operation)
// 	response, err := task.GetIdentification(page)
// 	if err != nil {
// 		operation.identificationFailed()
// 	} else {
// 		operation.insertIdentification(response)
// 	}
// }
//
// func (operation *HarvestingOperation) insertIdentification(result *Identificationer) {
// 	task := newInsertIdentificationTask(operation, operation.GetRepository())
// 	err := task.Insert(result)
// 	if err != nil {
// 		operation.identificationFailed()
// 	} else {
// 		operation.identificationSucceded()
// 	}
// }
//
// func (operation *HarvestingOperation) identificationFailed() {
// 	operation.ChangeRepositoryStatus("verification-failed")
// 	operation.ChangeResult("danger")
// 	operation.Finish()
// }
//
// func (operation *HarvestingOperation) identificationSucceded() {
// 	operation.ChangeRepositoryStatus("ok")
// 	operation.Finish()
// }
//

func newHarvestingFormats(harvestProcess *Process) Operationer {
	return &HarvestingFormats{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "Harvesting"),
		},
	}
}
