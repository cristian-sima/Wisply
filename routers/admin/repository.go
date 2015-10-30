package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

func getRepository() func(*beego.Namespace) {
	ns := beego.NSNamespace("/repositories",
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
	return ns
}
