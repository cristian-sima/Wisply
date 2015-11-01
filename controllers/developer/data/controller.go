package data

import "github.com/cristian-sima/Wisply/controllers/developer"

// Controller manages the operations for exporting data
type Controller struct {
	developer.Controller
}

// Prepare changes the path
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("developer/data")
}
