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

// GetApplicationMode returns the mode in which the app is running
// dev - for developing
func (controller *Controller) GetApplicationMode() string {
	return beego.AppConfig.String("runmode")
}

// IsProductionMode checks if the application is running in the production mode
func (controller *Controller) IsProductionMode() bool {
	return controller.GetApplicationMode() == "pro"
}

// IsDevelopingMode checks if the application is running in the developing mode
func (controller *Controller) IsDevelopingMode() bool {
	return controller.GetApplicationMode() == "dev"
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
	if !controller.IsProductionMode() {
		return false
	}
	return captchaModel.RequireCaptcha(page, controller.getIP())
}

// IsCaptchaValid checks (when the app is not in the developing mode) if the
// captcha is correct for a particular name
func (controller *Controller) IsCaptchaValid(page string) bool {
	value := strings.TrimSpace(controller.GetString(page + "-captcha-value"))
	id := strings.TrimSpace(controller.GetString(page + "-captcha-id"))
	if !controller.IsProductionMode() {
		return true
	}

	if !controller.IsCaptchaRequired(page) {
		return true
	}
	return captcha.VerifyString(id, value)

}

// LoadCaptcha checks if the captcha image is needed. If so, it loads in
// to the template
func (controller *Controller) LoadCaptcha(page string) {
	if controller.IsCaptchaRequired(page) {
		controller.Data["captcha"] = captchaModel.New()
		controller.Data["showCaptcha"] = true
	}
}
