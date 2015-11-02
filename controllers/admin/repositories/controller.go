package repositories

import "github.com/cristian-sima/Wisply/controllers/admin"

// Controller manages the operations for the repositories
type Controller struct {
	admin.Controller
}

// Prepare sets the path for the package
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/repositories")
}
