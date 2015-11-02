package api

import "github.com/cristian-sima/Wisply/controllers/developer"

// Controller manages the operations for API
type Controller struct {
	developer.Controller
}

// Prepare changes the path
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("developer/api")
}
