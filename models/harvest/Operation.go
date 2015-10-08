package harvest

import (
	"fmt"

	action "github.com/cristian-sima/Wisply/models/action"
)

// Operationer ... defines the set of methods which should be implemented by the harvest operations
type Operationer interface {
	Start()
	GetOperation() *action.Operation
}

// Operation represents an operation
type Operation struct {
	*action.Operation
}

// Finish calls its finish parents method
func (operation *Operation) Finish() {
	fmt.Println("I am finishing")
	operation.Operation.Finish()
}

func newOperation(process *action.Process, content string) *action.Operation {
	return &*process.CreateOperation("Verification")
}
