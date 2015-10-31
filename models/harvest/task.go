package harvest

import (
	action "github.com/cristian-sima/Wisply/models/action"
	"github.com/cristian-sima/Wisply/models/harvest/remote"
)

// Tasker ... defines the set of methods which should be implemented by the harvest task
type Tasker interface {
	GetTask() *action.Task
}

// Task represents a harvest task
type Task struct {
	Tasker
	*action.Task
	remote    *remote.RepositoryInterface
	operation Operationer // it is the harvest operation
}

func newTask(operation *action.Operation, content string) *action.Task {
	return &*operation.CreateTask(content)
}

// remoteTask represents a task that manages the operations with the remote repository
type remoteTask struct {
	Tasker
	*Task
	remote remote.RepositoryInterface
	name   string
}

func (task *remoteTask) addContent(verb string) {
	task.Content = task.name + " " + verb
}

func (task *remoteTask) finishRequest(err error, success string) {
	if err != nil {
		task.ChangeResult("danger")
		task.Finish(err.Error())
	} else {
		task.Finish(success)
	}
}
