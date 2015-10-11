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

//
// // Identify returns the content for the Identify request
// func (task *GetRequestTask) Identify() ([]byte, error) {
// 	request := &oai.Request{
// 	// BaseURL: task.repository.URL,
// 	// Verb: "Identify",
// 	}
// 	return task.get(request)
// }
//
// // RequestFormats returns all the formats of the repository
// func (task *GetRequestTask) RequestFormats() ([]byte, error) {
// 	request := &oai.Request{
// 		BaseURL: task.repository.URL,
// 		Verb:    "ListMetadataFormats",
// 	}
// 	return task.get(request)
// }
//
// // RequestCollections returns all the collections of the repository
// func (task *GetRequestTask) RequestCollections() ([]byte, error) {
// 	request := &oai.Request{
// 		BaseURL: task.repository.URL,
// 		Verb:    "ListSets",
// 	}
// 	return task.get(request)
// }
//
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
