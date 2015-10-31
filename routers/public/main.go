// Package public contains all the pages which can be accessed by anyone
// without a connection
package public

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/general"
	"github.com/cristian-sima/Wisply/controllers/public"
)

// Load tells the framework to load the addresses for the router
func Load() {
	loadNS()
	loadRoot()
}

func loadNS() {

	auth := getAuth()
	curriculum := getCurriculum()
	institution := getInstitution()
	repository := getRepository()

	beego.AddNamespace(auth)
	beego.AddNamespace(curriculum)
	beego.AddNamespace(institution)
	beego.AddNamespace(repository)
}

func loadRoot() {
	loadDefault()
	loadHelp()
	loadCaptcha()
	loadOthers()
}

func loadDefault() {
	beego.Router("/", &public.Static{}, "*:ShowIndex")
	beego.Router("/about", &public.Static{}, "*:ShowAbout")
	beego.Router("/contact", &public.Static{}, "*:ShowContact")
	beego.Router("/sample", &public.Static{}, "*:ShowSample")
	beego.Router("/accessibility", &public.Static{}, "*:ShowAccessibility")
	beego.Router("/thank-you", &public.Static{}, "*:ShowThankYouPage")
}

func loadHelp() {
	beego.Router("/help", &public.Static{}, "*:ShowHelp")
	beego.Router("/privacy", &public.Static{}, "*:ShowPrivacyPolicy")
	beego.Router("/cookies", &public.Static{}, "*:ShowCookiesPolicy")
	beego.Router("/terms-and-conditions", &public.Static{}, "*:ShowTerms")
	beego.Router("/take-down-policy", &public.Static{}, "*:ShowTakeDownPolicy")
}

func loadCaptcha() {
	beego.Router("/captcha/:id\\.:type", &general.Captcha{}, "*:Serve")
}

func loadOthers() {
	beego.Router("/test", &public.Analyse{}, "*:AnalyseText")
}
