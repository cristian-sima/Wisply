package harvest

import "github.com/cristian-sima/Wisply/controllers/admin/log"

// Controller is the parent for harvest logs
type Controller struct {
	log.Controller
}

// Prepare changes the path
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/log/harvest")
}
