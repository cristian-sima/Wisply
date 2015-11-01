package developer

// Static shows the public HTML page for the API
type Static struct {
	Controller
}

// ShowHomePage shows the static home page for API
func (controller *Static) ShowHomePage() {
	controller.Layout = "site/public-layout.tpl"
	controller.TplNames = "site/developer/developer.tpl"
	controller.SetCustomTitle("Developers & Research")
	// Please use http://www.timestampgenerator.com/ for generating the timestamp
	controller.IndicateLastModification(1441987477)
}