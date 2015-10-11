package harvest

import "github.com/cristian-sima/Wisply/models/harvest/remote"

// ParseRequestTask represents a task for parsing the content of a request from a remote repository
// It is a decorator pattern
type ParseRequestTask struct {
	Tasker
	*Task
	remote remote.RepositoryInterface
}

func (task *ParseRequestTask) setContent(verb string) {
	task.Content = "Parse " + verb
}

func (task *ParseRequestTask) finishRequest(err error, success string) {
	if err != nil {
		task.ChangeResult("danger")
		task.Finish(err.Error())
	} else {
		task.Finish(success)
	}
}

// Verify checks if the content is valid or not as a result of parsing
func (task *ParseRequestTask) Verify(content []byte) error {

	task.setContent("verify")
	err := task.remote.IsValidResponse(content)
	task.finishRequest(err, "The response has been parsed")

	return err
}

//
// // GetIdentification returns an identificationer interface as a result of parsing the identification of the page
// func (task *ParseRequestTask) GetIdentification(content []byte) (*Identificationer, error) {
//
// 	var idenfitication Identificationer
//
// 	response, err := task.parse(content)
// 	if err != nil {
// 		return &idenfitication, err
// 	}
//
// 	remoteIdentity := response.Identify
//
// 	idenfitication = &remote.OAIIdentification{
// 		Name:              remoteIdentity.RepositoryName,
// 		URL:               remoteIdentity.BaseURL,
// 		Protocol:          remoteIdentity.ProtocolVersion,
// 		AdminEmails:       remoteIdentity.AdminEmail,
// 		EarliestDatestamp: remoteIdentity.EarliestDatestamp,
// 		RecordPolicy:      remoteIdentity.DeletedRecord,
// 		Granularity:       remoteIdentity.Granularity,
// 	}
//
// 	task.Finish("The identification has been parsed")
//
// 	return &idenfitication, nil
// }
//
// // GetFormats returns a formater interface as a result of parsing the formats of the page
// func (task *ParseRequestTask) GetFormats(content []byte) ([]Formater, error) {
//
// 	var formats []Formater
//
// 	response, err := task.parse(content)
// 	if err != nil {
// 		return formats, err
// 	}
//
// 	remoteFormats := response.ListMetadataFormats.MetadataFormat
//
// 	for _, format := range remoteFormats {
// 		format := &remote.OAIFormat{
// 			Prefix:    format.MetadataPrefix,
// 			Namespace: format.MetadataNamespace,
// 			Schema:    format.Schema,
// 		}
// 		formats = append(formats, format)
// 	}
//
// 	number := strconv.Itoa(len(formats))
// 	task.Finish(number + " has been identified")
//
// 	return formats, nil
// }
//
// // GetCollections returns an array with all the collections
// func (task *ParseRequestTask) GetCollections(content []byte) ([]Collectioner, error) {
//
// 	var collections []Collectioner
//
// 	response, err := task.parse(content)
// 	if err != nil {
// 		return collections, err
// 	}
//
// 	remoteCollections := response.ListSets.Set
//
// 	for _, collection := range remoteCollections {
// 		collection := &remote.OAICollection{
// 			Name: collection.SetName,
// 			Spec: collection.SetSpec,
// 		}
// 		collections = append(collections, collection)
// 	}
//
// 	number := strconv.Itoa(len(collections))
// 	task.Finish(number + " has been identified")
//
// 	return collections, nil
// }
//
// // GetRecords returns an array with records
// func (task *ParseRequestTask) GetRecords(content []byte) ([]Recorder, error) {
//
// 	var records []Recorder
//
// 	response, err := task.parse(content)
// 	if err != nil {
// 		return records, err
// 	}
//
// 	remoteRecords := response.ListRecords.Records
//
// 	for _, record := range remoteRecords {
//
// 		keys, _ := task.getKeys(record.Metadata.Body)
//
// 		record := &remote.OAIRecord{
// 			Identifier: record.Header.Identifier,
// 			Datestamp:  record.Header.DateStamp,
// 			Keys:       keys,
// 		}
// 		records = append(records, record)
// 	}
//
// 	number := strconv.Itoa(len(records))
// 	task.Finish(number + " records has been identified")
//
// 	return records, nil
// }
//
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

func newParseRequestTask(operationHarvest Operationer, remoteRepository remote.RepositoryInterface) *ParseRequestTask {
	return &ParseRequestTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Parse request content"),
		},
		remote: remoteRepository,
	}
}
