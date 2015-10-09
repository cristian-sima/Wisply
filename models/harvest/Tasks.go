package harvest

import oai "github.com/cristian-sima/Wisply/models/harvest/protocols/oai"

// GetRequestTask represents a task for requesting something
type GetRequestTask struct {
	Tasker
	*Task
}

// Identify returns the content for the Identify request
func (task *GetRequestTask) Identify(URL string) ([]byte, error) {

	request := &oai.Request{
		BaseURL: URL,
		Verb:    "Identify",
	}
	return task.get(request)
}

// get returns the content of the remote repository
func (task *GetRequestTask) get(request *oai.Request) ([]byte, error) {

	task.Content = "HTTP Request " + request.Verb
	body, err := request.Get()

	if err != nil {
		task.ChangeResult("danger")
		task.Finish(err.Error())
	} else {
		task.Finish("Success")
	}
	return body, err
}

// ParseRequestTask represents a task for parsing the request
type ParseRequestTask struct {
	Tasker
	*Task
}

// GetIdentification returns the identification parsed of a page
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

// Verify checks if the content is valid or not
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

func newGetRequestTask(operationHarvest Operationer) *GetRequestTask {
	return &GetRequestTask{
		Task: &Task{
			operation: operationHarvest,
			Task:      newTask(operationHarvest.GetOperation(), "request verification"),
		},
	}
}
