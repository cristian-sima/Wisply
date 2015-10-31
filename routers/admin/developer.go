package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

func getDeveloper() func(*beego.Namespace) {
	ns := beego.NSNamespace("/api",
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
	return ns
}
