package controllers

type StaticController struct {
	WisplyController
}

func (controller *StaticController) ShowAbout() {
	pageName := "about"
	controller.showStaticPage(pageName)
}

func (controller *StaticController) ShowContact() {
	pageName := "contact"
	controller.showStaticPage(pageName)
}

func (controller *StaticController) ShowIndex() {
	pageName := "index"
	controller.showStaticPage(pageName)
}

func (controller *StaticController) ShowWebscience() {
	pageName := "webscience"
	controller.showStaticPage(pageName)
}

func (controller *StaticController) ShowAccessibility() {
	pageName := "accessibility"
	controller.showStaticPage(pageName)
}

func (controller *StaticController) ShowHelp() {
	pageName := "help"
	controller.showStaticPage(pageName)
}

func (controller *StaticController) showStaticPage(pageName string) {
	controller.Layout = "site/layout.tpl"
	controller.TplNames = "site/static/" + pageName + ".tpl"
}

// different

func (controller *StaticController) ShowSample() {
	controller.TplNames = "site/static/sample.tpl"
}
