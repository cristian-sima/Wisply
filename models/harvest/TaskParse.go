package harvest

import (
	"strconv"

	"github.com/cristian-sima/Wisply/models/harvest/remote"
	"github.com/cristian-sima/Wisply/models/harvest/wisply"
)

// ParseTask represents a task for parsing the content of a request from a remote repository
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

// GetIdentification tells the remote server to parse the content and to return the identification
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
	task.finishRequest(err, "Success. "+number+" formats has been identified")

	return formats, err
}

// GetCollections tells the remote server to parse the content and to return the `collections`
func (task *ParseTask) GetCollections(content []byte) ([]wisply.Collectioner, error) {
	task.addContent("collections")
	collections, err := task.remote.GetCollections(content)
	number := strconv.Itoa(len(collections))
	task.finishRequest(err, "Success. "+number+" collections has been identified")

	return collections, err
}

// GetRecords tells the remote server to parse the content and to return the `records`
func (task *ParseTask) GetRecords(content []byte) ([]wisply.Recorder, error) {
	task.addContent("records")
	records, err := task.remote.GetRecords(content)
	number := strconv.Itoa(len(records))
	task.finishRequest(err, "Success. "+number+" records has been identified")

	return records, err
}

// // GetIdentifiers returns an array with identifiers
// func (task *ParseRequestTask) GetIdentifiers(content []byte) ([]Identifier, error) {
//
// 	var identifiers []Identifier
//
// 	response, err := task.parse(content)
// 	if err != nil {
// 		return identifiers, err
// 	}
//
// 	remoteIdentifiers := response.ListIdentifiers.Headers
//
// 	for _, remoteIdentifier := range remoteIdentifiers {
//
// 		identifier := &remote.OAIIdentifier{
// 			Identifier: remoteIdentifier.Identifier,
// 			Datestamp:  remoteIdentifier.DateStamp,
// 			Spec:       remoteIdentifier.SetSpec,
// 		}
// 		identifiers = append(identifiers, identifier)
// 	}
//
// 	number := strconv.Itoa(len(identifiers))
// 	task.Finish(number + " identifiers has been identified")
//
// 	return identifiers, nil
// }
//
// func (task *ParseRequestTask) getKeys(plainText []byte) (*remote.Keys, error) {
//
// 	keys := &remote.Keys{}
//
// 	// Unmarshall all the data
// 	err := xml.Unmarshal(plainText, &keys)
// 	if err != nil {
// 		return keys, err
// 	}
//
// 	return keys, nil
// }

func newParseTask(operationHarvest Operationer, remoteRepository remote.RepositoryInterface) *ParseTask {
	return &ParseTask{
		remoteTask: &remoteTask{
			Task: &Task{
				operation: operationHarvest,
				Task:      newTask(operationHarvest.GetOperation(), "Remote Request"),
			},
			remote: remoteRepository,
			name:   "Parse ",
		},
	}
}
