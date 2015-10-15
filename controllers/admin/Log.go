package admin

import (
	"strconv"

	action "github.com/cristian-sima/Wisply/models/action"
	harvest "github.com/cristian-sima/Wisply/models/harvest"
	"github.com/cristian-sima/Wisply/models/repository"
)

// LogController manages the operations for showing the logs
type LogController struct {
	Controller
}

// ShowGeneralPage displays the processes
func (controller *LogController) ShowGeneralPage() {
	controller.Data["processes"] = action.GetAllProcesses()
	controller.SetCustomTitle("Admin - Event Log")
	controller.TplNames = "site/admin/log/all-processes.tpl"
}

// ShowProcessAdvanceOptions displays the advance options for a process
func (controller *LogController) ShowProcessAdvanceOptions() {
	controller.SetCustomTitle("Advance options")
	idString := controller.Ctx.Input.Param(":process")
	ID, _ := strconv.Atoi(idString)
	controller.Data["process"] = action.NewProcess(ID)
	controller.TplNames = "site/admin/log/process-advance-options.tpl"
}

// ShowLogAdvanceOptions displays the advance options for the entire log
func (controller *LogController) ShowLogAdvanceOptions() {
	controller.SetCustomTitle("Log - Advance options")
	controller.TplNames = "site/admin/log/log-advance-options.tpl"
}

// DeleteEntireLog deletes the entire log
func (controller *LogController) DeleteEntireLog() {
	action.DeleteLog()
	repository.ResetAllRepositories()
	controller.TplNames = "site/admin/log/log-advance-options.tpl"
}

// DeleteProcess deletes a process
func (controller *LogController) DeleteProcess() {
	idString := controller.Ctx.Input.Param(":process")
	ID, _ := strconv.Atoi(idString)
	process := action.NewProcess(ID)
	controller.deleteProcess(process)
}

// DeleteProcess deletes the process
func (controller *LogController) deleteProcess(process *action.Process) {

	process.Delete()
	controller.TplNames = "site/admin/log/log-advance-options.tpl"
}

// ShowProcess displays the log of a process
func (controller *LogController) ShowProcess() {
	controller.showProcess()
	controller.TplNames = "site/admin/log/process.tpl"
}

// ShowProgressHistory displays the entire log of a process
func (controller *LogController) ShowProgressHistory() {
	controller.showProcess()
	controller.TplNames = "site/admin/log/progress-history.tpl"
}

// ShowProcess displays the log of a process
func (controller *LogController) showProcess() {
	idString := controller.Ctx.Input.Param(":process")
	ID, _ := strconv.Atoi(idString)
	process := action.NewProcess(ID)

	if process.IsRunning {
		operationID := process.GetCurrentOperation().ID
		operation := action.NewOperation(operationID)
		controller.Data["operation"] = operation
	}

	controller.Data["process"] = process
	controller.Data["operations"] = process.GetOperations()

	switch process.Content {
	case "Harvest":
		controller.showHarvestProcess(process)
		break
	}
}

func (controller *LogController) showHarvestProcess(process *action.Process) {
	controller.Data["harvestProcess"] = harvest.NewProcess(process.Action.ID)
}

// ShowOperation display the
func (controller *LogController) ShowOperation() {
	processID := controller.Ctx.Input.Param(":process")
	operationID := controller.Ctx.Input.Param(":operation")

	IDProcess, _ := strconv.Atoi(processID)
	IDOperation, _ := strconv.Atoi(operationID)

	process := action.NewProcess(IDProcess)
	operation := action.NewOperation(IDOperation)

	controller.Data["process"] = process
	controller.Data["operation"] = operation
	controller.Data["tasks"] = operation.GetTasks()

	controller.TplNames = "site/admin/log/operation.tpl"
}
