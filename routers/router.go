// Package routers contains all the addresses of the application
package routers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/account"
	"github.com/cristian-sima/Wisply/controllers/admin"
	"github.com/cristian-sima/Wisply/controllers/api"
	"github.com/cristian-sima/Wisply/controllers/general"
	"github.com/cristian-sima/Wisply/controllers/public"
)

func init() {

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

	beego.Router("/test", &admin.Analyse{}, "*:AnalyseText")

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

	// ----------------------------- ADMIN --------------------------------------

	// admin
	// ----------------------------- Repositories -------------------------------

	adminRepositoryNS := beego.NSNamespace("/repositories",
		beego.NSRouter("", &admin.Repository{}, "*:List"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &admin.Repository{}, "GET:ShowTypes"),
		),
		beego.NSNamespace("/newRepository",
			beego.NSRouter("", &admin.Repository{}, "GET:Add"),
			beego.NSRouter("", &admin.Repository{}, "POST:Insert"),
		),
		beego.NSNamespace("/repository/:id",
			beego.NSRouter("", &admin.Repository{}, "GET:ShowRepository"),
			beego.NSNamespace("/advance-options",
				beego.NSRouter("", &admin.Repository{}, "GET:ShowAdvanceOptions"),
				beego.NSNamespace("/modify",
					beego.NSRouter("", &admin.Repository{}, "GET:Modify"),
					beego.NSRouter("", &admin.Repository{}, "POST:Update"),
					beego.NSNamespace("/filter",
						beego.NSRouter("", &admin.Repository{}, "GET:ShowFilter"),
						beego.NSRouter("", &admin.Repository{}, "POST:ChangeFilter"),
					),
				),
			),
			beego.NSNamespace("/empty",
				beego.NSRouter("", &admin.Repository{}, "POST:EmptyRepository"),
			),
			beego.NSNamespace("/delete",
				beego.NSRouter("", &admin.Repository{}, "POST:Delete"),
			),
		),
	)

	// admin
	// ----------------------------- Institutions -------------------------------

	adminInstitutionsNS := beego.NSNamespace("/institutions",
		beego.NSRouter("", &admin.Institution{}, "*:DisplayAll"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &admin.Institution{}, "GET:Add"),
			beego.NSRouter("", &admin.Institution{}, "POST:Insert"),
		),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &admin.Institution{}, "GET:Modify"),
			beego.NSRouter(":id", &admin.Institution{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &admin.Institution{}, "POST:Delete"),
		),
		beego.NSNamespace("/institution",
			beego.NSNamespace("/:id",
				beego.NSRouter("", &admin.Repository{}, "GET:ShowInstitution"),
				beego.NSRouter("/advance-options", &admin.Repository{}, "GET:ShowAdvanceInstitutionOptions"),
			),
		),
	)

	// admin
	// ----------------------------- Harvest -----------------------------------

	adminHarvestNS := beego.NSNamespace("/harvest",
		beego.NSNamespace("/init",
			beego.NSRouter(":id", &admin.Harvest{}, "POST:ShowPanel"),
			beego.NSRouter("/ws", &admin.Harvest{}, "GET:InitWebsocketConnection"),
		),
		beego.NSNamespace("/recover",
			beego.NSRouter(":id", &admin.Harvest{}, "POST:RecoverProcess"),
		),
		beego.NSNamespace("/finish",
			beego.NSRouter(":id", &admin.Harvest{}, "POST:ForceFinishProcess"),
		),
	)

	// admin
	// ----------------------------- Log -----------------------------------

	adminLogNS := beego.NSNamespace("/log",
		beego.NSRouter("", &admin.Log{}, "*:ShowGeneralPage"),
		beego.NSNamespace("/process",
			beego.NSRouter(":process", &admin.Log{}, "*:ShowProcess"),
			beego.NSNamespace(":process/operation",
				beego.NSRouter(":operation", &admin.Log{}, "*:ShowOperation"),
			),
			beego.NSRouter(":process/history", &admin.Log{}, "*:ShowProgressHistory"),
			beego.NSRouter(":process/advance-options", &admin.Log{}, "*:ShowProcessAdvanceOptions"),
			beego.NSRouter("/delete/:process", &admin.Log{}, "POST:DeleteProcess"),
		),
		beego.NSRouter("/advance-options", &admin.Log{}, "*:ShowLogAdvanceOptions"),
		beego.NSRouter("/delete", &admin.Log{}, "POST:DeleteEntireLog"),
	)

	// admin
	// ----------------------------- Accounts ----------------------------------

	adminAccountsNS := beego.NSNamespace("/accounts",
		beego.NSRouter("", &admin.Account{}, "*:List"),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &admin.Account{}, "GET:Modify"),
			beego.NSRouter(":id", &admin.Account{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &admin.Account{}, "POST:Delete"),
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
				beego.NSNamespace("/modify-static-description",
					beego.NSRouter("", &admin.Curriculum{}, "GET:ShowModifyDescriptionForm"),
					beego.NSRouter("", &admin.Curriculum{}, "POST:UpdateProgramDescription"),
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
		// 	beego.NSRouter(":id", &admin.Account{}, "POST:Delete"),
		// ),
	)

	// admin
	// ----------------------------- Admin API ----------------------------------

	adminAPINS := beego.NSNamespace("/api",
		beego.NSRouter("", &admin.Developers{}, "*:ShowHomePage"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &admin.Developers{}, "GET:ShowAddForm"),
			beego.NSRouter("", &admin.Developers{}, "POST:InsertNewTable"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter("", &admin.Developers{}, "POST:RemoveAllowedTable"),
		),
		beego.NSNamespace("/modify/:id",
			beego.NSRouter("", &admin.Developers{}, "GET:ShowModifyForm"),
			beego.NSRouter("", &admin.Developers{}, "POST:ModifyTable"),
		),
	)

	// admin
	// ----------------------------- Admin -------------------------------

	adminNS :=
		beego.NewNamespace("/admin",
			beego.NSRouter("", &admin.Home{}, "*:DisplayDashboard"),
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
