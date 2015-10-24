// Package routers contains all the addresses of the application
package routers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/account"
	"github.com/cristian-sima/Wisply/controllers/admin"
	"github.com/cristian-sima/Wisply/controllers/api"
	"github.com/cristian-sima/Wisply/controllers/public"
)

func init() {

	// ----------------------------- PUBLIC  --------------------------------------

	// Note: I can not group these into namespace because they share "/" path
	// Note: The public namespace should be created (NewNamespace)

	beego.Router("/", &public.StaticController{}, "*:ShowIndex")
	beego.Router("/about", &public.StaticController{}, "*:ShowAbout")
	beego.Router("/contact", &public.StaticController{}, "*:ShowContact")
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
	// ----------------------------- Curriculum --------------------------------------

	publicCurriculumNS := beego.NewNamespace("curriculum/",
		beego.NSNamespace(":id",
			beego.NSRouter("", &public.Curriculum{}, "GET:ShowProgram"),
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
		beego.NSNamespace("/repository/:id",
			beego.NSRouter("", &admin.RepositoryController{}, "GET:ShowRepository"),
			beego.NSNamespace("/advance-options",
				beego.NSRouter("", &admin.RepositoryController{}, "GET:ShowAdvanceOptions"),
				beego.NSNamespace("/modify",
					beego.NSRouter("", &admin.RepositoryController{}, "GET:Modify"),
					beego.NSRouter("", &admin.RepositoryController{}, "POST:Update"),
					beego.NSNamespace("/filter",
						beego.NSRouter("", &admin.RepositoryController{}, "GET:ShowFilter"),
						beego.NSRouter("", &admin.RepositoryController{}, "POST:ChangeFilter"),
					),
				),
			),
			beego.NSNamespace("/empty",
				beego.NSRouter("", &admin.RepositoryController{}, "POST:EmptyRepository"),
			),
			beego.NSNamespace("/delete",
				beego.NSRouter("", &admin.RepositoryController{}, "POST:Delete"),
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
	// ----------------------------- Curriculum ----------------------------------

	adminCurriculumNS := beego.NSNamespace("/curriculum",
		beego.NSRouter("", &admin.Curriculum{}, "*:ShowHomePage"),
		beego.NSNamespace("/programs/:id",
			beego.NSRouter("", &admin.Curriculum{}, "GET:ShowProgram"),
			beego.NSNamespace("/advance-options",
				beego.NSRouter("", &admin.Curriculum{}, "GET:ShowProgramAdvanceOptions"),
				beego.NSNamespace("/modify",
					beego.NSRouter("", &admin.Curriculum{}, "GET:ShowModifyProgramForm"),
					beego.NSRouter("", &admin.Curriculum{}, "POST:UpdateProgram"),
				),
			),
			beego.NSNamespace("/delete",
				beego.NSRouter("", &admin.Curriculum{}, "POST:DeleteProgram"),
			),
		),
		beego.NSNamespace("add",
			beego.NSRouter("", &admin.Curriculum{}, "GET:ShowAddProgramForm"),
			beego.NSRouter("", &admin.Curriculum{}, "POST:CreateProgram"),
		),
		// beego.NSNamespace("/delete",
		// 	beego.NSRouter(":id", &admin.AccountController{}, "POST:Delete"),
		// ),
	)

	// admin
	// ----------------------------- Admin API ----------------------------------

	adminAPINS := beego.NSNamespace("/api",
		beego.NSRouter("", &admin.APIController{}, "*:ShowHomePage"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &admin.APIController{}, "GET:ShowAddForm"),
			beego.NSRouter("", &admin.APIController{}, "POST:InsertNewTable"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter("", &admin.APIController{}, "POST:RemoveAllowedTable"),
		),
		beego.NSNamespace("/modify/:id",
			beego.NSRouter("", &admin.APIController{}, "GET:ShowModifyForm"),
			beego.NSRouter("", &admin.APIController{}, "POST:ModifyTable"),
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
			adminAPINS,
			adminCurriculumNS,
		)

	// api
	// ----------------------------- Repository ----------------------------------

	apiRepositoryNS := beego.NSNamespace("/repository",
		beego.NSNamespace("/resources/:id",
			beego.NSNamespace("/get",
				beego.NSRouter("/:min/:number", &api.Repository{}, "GET:GetResources"),
			),
		),
	)

	// api
	// ----------------------------- Search ----------------------------------

	apiSearchNS := beego.NSNamespace("/search",
		beego.NSNamespace("/anything/:query",
			beego.NSRouter("", &api.Search{}, "*:SearchAnything"),
		),
		beego.NSNamespace("/save/:query",
			beego.NSRouter("", &api.Search{}, "POST:JustSaveAccountQuery"),
		),
	)

	// api
	// ----------------------------- API -------------------------------

	apiNS :=
		beego.NewNamespace("/api",
			beego.NSRouter("", &api.Static{}, "GET:ShowHomePage"),
			beego.NSNamespace("/table/",
				beego.NSRouter("list", &api.Table{}, "GET:ShowList"),
				beego.NSRouter("generate/:name", &api.Table{}, "*:GenerateTable"),
				beego.NSRouter("download/:name", &api.Table{}, "*:DownloadTable"),
			),
			apiRepositoryNS,
			apiSearchNS,
		)

	// api
	// ----------------------------- ACCOUNT -------------------------------

	accountNS :=
		beego.NewNamespace("/account",
			beego.NSRouter("", &account.Home{}, "GET:Show"),
			beego.NSNamespace("/search/",
				beego.NSRouter("", &account.Search{}, "GET:DisplayHistory"),
				beego.NSRouter("clear", &account.Search{}, "POST:ClearHistory"),
			),
			beego.NSNamespace("/settings",
				beego.NSRouter("", &account.Settings{}, "GET:DisplayPage"),
				beego.NSRouter("/delete", &account.Settings{}, "POST:DeleteAccount"),
			),
		)

	// -------------------------------- REGISTER -----------------------------

	// public
	beego.AddNamespace(publicAuthNS)
	beego.AddNamespace(publicInstitutionsNS)
	beego.AddNamespace(publicRepositoryNS)
	beego.AddNamespace(publicCurriculumNS)

	// other NS
	beego.AddNamespace(adminNS)
	beego.AddNamespace(apiNS)
	beego.AddNamespace(accountNS)

}
