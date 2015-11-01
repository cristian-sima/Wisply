package curriculum

import (
	"github.com/astaxie/beego"

	program "github.com/cristian-sima/Wisply/controllers/admin/education/program"
)

func getProgram() func(*beego.Namespace) {
	ns := beego.NSNamespace("/program/:id",
		beego.NSRouter("", &program.Program{}, "*:Display"),
		beego.NSNamespace("/advance-options",
			beego.NSRouter("", &program.Program{}, "*:ShowAdvanceOptions"),
			beego.NSNamespace("/modify",
				beego.NSRouter("", &program.Program{}, "*:ShowModifyForm"),
				beego.NSRouter("", &program.Program{}, "POST:Update"),
			),
			beego.NSNamespace("/modify-static-description",
				beego.NSRouter("", &program.Program{}, "*:ShowModifyForm"),
				beego.NSRouter("", &program.Program{}, "POST:UpdateDescription"),
			),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter("", &program.Program{}, "POST:DeleteProgram"),
		),
		beego.NSNamespace("add",
			beego.NSRouter("", &program.Program{}, "*:ShowAddProgramForm"),
			beego.NSRouter("", &program.Program{}, "POST:CreateProgram"),
		),
	)
	return ns
}
