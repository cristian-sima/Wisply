package institutions

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin/institutions"
	"github.com/cristian-sima/Wisply/controllers/admin/institutions/institution"
)

// Get returns the namespace for institutions
func Get() func(*beego.Namespace) {
	ns := beego.NSNamespace("/institutions",
		beego.NSRouter("", &institutions.Home{}, "*:Display"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &institution.Institution{}, "GET:ShowInsertForm"),
			beego.NSRouter("", &institution.Institution{}, "POST:Insert"),
		),
		beego.NSNamespace("/:institution",
			beego.NSRouter("/analyse", &institution.Institution{}, "GET:Analyse"),
			beego.NSRouter("", &institution.Institution{}, "GET:Display"),
			beego.NSRouter("/advance-options", &institution.Institution{}, "GET:ShowAdvanceOptions"),
			beego.NSNamespace("/modify",
				beego.NSRouter("", &institution.Institution{}, "GET:ShowModifyForm"),
				beego.NSRouter("", &institution.Institution{}, "POST:Modify"),
			),
			beego.NSNamespace("/delete",
				beego.NSRouter("", &institution.Institution{}, "POST:Delete"),
			),
			// program
			beego.NSNamespace("/program",
				beego.NSNamespace("/add",
					beego.NSRouter("", &institution.Program{}, "GET:ShowInsertForm"),
					beego.NSRouter("", &institution.Program{}, "POST:CreateProgram"),
				),
				beego.NSNamespace("/:program",
					beego.NSRouter("", &institution.Program{}, "GET:Display"),
					beego.NSNamespace("/modify",
						beego.NSRouter("", &institution.Program{}, "GET:ShowModifyForm"),
						beego.NSRouter("", &institution.Program{}, "POST:Modify"),
					),
					beego.NSRouter("/delete", &institution.Program{}, "POST:Delete"),

					// module within program of study
					beego.NSNamespace("/module",
						beego.NSNamespace("/add",
							beego.NSRouter("", &institution.Program{}, "GET:ShowAddModuleForm"),
							beego.NSRouter("", &institution.Program{}, "POST:AddModule"),
						),
						beego.NSRouter("/:module/delete", &institution.Program{}, "POST:DeleteModule"),
					),
				),
			),
			// module within institution
			beego.NSNamespace("/module",
				beego.NSNamespace("/add",
					beego.NSRouter("", &institution.Module{}, "GET:ShowInsertForm"),
					beego.NSRouter("", &institution.Module{}, "POST:CreateModule"),
				),
				beego.NSNamespace("/:module",
					beego.NSNamespace("/modify",
						beego.NSRouter("", &institution.Module{}, "GET:ShowModifyForm"),
						beego.NSRouter("", &institution.Module{}, "POST:Modify"),
					),
					beego.NSRouter("/delete", &institution.Module{}, "POST:Delete"),
				),
			),
		),
	)
	return ns
}
