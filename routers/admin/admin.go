package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

// Load tells the framework to load the addresses for the router
func Load() {

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

	adminNS := beego.NewNamespace("/admin",
		beego.NSRouter("", &admin.Home{}, "*:DisplayDashboard"),
		adminAccountsNS,
		adminRepositoryNS,
		adminInstitutionsNS,
		adminHarvestNS,
		adminLogNS,
		adminAPINS,
		adminCurriculumNS,
	)

	beego.AddNamespace(adminNS)
}
