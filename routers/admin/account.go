package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

func getAccount() func(*beego.Namespace) {
	ns := beego.NSNamespace("/accounts",
		beego.NSRouter("", &admin.Account{}, "*:List"),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &admin.Account{}, "GET:Modify"),
			beego.NSRouter(":id", &admin.Account{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &admin.Account{}, "POST:Delete"),
		),
	)
	return ns
}
