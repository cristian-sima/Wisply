package tools

import "github.com/cristian-sima/Wisply/controllers/developer"

// Controller manages the operations for the tools
type Controller struct {
	developer.Controller
}

// Prepare sets the path for the package
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("developer/tools")
}
