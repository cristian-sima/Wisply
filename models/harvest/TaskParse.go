package harvest

import (
	"strconv"

	"github.com/cristian-sima/Wisply/models/harvest/remote"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// ParseTask represents a task for parsing the content of a request from a
// remote repository
// It is a decorator pattern
type ParseTask struct {
	*remoteTask
}

// Verify checks if the content is valid or not as a result of parsing
func (task *ParseTask) Verify(content []byte) error {

	task.addContent("verify")
	err := task.remote.IsValidResponse(content)
	task.finishRequest(err, "The content is valid")

	return err
}

// GetIdentification tells the remote server to parse the content and to
// return the identification
func (task *ParseTask) GetIdentification(content []byte) (*wisply.Identificationer, error) {

	task.addContent("identification")
	identification, err := task.remote.GetIdentification(content)
	task.finishRequest(err, "The identification has been parsed")

	return identification, err

}

// GetFormats tells the remote server to parse the content and to return the `formats`
func (task *ParseTask) GetFormats(content []byte) ([]wisply.Formater, error) {
	task.addContent("formats")
	formats, err := task.remote.GetFormats(content)
	number := strconv.Itoa(len(formats))
	task.finishRequest(err, "Success. "+number+" `formats` has been identified")

	return formats, err
}

// GetCollections tells the remote server to parse the content and to return
// the `collections`
func (task *ParseTask) GetCollections(content []byte) ([]wisply.Collectioner, error) {
	task.addContent("collections")
	collections, err := task.remote.GetCollections(content)
	number := strconv.Itoa(len(collections))
	task.finishRequest(err, "Success. "+number+" `collections` has been identified")

	return collections, err
}

// GetRecords tells the remote server to parse the content and to return the `records`
func (task *ParseTask) GetRecords(content []byte) ([]wisply.Recorder, error) {
	msg := ""
	task.addContent("records")
	records, err := task.remote.GetRecords(content)
	number := len(records)
	if number == 0 {
		msg = "There are no new records"
	} else {
		msg = "Success. " + strconv.Itoa(number) + " `records` has been identified"
	}
	task.finishRequest(err, msg)

	return records, err
}

// GetIdentifiers tells the remote server to parse the content and to return the `identifiers`
func (task *ParseTask) GetIdentifiers(content []byte) ([]wisply.Identifier, error) {
	msg := ""
	task.addContent("identifiers")
	identifiers, err := task.remote.GetIdentifiers(content)
	number := len(identifiers)
	if number == 0 {
		msg = "There are no new identifiers"
	} else {
		msg = "Success. " + strconv.Itoa(number) + " `identifiers` has been identified"
	}
	task.finishRequest(err, msg)

	return identifiers, err
}

func newParseTask(operationHarvest Operationer, remoteRepository remote.RepositoryInterface) *ParseTask {
	return &ParseTask{
		remoteTask: &remoteTask{
			Task: &Task{
				operation: operationHarvest,
				Task:      newTask(operationHarvest.GetOperation(), "Remote Request"),
			},
			remote: remoteRepository,
			name:   "Parse",
		},
	}
}
