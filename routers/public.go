package routers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/general"
	"github.com/cristian-sima/Wisply/controllers/public"
)

func loadPublic() {
	// ----------------------------- PUBLIC  --------------------------------------

	// Note: I can not group these into namespace because they share "/" path
	// Note: The public namespace should be created (NewNamespace)

	beego.Router("/", &public.Static{}, "*:ShowIndex")
	beego.Router("/about", &public.Static{}, "*:ShowAbout")
	beego.Router("/learn-more", &public.Static{}, "*:ShowAbout")
	beego.Router("/contact", &public.Static{}, "*:ShowContact")
	beego.Router("/sample", &public.Static{}, "*:ShowSample")
	beego.Router("/accessibility", &public.Static{}, "*:ShowAccessibility")

	beego.Router("/help", &public.Static{}, "*:ShowHelp")

	beego.Router("/privacy", &public.Static{}, "*:ShowPrivacyPolicy")
	beego.Router("/cookies", &public.Static{}, "*:ShowCookiesPolicy")
	beego.Router("/terms-and-conditions", &public.Static{}, "*:ShowTerms")
	beego.Router("/take-down-policy", &public.Static{}, "*:ShowTakeDownPolicy")

	beego.Router("/thank-you", &public.Static{}, "*:ShowThankYouPage")

	beego.Router("/curricula", &public.Curriculum{}, "*:ShowCurricula")

	beego.Router("/test", &public.Analyse{}, "*:AnalyseText")

	beego.Router("/captcha/:id\\.:type", &general.Captcha{}, "*:Serve")

	// public
	// ----------------------------- Authentification --------------------------------------

	publicAuthNS := beego.NewNamespace("auth",
		beego.NSNamespace("/login",
			beego.NSRouter("", &public.Auth{}, "GET:ShowLoginPage"),
			beego.NSRouter("", &public.Auth{}, "POST:LoginAccount"),
		),
		beego.NSNamespace("/register",
			beego.NSRouter("", &public.Auth{}, "GET:ShowRegisterForm"),
			beego.NSRouter("", &public.Auth{}, "POST:CreateNewAccount"),
		),
		beego.NSNamespace("/logout",
			beego.NSRouter("", &public.Auth{}, "POST:Logout"),
		),
	)

	// public
	// ----------------------------- Curriculum --------------------------------------

	publicCurriculumNS := beego.NewNamespace("curriculum/",
		beego.NSNamespace(":id",
			beego.NSRouter("", &public.Curriculum{}, "GET:ShowProgram"),
		),
	)

	// public
	// ----------------------------- Institutions -------------------------------

	publicInstitutionsNS := beego.NewNamespace("/institutions",
		beego.NSRouter("", &public.Institution{}, "*:List"),
		beego.NSRouter("/:id", &public.Institution{}, "GET:ShowInstitution"),
	)

	// public
	// ----------------------------- Repositories -------------------------------

	publicRepositoryNS := beego.NewNamespace("/repository",
		beego.NSNamespace("/:repository",
			beego.NSRouter("", &public.Repository{}, "GET:ShowRepository"),
			beego.NSNamespace("/resource",
				beego.NSNamespace("/:resource",
					beego.NSRouter("", &public.Repository{}, "GET:ShowResource"),
					beego.NSRouter("/content", &public.Repository{}, "GET:GetResourceContent"),
				),
			),
		),
	)

	beego.AddNamespace(publicAuthNS)
	beego.AddNamespace(publicInstitutionsNS)
	beego.AddNamespace(publicRepositoryNS)
	beego.AddNamespace(publicCurriculumNS)

}
