package harvest

import wisply "github.com/cristian-sima/Wisply/models/wisply/data"

// HarvestingFormats is the operation which collects the formats from a remote repository
type HarvestingFormats struct {
	*HarvestingOperation
}

// Start marks the local repository status of the local repository as
// updating and gets the formats
func (operation *HarvestingFormats) Start() {
	operation.Operation.Start()
	operation.ChangeRepositoryStatus("updating")
	operation.tryToGet()
}

func (operation *HarvestingFormats) tryToGet() {
	rem := operation.GetRemote()
	task := newGetTask(operation, rem)
	content, err := task.GetFormats()
	if err != nil {
		operation.failed()
	} else {
		operation.tryToParse(content)
	}
}

func (operation *HarvestingFormats) tryToParse(content []byte) {
	rem := operation.GetRemote()
	task := newParseTask(operation, rem)
	formats, err := task.GetFormats(content)
	if err != nil {
		operation.failed()
	} else {
		operation.insertFormats(formats)
	}
}

func (operation *HarvestingFormats) insertFormats(formats []wisply.Formater) {
	repository := operation.Operation.GetRepository()
	task := newInsertFormatsTask(operation, repository)
	err := task.Insert(formats)
	if err != nil {
		operation.failed()
	}
	oldFormats := operation.process.Formats
	newFormats := len(formats)
	difference := newFormats - oldFormats
	if difference != 0 {
		err = operation.process.updateFormats(difference)
	}
	if err != nil {
		operation.failed()
	} else {
		operation.succedded()
	}
}

// constructor
func newHarvestingFormats(harvestProcess *Process) Operationer {
	return &HarvestingFormats{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Formats"),
	}
}
