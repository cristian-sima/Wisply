package wisply

import (
	"net"
	"strings"

	"github.com/astaxie/beego"
	captchaModel "github.com/cristian-sima/Wisply/models/captcha"
	"github.com/dchest/captcha"
)

// Adapter is the objects which connects Wisply with the framework Adapter
// It is the most basic type of Adapter
type Adapter struct {
	beego.Controller
	templatePath string
}

// LoadTemplate loads a filenmae for that template
// The filename should be in the same directory specified by the templatePath
// The templatePath can be changed by using SetTemplatePath
func (controller *Controller) LoadTemplate(filename string) {
	if controller.templatePath == "" {
		panic("This controller did not overwrite the default path")
	} else {
		controller.TplNames = templatesFolder + controller.templatePath + "/" + filename + ".tpl"
	}
}

// SetTemplatePath changes the default path for template
func (controller *Controller) SetTemplatePath(path string) {
	controller.templatePath = path
}

// SetLayout sets the layout for the controller
// Please see beego for more information about layout
func (controller *Controller) SetLayout(layout string) {
	controller.Layout = templatesFolder + layout + "-layout.tpl"
}

// RemoveLayout removes the layout
func (controller *Controller) RemoveLayout() {
	controller.Layout = ""
}

// GetApplicationMode returns the mode in which the app is running
// dev - for developing
func (Adapter *Adapter) GetApplicationMode() string {
	return beego.AppConfig.String("runmode")
}

// IsProductionMode checks if the application is running in the production mode
func (Adapter *Adapter) IsProductionMode() bool {
	return Adapter.GetApplicationMode() == "pro"
}

// IsDevelopingMode checks if the application is running in the developing mode
func (Adapter *Adapter) IsDevelopingMode() bool {
	return Adapter.GetApplicationMode() == "dev"
}

// RegisterCaptchaAction register a combination of ip and page
func (Adapter *Adapter) RegisterCaptchaAction(page string) {
	captchaModel.RegisterAction(page, Adapter.getIP())
}

func (Adapter *Adapter) getIP() string {
	ip, _, _ := net.SplitHostPort(Adapter.Ctx.Request.RemoteAddr)
	return ip
}

// IsCaptchaRequired checks if the application is running in production mode
// and for that page and ip the number has not exceeded the allowed one
func (Adapter *Adapter) IsCaptchaRequired(page string) bool {
	if !Adapter.IsProductionMode() {
		return false
	}
	return captchaModel.RequireCaptcha(page, Adapter.getIP())
}

// IsCaptchaValid checks (when the app is not in the developing mode) if the
// captcha is correct for a particular name
func (Adapter *Adapter) IsCaptchaValid(page string) bool {
	value := strings.TrimSpace(Adapter.GetString(page + "-captcha-value"))
	id := strings.TrimSpace(Adapter.GetString(page + "-captcha-id"))
	if !Adapter.IsProductionMode() {
		return true
	}

	if !Adapter.IsCaptchaRequired(page) {
		return true
	}
	return captcha.VerifyString(id, value)

}

// LoadCaptcha checks if the captcha image is needed. If so, it loads in
// to the template
func (Adapter *Adapter) LoadCaptcha(page string) {
	if Adapter.IsCaptchaRequired(page) {
		Adapter.Data["captcha"] = captchaModel.New()
		Adapter.Data["showCaptcha"] = true
	}
}
