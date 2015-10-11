package harvest

import "github.com/cristian-sima/Wisply/models/harvest/remote"

// GetTask represents a task for requesting something from the remote repository
// It is a decorator pattern
type GetTask struct {
	*remoteTask
}

// Verify tells the remote to verify itself
// It waits for the answer and it modifies it state according to it
func (task *GetTask) Verify() ([]byte, error) {
	task.addContent("Verify")
	body, err := task.remote.Test()
	task.finishRequest(err, "Request performed")
	return body, err
}

// Identify tells the remote to send the identification details
// It waits for the answer and it modifies it state according to it
func (task *GetTask) Identify() ([]byte, error) {
	task.addContent("Verify")
	body, err := task.remote.Identify()
	task.finishRequest(err, "Request performed")
	return body, err
}

// GetFormats returns the formats of the remote repository
// It waits for the answer and it modifies it state according to it
func (task *GetTask) GetFormats() ([]byte, error) {
	task.addContent("Get Formats")
	body, err := task.remote.ListFormats()
	task.finishRequest(err, "Request performed")
	return body, err
}

// GetCollections returns the collections of the remote repository
// It waits for the answer and it modifies it state according to it
func (task *GetTask) GetCollections() ([]byte, error) {
	task.addContent("Get Collections")
	body, err := task.remote.ListCollections()
	task.finishRequest(err, "Request performed")
	return body, err
}

// // RequestRecords returns all the records of the repository
// func (task *GetRequestTask) RequestRecords() ([]byte, error) {
// 	request := &oai.Request{
// 		BaseURL:        task.repository.URL,
// 		Verb:           "ListRecords",
// 		MetadataPrefix: "oai_dc",
// 	}
// 	return task.get(request)
// }
//
// // RequestIdentifiers returns all the identifiers of the repository
// func (task *GetRequestTask) RequestIdentifiers() ([]byte, error) {
// 	request := &oai.Request{
// 		BaseURL:        task.repository.URL,
// 		Verb:           "ListIdentifiers",
// 		MetadataPrefix: "oai_dc",
// 	}
// 	return task.get(request)
// }
//
// // get returns the content of the remote repository
// func (task *GetRequestTask) get() ([]byte, error) {
//
// 	task.Content = "HTTP Request " + request.Verb
// 	body, err := request.Get()
//
// 	if err != nil {
// 		task.ChangeResult("danger")
// 		task.Finish(err.Error())
// 	} else {
// 		task.Finish("Success")
// 	}
// 	return body, err
// }
//
// // get returns the content of the remote repository
// func (task *GetRequestTask) get(request *oai.Request) ([]byte, error) {
//
// 	task.Content = "HTTP Request " + request.Verb
// 	body, err := request.Get()
//
// 	if err != nil {
// 		task.ChangeResult("danger")
// 		task.Finish(err.Error())
// 	} else {
// 		task.Finish("Success")
// 	}
// 	return body, err
// }

func newGetTask(operationHarvest Operationer, remoteRepository remote.RepositoryInterface) *GetTask {
	return &GetTask{
		remoteTask: &remoteTask{
			Task: &Task{
				operation: operationHarvest,
				Task:      newTask(operationHarvest.GetOperation(), "Remote Request"),
			},
			remote: remoteRepository,
			name:   "HTTP Request ",
		},
	}
}
