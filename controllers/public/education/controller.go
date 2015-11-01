package education

import "github.com/cristian-sima/Wisply/controllers/public"

// Controller manages the operations for education which are public
type Controller struct {
	public.Controller
}

// Prepare sets the path for the package
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("public/education")
}
