package harvest

import "github.com/cristian-sima/Wisply/models/harvest/wisply"

// HarvestingIdentifiers is the operation which collects the identifiers from a remote repository
type HarvestingIdentifiers struct {
	*HarvestingOperation
}

// Start gets the identifiers
func (operation *HarvestingIdentifiers) Start() {
	operation.Operation.Start()
	operation.start()
}

func (operation *HarvestingIdentifiers) start() {
	err := operation.clear()
	if err != nil {
		operation.failed()
	} else {
		operation.multiRequest()
	}
}

func (operation *HarvestingIdentifiers) clear() error {
	rem := operation.Operation.GetRepository()
	task := newInsertIdentifiersTask(operation, rem)
	return task.Clear()
}

func (operation *HarvestingIdentifiers) multiRequest() {
	var (
		hasMoreIdentifiers bool
		err                error
	)
	token := operation.getCurrentToken()
	hasMoreIdentifiers = true // in order to enter in the loop
	for hasMoreIdentifiers && (err == nil) {
		err := operation.tryToGet(token)
		if err == nil {
			token, hasMoreIdentifiers = operation.GetRemote().GetNextPage()
		}
		if hasMoreIdentifiers {
			operation.process.SaveToken("identifiers", token)
		}
	}
	if err != nil {
		operation.failed()
	} else {
		operation.succeeded()
	}
}

func (operation *HarvestingIdentifiers) tryToGet(token string) error {
	rem := operation.GetRemote()
	task := newGetTask(operation, rem)
	content, err := task.GetIdentifiers(token)
	if err != nil {
		return err
	}
	return operation.tryToParse(content)
}

func (operation *HarvestingIdentifiers) tryToParse(page []byte) error {
	rem := operation.GetRemote()
	task := newParseTask(operation, rem)
	identifiers, err := task.GetIdentifiers(page)
	if err != nil {
		return err
	}
	return operation.insertIdentifiers(identifiers)
}

func (operation *HarvestingIdentifiers) insertIdentifiers(identifiers []wisply.Identifier) error {
	repository := operation.Operation.GetRepository()
	task := newInsertIdentifiersTask(operation, repository)
	err := task.Insert(identifiers)
	if err != nil {
		return err
	}
	err = operation.process.updateIdentifiers(len(identifiers))
	return err
}

// constructor
func newHarvestingIdentifiers(harvestProcess *Process) Operationer {
	return &HarvestingIdentifiers{
		HarvestingOperation: newHarvestingOperation(harvestProcess, "Identifiers"),
	}
}
