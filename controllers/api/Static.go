package api

// Controller
type Static struct {
	Controller
}

// ShowHomePage shows the static home page for API
func (controller *Static) ShowHomePage() {
	controller.Layout = "site/public-layout.tpl"
	controller.TplNames = "site/api/api.tpl"
	// Please use http://www.timestampgenerator.com/
	controller.SetCustomTitle("API & Developers")
	controller.IndicateLastModification(1441987477)
}
