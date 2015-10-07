package admin

import action "github.com/cristian-sima/Wisply/models/action"

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
