package data

import "github.com/cristian-sima/Wisply/controllers/admin"

// Controller manages the operations for exporting data
type Controller struct {
	admin.Controller
}

// Prepare changes the path
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("developer/data")
}
