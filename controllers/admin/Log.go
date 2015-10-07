package admin

import (
	"strconv"

	action "github.com/cristian-sima/Wisply/models/action"
)

// LogController manages the operations for showing the logs
type LogController struct {
	Controller
}

// ShowGeneralPage displays the processes
func (controller *LogController) ShowGeneralPage() {
	controller.Data["processes"] = action.GetAllProcesses()
	controller.SetCustomTitle("Admin - Event Log")
	controller.TplNames = "site/admin/log/home.tpl"
}

// ShowProcess displays the log of a process
func (controller *LogController) ShowProcess() {
	idString := controller.Ctx.Input.Param(":id")
	ID, _ := strconv.Atoi(idString)
	process := action.NewProcess(ID)
	if process.IsRunning {
		operation := action.NewOperation(process.GetCurrentOperation().ID)
		controller.Data["operation"] = operation
	}

	controller.Data["process"] = process
	// controller.Data["tasks"] = process.GetTasks();
	controller.TplNames = "site/admin/log/process.tpl"
}
