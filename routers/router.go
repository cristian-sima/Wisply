package routers

import (
	"github.com/astaxie/beego"
	admin "github.com/cristian-sima/Wisply/controllers/admin"
	public "github.com/cristian-sima/Wisply/controllers/public"
)

func init() {

	// ----------------------------- PUBLIC  --------------------------------------

	// Note: I can not group these into namespace because they share "/" path
	// Note: The public namespace should be created (NewNamespace)

	beego.Router("/", &public.StaticController{}, "*:ShowIndex")
	beego.Router("/about", &public.StaticController{}, "*:ShowAbout")
	beego.Router("/contact", &public.StaticController{}, "*:ShowContact")
	beego.Router("/webscience", &public.StaticController{}, "*:ShowWebscience")
	beego.Router("/sample", &public.StaticController{}, "*:ShowSample")
	beego.Router("/accessibility", &public.StaticController{}, "*:ShowAccessibility")
	beego.Router("/help", &public.StaticController{}, "*:ShowHelp")
	beego.Router("/privacy", &public.StaticController{}, "*:ShowPrivacyPolicy")
	beego.Router("/cookies", &public.StaticController{}, "*:ShowCookiesPolicy")
	beego.Router("/terms-and-conditions", &public.StaticController{}, "*:ShowTerms")

	// public
	// ----------------------------- Authentification --------------------------------------

	publicAuthNS := beego.NewNamespace("auth",
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

	// public
	// ----------------------------- Institutions -------------------------------

	publicInstitutionsNS := beego.NewNamespace("/institutions",
		beego.NSRouter("", &public.InstitutionController{}, "*:List"),
		beego.NSRouter("/:id", &public.InstitutionController{}, "GET:ShowInstitution"),
	)

	// public
	// ----------------------------- Repositories -------------------------------

	publicRepositoryNS := beego.NewNamespace("/repository",
		beego.NSRouter("/:id", &public.RepositoryController{}, "GET:ShowRepository"),
	)

	// ----------------------------- ADMIN --------------------------------------

	// admin
	// ----------------------------- Repositories -------------------------------

	adminRepositoryNS := beego.NSNamespace("/repositories",
		beego.NSRouter("", &admin.RepositoryController{}, "*:List"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &admin.RepositoryController{}, "GET:ShowTypes"),
		),
		beego.NSNamespace("/newRepository",
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
		beego.NSNamespace("/repository",
			beego.NSNamespace("/:id",
				beego.NSRouter("", &admin.RepositoryController{}, "GET:ShowRepository"),
				beego.NSRouter("/advance-options", &admin.RepositoryController{}, "GET:ShowAdvanceOptions"),
			),
		),
	)

	// admin
	// ----------------------------- Institutions -------------------------------

	adminInstitutionsNS := beego.NSNamespace("/institutions",
		beego.NSRouter("", &admin.InstitutionController{}, "*:DisplayAll"),
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
		beego.NSNamespace("/institution",
			beego.NSNamespace("/:id",
				beego.NSRouter("", &admin.RepositoryController{}, "GET:ShowInstitution"),
				beego.NSRouter("/advance-options", &admin.RepositoryController{}, "GET:ShowAdvanceInstitutionOptions"),
			),
		),
	)

	// admin
	// ----------------------------- Harvest -----------------------------------

	adminHarvestNS := beego.NSNamespace("/harvest",
		beego.NSNamespace("/init",
			beego.NSRouter(":id", &admin.HarvestController{}, "POST:ShowPanel"),
			beego.NSRouter("/ws", &admin.HarvestController{}, "GET:InitWebsocketConnection"),
		),
		beego.NSNamespace("/recover",
			beego.NSRouter(":id", &admin.HarvestController{}, "POST:RecoverProcess"),
		),
		beego.NSNamespace("/finish",
			beego.NSRouter(":id", &admin.HarvestController{}, "POST:ForceFinishProcess"),
		),
	)

	// admin
	// ----------------------------- Log -----------------------------------

	adminLogNS := beego.NSNamespace("/log",
		beego.NSRouter("", &admin.LogController{}, "*:ShowGeneralPage"),
		beego.NSNamespace("/process",
			beego.NSRouter(":process", &admin.LogController{}, "*:ShowProcess"),
			beego.NSNamespace(":process/operation",
				beego.NSRouter(":operation", &admin.LogController{}, "*:ShowOperation"),
			),
			beego.NSRouter(":process/history", &admin.LogController{}, "*:ShowProgressHistory"),
			beego.NSRouter(":process/advance-options", &admin.LogController{}, "*:ShowProcessAdvanceOptions"),
			beego.NSRouter("/delete/:process", &admin.LogController{}, "POST:DeleteProcess"),
		),
		beego.NSRouter("/advance-options", &admin.LogController{}, "*:ShowLogAdvanceOptions"),
		beego.NSRouter("/delete", &admin.LogController{}, "POST:DeleteEntireLog"),
	)

	// admin
	// ----------------------------- Accounts ----------------------------------

	adminAccountsNS := beego.NSNamespace("/accounts",
		beego.NSRouter("", &admin.AccountController{}, "*:List"),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &admin.AccountController{}, "GET:Modify"),
			beego.NSRouter(":id", &admin.AccountController{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &admin.AccountController{}, "POST:Delete"),
		),
	)

	// admin
	// ----------------------------- Admin -------------------------------

	adminNS :=
		beego.NewNamespace("/admin",
			beego.NSRouter("", &admin.Controller{}, "*:DisplayDashboard"),
			adminAccountsNS,
			adminRepositoryNS,
			adminInstitutionsNS,
			adminHarvestNS,
			adminLogNS,
		)

	// -------------------------------- REGISTER -----------------------------

	// public
	beego.AddNamespace(publicAuthNS)
	beego.AddNamespace(publicInstitutionsNS)
	beego.AddNamespace(publicRepositoryNS)

	// admin
	beego.AddNamespace(adminNS)

}
