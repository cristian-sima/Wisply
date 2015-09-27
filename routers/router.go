package routers

import (
	"github.com/astaxie/beego"
	admin "github.com/cristian-sima/Wisply/controllers/admin"
	public "github.com/cristian-sima/Wisply/controllers/public"
)

func init() {

	beego.Router("/", &public.StaticController{}, "*:ShowIndex")
	beego.Router("/about", &public.StaticController{}, "*:ShowAbout")
	beego.Router("/contact", &public.StaticController{}, "*:ShowContact")
	beego.Router("/webscience", &public.StaticController{}, "*:ShowWebscience")
	beego.Router("/sample", &public.StaticController{}, "*:ShowSample")
	beego.Router("/accessibility", &public.StaticController{}, "*:ShowAccessibility")

	beego.Router("/help", &public.StaticController{}, "*:ShowHelp")
	beego.Router("/privacy", &public.StaticController{}, "*:ShowPrivacyPolicy")
	beego.Router("/cookies", &public.StaticController{}, "*:ShowCookiesPolicy")
	beego.Router("/terms-and-conditions", &public.StaticController{}, "*:ShowTermsPage")

	// ----------------------------- Authentification --------------------------------------

	authNamespace := beego.NewNamespace("/auth",
		beego.NSNamespace("/login",
			beego.NSRouter("", &public.AuthController{}, "GET:ShowLoginForm"),
			beego.NSRouter("", &public.AuthController{}, "POST:LoginAccount"),
		),
		beego.NSNamespace("/register",
			beego.NSRouter("", &public.AuthController{}, "GET:ShowRegisterForm"),
			beego.NSRouter("", &public.AuthController{}, "POST:CreateNewAccount"),
		),
		beego.NSNamespace("/logout",
			beego.NSRouter("", &public.AuthController{}, "POST:Logout"),
		),
	)

	// ----------------------------- Admin --------------------------------------

	repositoryNamespace := beego.NSNamespace("/repositories",
		beego.NSRouter("", &admin.RepositoryController{}, "*:List"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &admin.RepositoryController{}, "GET:Add"),
			beego.NSRouter("", &admin.RepositoryController{}, "POST:Insert"),
		),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &admin.RepositoryController{}, "GET:Modify"),
			beego.NSRouter(":id", &admin.RepositoryController{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &admin.RepositoryController{}, "POST:Delete"),
		),
	)

	institutionsNamespace := beego.NSNamespace("/institutions",
		beego.NSRouter("", &admin.InstitutionController{}, "*:List"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &admin.InstitutionController{}, "GET:Add"),
			beego.NSRouter("", &admin.InstitutionController{}, "POST:Insert"),
		),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &admin.InstitutionController{}, "GET:Modify"),
			beego.NSRouter(":id", &admin.InstitutionController{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &admin.InstitutionController{}, "POST:Delete"),
		),
	)

	harvestNamespace := beego.NSNamespace("/harvest",
		beego.NSNamespace("/init",
			beego.NSRouter(":id", &admin.HarvestController{}, "POST:ShowPanel"),
			beego.NSRouter("/ws", &admin.HarvestController{}, "GET:InitWebsocketConnection"),
		),
	)

	accountsNamespace := beego.NSNamespace("/accounts",
		beego.NSRouter("", &admin.AccountController{}, "*:List"),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &admin.AccountController{}, "GET:Modify"),
			beego.NSRouter(":id", &admin.AccountController{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &admin.AccountController{}, "POST:Delete"),
		),
	)

	adminNamespace :=
		beego.NewNamespace("/admin",
			beego.NSRouter("", &admin.Controller{}, "*:DisplayDashboard"),
			accountsNamespace,
			repositoryNamespace,
			harvestNamespace,
			institutionsNamespace,
		)

	// register namespace
	beego.AddNamespace(authNamespace)
	beego.AddNamespace(adminNamespace)

}
