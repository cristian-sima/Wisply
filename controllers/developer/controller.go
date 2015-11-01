package developer

import "github.com/cristian-sima/Wisply/controllers/wisply"

var messages = map[string]string{
	"tableNotAllowed": "Wisply does not know this table :(",
}

// Controller represents the basic API controller
type Controller struct {
	wisply.Controller
}

// ShowHomePage shows the static home page for API
func (controller Controller) ShowHomePage() {
	controller.SetCustomTitle("API & Developers")
	// Please use http://www.timestampgenerator.com/ for getting the timestamp
	controller.IndicateLastModification(1441987477)
	controller.Layout = "site/public-layout.tpl"
	controller.TplNames = "site/developer/developer.tpl"
}
