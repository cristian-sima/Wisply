package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

func getInstitution() func(*beego.Namespace) {
	ns := beego.NSNamespace("/institutions",
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
	return ns
}
