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
