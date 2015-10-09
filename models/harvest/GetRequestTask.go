package harvest

import (
	"fmt"

	oai "github.com/cristian-sima/Wisply/models/harvest/protocols/oai"
)

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

// Parse returns the content of the remote repository
func (task *ParseRequestTask) Parse(content []byte) (*oai.Response, error) {
	fmt.Println(":)")
	request := &oai.Request{
		Verb: "Identify",
	}
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
