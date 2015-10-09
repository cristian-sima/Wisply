package harvest

import oai "github.com/cristian-sima/Wisply/models/harvest/protocols/oai"

// ParseRequestTask represents a task for parsing the request
type ParseRequestTask struct {
	Tasker
	*Task
}

// GetIdentification returns an identificationer interface as a result of parsing the identification of the page
func (task *ParseRequestTask) GetIdentification(content []byte) (*Identificationer, error) {

	var idenfitication Identificationer

	response, err := task.parse(content)
	if err != nil {
		return &idenfitication, err
	}

	remoteIdentity := response.Identify

	idenfitication = &OAIIdentification{
		Name:              remoteIdentity.RepositoryName,
		URL:               remoteIdentity.BaseURL,
		Protocol:          remoteIdentity.ProtocolVersion,
		AdminEmails:       remoteIdentity.AdminEmail,
		EarliestDatestamp: remoteIdentity.EarliestDatestamp,
		RecordPolicy:      remoteIdentity.DeletedRecord,
		Granularity:       remoteIdentity.Granularity,
	}

	return &idenfitication, nil
}

// Verify checks if the content is valid or not as a result of parsing
func (task *ParseRequestTask) Verify(content []byte) error {
	_, err := task.parse(content)
	return err
}

// parse returns the content of the remote repository
func (task *ParseRequestTask) parse(content []byte) (*oai.Response, error) {
	request := &oai.Request{}
	response, err := request.Parse(content)
	if err != nil {
		task.ChangeResult("danger")
		task.Finish(err.Error())
	} else {
		task.Finish("Success")
	}
	return response, err
}

func newParseRequestTask(operationHarvest Operationer) *ParseRequestTask {
	return &ParseRequestTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "Parse request content"),
		},
	}
}
