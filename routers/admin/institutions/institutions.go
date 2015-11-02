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
			beego.NSRouter("", &institution.Institution{}, "GET:Display"),
			beego.NSRouter("/advance-options", &institution.Institution{}, "GET:ShowAdvanceOptions"),
			beego.NSNamespace("/modify",
				beego.NSRouter("", &institution.Institution{}, "GET:ShowModifyForm"),
				beego.NSRouter("", &institution.Institution{}, "POST:Modify"),
			),
			beego.NSNamespace("/delete",
				beego.NSRouter("", &institution.Institution{}, "POST:Delete"),
			),
		),
	)
	return ns
}
