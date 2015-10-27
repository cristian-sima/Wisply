package general

import "github.com/astaxie/beego"

// Controller is the objects which connects Wisply with the framework controller
type Controller struct {
	beego.Controller
}

// GetDeveloperMode returns the mode in which the app is running
// dev - for developing
func (controller *Controller) GetDeveloperMode() string {
	return beego.AppConfig.String("runmode")
}

// IsProductionMode checks if the application is running in the production mode
func (controller *Controller) IsProductionMode() bool {
	return controller.GetDeveloperMode() == "pro"
}

// IsDevelopingMode checks if the application is running in the developing mode
func (controller *Controller) IsDevelopingMode() bool {
	return controller.GetDeveloperMode() == "dev"
}
