package routers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers"
)

func init() {

	beego.Router("/", &controllers.StaticController{}, "*:ShowIndex")
	beego.Router("/about", &controllers.StaticController{}, "*:ShowAbout")
	beego.Router("/contact", &controllers.StaticController{}, "*:ShowContact")
	beego.Router("/webscience", &controllers.StaticController{}, "*:ShowWebscience")
	beego.Router("/sample", &controllers.StaticController{}, "*:ShowSample")
	beego.Router("/accessibility", &controllers.StaticController{}, "*:ShowAccessibility")

	beego.Router("/help", &controllers.StaticController{}, "*:ShowHelp")
	beego.Router("/privacy", &controllers.StaticController{}, "*:ShowPrivacyPolicy")
	beego.Router("/cookies", &controllers.StaticController{}, "*:ShowCookiesPolicy")
	beego.Router("/terms-and-conditions", &controllers.StaticController{}, "*:ShowTermsPage")

	// ----------------------------- Authentification --------------------------------------

	authNamespace := beego.NewNamespace("/auth",
		beego.NSNamespace("/login",
			beego.NSRouter("", &controllers.AuthController{}, "GET:ShowLoginForm"),
			beego.NSRouter("", &controllers.AuthController{}, "POST:LoginAccount"),
		),
		beego.NSNamespace("/register",
			beego.NSRouter("", &controllers.AuthController{}, "GET:ShowRegisterForm"),
			beego.NSRouter("", &controllers.AuthController{}, "POST:CreateNewAccount"),
		),
		beego.NSNamespace("/logout",
			beego.NSRouter("", &controllers.AuthController{}, "POST:Logout"),
		),
	)

	// ----------------------------- Admin --------------------------------------

	sourcesNamespace := beego.NSNamespace("/sources",
		beego.NSRouter("", &controllers.SourceController{}, "*:ListSources"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &controllers.SourceController{}, "GET:AddNewSource"),
			beego.NSRouter("", &controllers.SourceController{}, "POST:InsertSource"),
		),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &controllers.SourceController{}, "GET:Modify"),
			beego.NSRouter(":id", &controllers.SourceController{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &controllers.SourceController{}, "POST:Delete"),
		),
	)

	accountsNamespace := beego.NSNamespace("/accounts",
		beego.NSRouter("", &controllers.AccountController{}, "*:ListAccounts"),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &controllers.AccountController{}, "GET:Modify"),
			beego.NSRouter(":id", &controllers.AccountController{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &controllers.AccountController{}, "POST:Delete"),
		),
	)

	adminNamespace :=
		beego.NewNamespace("/admin",
			beego.NSRouter("", &controllers.AdminController{}, "*:DisplayDashboard"),
			sourcesNamespace,
			accountsNamespace,
		)

	// register namespace
	beego.AddNamespace(authNamespace)
	beego.AddNamespace(adminNamespace)

}
