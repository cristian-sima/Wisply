package public

// Static It contains the all the static pages
type Static struct {
	Controller
}

// Prepare sets the template path
func (controller *Static) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("public/static")
}

// ShowAbout shows the about page
func (controller *Static) ShowAbout() {
	controller.LoadTemplate("about")
	controller.SetCustomTitle("About Wisply")
}

// ShowContact shows the about page
func (controller *Static) ShowContact() {
	controller.LoadTemplate("contact")
	controller.SetCustomTitle("Contact Wisply")
}

// ShowTakeDownPolicy shows the take down policy
func (controller *Static) ShowTakeDownPolicy() {
	controller.LoadTemplate("take-down-policy")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1446153328)
	controller.SetCustomTitle("Take down policy")
}

// ShowThankYouPage shows the thank you page
func (controller *Static) ShowThankYouPage() {
	controller.LoadTemplate("thank-you")
	controller.SetCustomTitle("Thank you")
}

// ShowIndex shows the index page
func (controller *Static) ShowIndex() {
	controller.LoadTemplate("index")
	controller.SetCustomTitle("Wisply - Building the hive of education")
}

// ShowAccessibility shows the accessibility page
func (controller *Static) ShowAccessibility() {
	controller.LoadTemplate("accessibility")
	// Please use http://www.timestampgenerator.com/
	controller.SetCustomTitle("Accessibility")
	controller.IndicateLastModification(1441987477)
}

// ShowHelp shows the help page
func (controller *Static) ShowHelp() {
	controller.LoadTemplate("help")
	controller.SetCustomTitle("Help")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1446153828)
}

// ShowPrivacyPolicy shows the privacy policy of the website
func (controller *Static) ShowPrivacyPolicy() {
	controller.LoadTemplate("privacy")
	controller.SetCustomTitle("Wisply Privacy Policy")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1442660323)
}

// ShowTerms shows the privacy policy of the website
func (controller *Static) ShowTerms() {
	controller.LoadTemplate("terms-and-conditions")
	controller.SetCustomTitle("Wisply Terms and Conditions")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1442661323)
}

// ShowCookiesPolicy shows the policy for cookies
func (controller *Static) ShowCookiesPolicy() {
	controller.LoadTemplate("cookies")
	controller.SetCustomTitle("Wisply Cookies Policy")
	// Please use http://www.timestampgenerator.com/
	controller.IndicateLastModification(1442660323)
}

// ShowSample shows the page used by developers
func (controller *Static) ShowSample() {
	controller.LoadTemplate("site/public/static/sample.tpl")
}
