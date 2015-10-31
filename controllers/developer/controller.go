// Package developer contains all the objects which manage the developer requests
package developer

import (
	"github.com/cristian-sima/Wisply/controllers/general"
)

var messages = map[string]string{
	"tableNotAllowed": "Wisply does not know this table :(",
}

// Controller represents the basic developer controller
type Controller struct {
	general.WisplyController
}

// ShowHomePage shows the static home page for developer
func (controller Controller) ShowHomePage() {
	controller.SetCustomTitle("Developers & Research")
	// Please use http://www.timestampgenerator.com/ for getting the timestamp
	controller.IndicateLastModification(1441987477)
	controller.Layout = "site/public-layout.tpl"
	controller.TplNames = "site/developer/developer.tpl"
}
