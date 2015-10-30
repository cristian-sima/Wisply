package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

func getLog() func(*beego.Namespace) {
	ns := beego.NSNamespace("/log",
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
	return ns
}
