package public

// Static It contains the all the static pages
type Static struct {
	Controller
}

// ShowAbout shows the about page
func (controller *Static) ShowAbout() {
	pageName := "about"
	controller.SetCustomTitle("About Wisply")
	controller.showStaticPage(pageName)
}

// ShowContact shows the about page
func (controller *Static) ShowContact() {
	pageName := "contact"
	controller.SetCustomTitle("Contact Wisply")
	controller.showStaticPage(pageName)
}

// ShowTakeDownPolicy shows the take down policy
func (controller *Static) ShowTakeDownPolicy() {
	pageName := "take-down-policy"
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1446153328)
	controller.SetCustomTitle("Take down policy")
	controller.showStaticPage(pageName)
}

// ShowThankYouPage shows the thank you page
func (controller *Static) ShowThankYouPage() {
	pageName := "thank-you"
	controller.SetCustomTitle("Thank you")
	controller.showStaticPage(pageName)
}

// ShowIndex shows the index page
func (controller *Static) ShowIndex() {
	pageName := "index"
	controller.SetCustomTitle("Wisply - Building the hive of education")
	controller.showStaticPage(pageName)
}

// ShowAccessibility shows the accessibility page
func (controller *Static) ShowAccessibility() {
	pageName := "accessibility"
	controller.showStaticPage(pageName)
	// Please use http://www.timestampgenerator.com/
	controller.SetCustomTitle("Accessibility")
	controller.IndicateLastModification(1441987477)
}

// ShowHelp shows the help page
func (controller *Static) ShowHelp() {
	pageName := "help"
	controller.showStaticPage(pageName)
	controller.SetCustomTitle("Help")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1446153828)
}

// ShowPrivacyPolicy shows the privacy policy of the website
func (controller *Static) ShowPrivacyPolicy() {
	pageName := "privacy"
	controller.showStaticPage(pageName)
	controller.SetCustomTitle("Wisply Privacy Policy")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1442660323)
}

// ShowTerms shows the privacy policy of the website
func (controller *Static) ShowTerms() {
	pageName := "terms-and-conditions"
	controller.showStaticPage(pageName)
	controller.SetCustomTitle("Wisply Terms and Conditions")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1442661323)
}

// ShowCookiesPolicy shows the policy for cookies
func (controller *Static) ShowCookiesPolicy() {
	pageName := "cookies"
	controller.showStaticPage(pageName)
	controller.SetCustomTitle("Wisply Cookies Policy")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1442660323)
}

func (controller *Static) showStaticPage(pageName string) {
	controller.TplNames = "site/public/static/" + pageName + ".tpl"
}

// ShowSample shows the sample page. This page contains visual elements.
// It is used by developers
func (controller *Static) ShowSample() {
	controller.TplNames = "site/public/static/sample.tpl"
}
