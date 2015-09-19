package controllers

// StaticController It contains the all the static pages
type StaticController struct {
	WisplyController
}

// ShowAbout It shows the about page
func (controller *StaticController) ShowAbout() {
	pageName := "about"
	controller.SetCustomTitle("About Wisply")
	controller.showStaticPage(pageName)
}

// ShowContact It shows the about page
func (controller *StaticController) ShowContact() {
	pageName := "contact"
	controller.SetCustomTitle("Contact Wisply")
	controller.showStaticPage(pageName)
}

// ShowIndex It shows the index page
func (controller *StaticController) ShowIndex() {
	pageName := "index"
	controller.SetCustomTitle("Wisply - Building the hive of education")
	controller.showStaticPage(pageName)
}

// ShowWebscience It shows the webscience page
func (controller *StaticController) ShowWebscience() {
	pageName := "webscience"
	controller.SetCustomTitle("Webscience")
	controller.showStaticPage(pageName)
}

// ShowAccessibility It shows the accessibility page
func (controller *StaticController) ShowAccessibility() {
	pageName := "accessibility"
	controller.showStaticPage(pageName)
	// Please use http://www.timestampgenerator.com/
	controller.SetCustomTitle("Accessibility")
	controller.IndicateLastModification(1441987477)
}

// ShowHelp It shows the help page
func (controller *StaticController) ShowHelp() {
	pageName := "help"
	controller.showStaticPage(pageName)
	controller.SetCustomTitle("Help")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1441987477)
}

// legal aspects

// ShowPrivacyPolicy shows the privacy policy of the website
func (controller *StaticController) ShowPrivacyPolicy() {
	pageName := "privacy"
	controller.showStaticPage(pageName)

	controller.SetCustomTitle("Wisply Privicy Policy")
	controller.IndicateLastModification(1442660323)
}

// ShowTermsPage shows the privacy policy of the website
func (controller *StaticController) ShowTermsPage() {
	pageName := "terms-and-conditions"
	controller.showStaticPage(pageName)

	controller.SetCustomTitle("Wisply Terms and Conditions")
	controller.IndicateLastModification(1442661323)
}

// ShowCookiesPolicy shows the policy for cookies
func (controller *StaticController) ShowCookiesPolicy() {
	pageName := "cookies"
	controller.showStaticPage(pageName)

	controller.SetCustomTitle("Wisply Cookie Policy")
	controller.IndicateLastModification(1442660323)
}

// --- end

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
