package education

import (
	"github.com/astaxie/beego"

	program "github.com/cristian-sima/Wisply/controllers/admin/education/program"
)

func getProgram() func(*beego.Namespace) {
	ns := beego.NSNamespace("/programs",
		beego.NSNamespace("/:program",
			beego.NSRouter("", &program.Program{}, "*:Display"),
			beego.NSNamespace("/advance-options",
				beego.NSRouter("", &program.Program{}, "*:ShowAdvanceOptions"),
				beego.NSNamespace("/modify",
					beego.NSRouter("", &program.Program{}, "GET:ShowModifyForm"),
					beego.NSRouter("", &program.Program{}, "POST:Update"),
				),
				beego.NSNamespace("/modify-static-description",
					beego.NSRouter("", &program.Program{}, "GET:ShowModifyForm"),
					beego.NSRouter("", &program.Program{}, "POST:UpdateDescription"),
				),
			),
			beego.NSNamespace("/delete",
				beego.NSRouter("", &program.Program{}, "POST:Delete"),
			),
			beego.NSNamespace("/definition",
				beego.NSNamespace("/add",
					beego.NSRouter("", &program.Definition{}, "GET:ShowAddForm"),
					beego.NSRouter("", &program.Definition{}, "POST:CreateDefinition"),
				),
				beego.NSNamespace("/:definition",
					beego.NSNamespace("/modify",
						beego.NSRouter("", &program.Definition{}, "GET:ShowModifyForm"),
						beego.NSRouter("", &program.Definition{}, "POST:Update"),
					),
					beego.NSNamespace("/delete",
						beego.NSRouter("", &program.Definition{}, "POST:Delete"),
					),
				),
			),
		),
		beego.NSNamespace("add",
			beego.NSRouter("", &program.Program{}, "GET:ShowAddForm"),
			beego.NSRouter("", &program.Program{}, "POST:CreateProgram"),
		),
	)
	return ns
}
