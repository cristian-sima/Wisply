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

// GetRecords returns the records of the remote repository
// It waits for the answer and it modifies it state according to it
func (task *GetTask) GetRecords() ([]byte, error) {
	task.addContent("Get Records")
	body, err := task.remote.ListRecords()
	task.finishRequest(err, "Request performed")
	return body, err
}

// GetIdentifiers returns the identifiers of the remote repository
// It waits for the answer and it modifies it state according to it
func (task *GetTask) GetIdentifiers() ([]byte, error) {
	task.addContent("Get Identifiers")
	body, err := task.remote.ListIdentifiers()
	task.finishRequest(err, "Request performed")
	return body, err
}

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
