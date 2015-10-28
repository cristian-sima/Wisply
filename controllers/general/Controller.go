package general

import (
	"net"
	"strings"

	"github.com/astaxie/beego"
	captchaModel "github.com/cristian-sima/Wisply/models/captcha"
	"github.com/dchest/captcha"
)

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

// RegisterCaptchaAction register a combination of ip and page
func (controller *Controller) RegisterCaptchaAction(page string) {
	captchaModel.RegisterAction(page, controller.getIP())
}

func (controller *Controller) getIP() string {
	ip, _, _ := net.SplitHostPort(controller.Ctx.Request.RemoteAddr)
	return ip
}

// IsCaptchaRequired checks if the application is running in production mode
// and for that page and ip the number has not exceeded the allowed one
func (controller *Controller) IsCaptchaRequired(page string) bool {
	// if !controller.IsProductionMode() {
	// 	return false
	// }
	return captchaModel.RequireCaptcha(page, controller.getIP())
}

// IsCaptchaValid checks (when the app is not in the developing mode) if the
// captcha is correct for a particular name
func (controller *Controller) IsCaptchaValid(page string) bool {
	value := strings.TrimSpace(controller.GetString(page + "-captcha-value"))
	id := strings.TrimSpace(controller.GetString(page + "-captcha-id"))
	// if !controller.IsProductionMode() {
	// 	return true
	// }
	return captcha.VerifyString(id, value)

}
