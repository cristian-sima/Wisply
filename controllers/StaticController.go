package controllers

// StaticController It contains the all the static pages
type StaticController struct {
	WisplyController
}

// ShowAbout It shows the about page
func (controller *StaticController) ShowAbout() {
	pageName := "about"
	controller.showStaticPage(pageName)
}

// ShowContact It shows the about page
func (controller *StaticController) ShowContact() {
	pageName := "contact"
	controller.showStaticPage(pageName)
}

// ShowIndex It shows the index page
func (controller *StaticController) ShowIndex() {
	pageName := "index"
	controller.showStaticPage(pageName)
}

// ShowWebscience It shows the webscience page
func (controller *StaticController) ShowWebscience() {
	pageName := "webscience"
	controller.showStaticPage(pageName)
}

// ShowAccessibility It shows the accessibility page
func (controller *StaticController) ShowAccessibility() {
	pageName := "accessibility"
	controller.showStaticPage(pageName)
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1441987477)
}

// ShowHelp It shows the help page
func (controller *StaticController) ShowHelp() {
	pageName := "help"
	controller.showStaticPage(pageName)
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1441987477)
}

func (controller *StaticController) showStaticPage(pageName string) {
	controller.Layout = "site/layout.tpl"
	controller.TplNames = "site/static/" + pageName + ".tpl"
}

// different

// ShowSample It shows the sample page. This contains visual elements.
// It is used by developers
func (controller *StaticController) ShowSample() {
	controller.TplNames = "site/static/sample.tpl"
}
