package log

import "github.com/cristian-sima/Wisply/controllers/admin"

// Controller is the parent for all the log controllers
type Controller struct {
	admin.Controller
}

// Prepare changes the path
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/log")
}
