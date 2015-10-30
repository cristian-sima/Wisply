package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

func getCurriculum() func(*beego.Namespace) {
	ns := beego.NSNamespace("/curriculum",
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
	)
	return ns
}
