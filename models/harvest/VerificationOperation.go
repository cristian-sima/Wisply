package harvest

import (
	action "github.com/cristian-sima/Wisply/models/action"
)

// VerificationOperation encapsulates the methods for validating the repository
type VerificationOperation struct {
	Operation
}

// Operation represents a harvest operation
type Operation struct {
	*action.Operation
}

// CreateOperation creates a new harvest operation
// func CreateOperation(process *action.Process) *Operation {
// 	return &Operation{
// 		Operation: &*action.CreateOperation(process),
// 	}
// }
