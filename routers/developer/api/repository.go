package api

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/developer/api"
)

func getRepository() func(*beego.Namespace) {
	ns := beego.NSNamespace("/repository",
		beego.NSNamespace("/resources/:id",
			beego.NSNamespace("/get",
				beego.NSRouter("/:min/:number", &api.Repository{}, "GET:GetResources"),
			),
		),
	)
	return ns
}
