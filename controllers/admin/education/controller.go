package education

import "github.com/cristian-sima/Wisply/controllers/admin"

// Controller manages the operations for education
type Controller struct {
	admin.Controller
}

// Prepare changes the path
func (controller *Controller) Prepare() {
	controller.Prepare()
	controller.SetTemplatePath("education")
}
