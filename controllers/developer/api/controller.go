package api

import "github.com/cristian-sima/Wisply/controllers/admin"

// Controller manages the operations for API
type Controller struct {
	admin.Controller
}

// Prepare changes the path
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("developer/api")
}
