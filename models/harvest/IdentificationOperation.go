package harvest

// IdentificationOperation encapsulates the methods for requesting information from the repository
type IdentificationOperation struct {
	*Operation
}

// Start starts the action. Gets the page, get the content, clean database and store it
func (operation *IdentificationOperation) Start() {
	operation.Operation.Start()
	operation.ChangeRepositoryStatus("initializing")
	operation.tryToGet()
}

func (operation *IdentificationOperation) tryToGet() {

	// remote := operation.Process.GetRemote()
	repository := operation.process.GetRepository()

	// create a task to request the server
	task := newGetRequestTask(operation, repository)

	page, err := task.Identify()

	if err != nil {
		operation.identificationFailed()
	} else {
		operation.tryToParse(page)
	}
	// create a task to check the result
}

func (operation *IdentificationOperation) tryToParse(page []byte) {
	task := newParseRequestTask(operation)
	response, err := task.GetIdentification(page)
	if err != nil {
		operation.identificationFailed()
	} else {
		operation.insertIdentification(response)
	}
}

func (operation *IdentificationOperation) insertIdentification(result *Identificationer) {
	task := newInsertIdentificationTask(operation, operation.GetRepository())
	err := task.Insert(result)
	if err != nil {
		operation.identificationFailed()
	} else {
		operation.identificationSucceded()
	}
}

func (operation *IdentificationOperation) identificationFailed() {
	operation.ChangeRepositoryStatus("verification-failed")
	operation.ChangeResult("danger")
	operation.Finish()
}

func (operation *IdentificationOperation) identificationSucceded() {
	operation.ChangeRepositoryStatus("ok")
	operation.Finish()
}

func newIdentificationOperation(harvestProcess *Process) Operationer {
	return &IdentificationOperation{
		Operation: &Operation{
			process:   harvestProcess,
			Operation: newOperation(harvestProcess.Process, "Identifying"),
		},
	}
}
