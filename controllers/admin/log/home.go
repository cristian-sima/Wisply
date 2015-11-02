package log

import (
	"github.com/cristian-sima/Wisply/models/action"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Home represents manages the home page for log area
type Home struct {
	Controller
}

// Display shows the home page for the log
func (controller *Home) Display() {
	controller.SetCustomTitle("Admin - Log")
	controller.LoadTemplate("home")
	controller.Data["processes"] = action.GetAllProcesses()
}

// DisplayAdvanceOptions displays the advance options for the entire log
func (controller *Home) DisplayAdvanceOptions() {
	controller.SetCustomTitle("Log - Advance options")
	controller.LoadTemplate("advance-options")
}

// Delete deletes the entire log
func (controller *Home) Delete() {
	action.DeleteLog()
	repository.ResetAllRepositories()
	controller.LoadTemplate("advance-options")
}
