package process

import (
	"strconv"

	"github.com/cristian-sima/Wisply/models/action"
)

// Operation represents an operation from a harvesting process
type Operation struct {
	Controller
	operation *action.Operation
}

// Display displays the operation page
func (controller *Operation) Display() {
	operation := controller.GetOperation()
	controller.Data["tasks"] = operation.GetTasks()
	controller.LoadTemplate("operation")
}

// Prepare loads the process
func (controller *Operation) Prepare() {
	controller.Controller.Prepare()
	controller.loadOperation()
}

// GetOperation returns the reference to the operation
func (controller *Operation) GetOperation() *action.Operation {
	return controller.operation
}

func (controller *Operation) loadOperation() {
	ID := controller.Ctx.Input.Param(":operation")
	intID, err := strconv.Atoi(ID)
	if err == nil {
		operation := action.NewOperation(intID)
		controller.Data["operation"] = operation
		controller.operation = operation
		controller.SetCustomTitle("Operation #" + strconv.Itoa(operation.ID))
	}
}
