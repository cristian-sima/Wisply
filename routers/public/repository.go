package public

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/public"
)

func getRepository() *beego.Namespace {
	ns := beego.NewNamespace("/repository",
		beego.NSNamespace("/:repository",
			beego.NSRouter("", &public.Repository{}, "GET:ShowRepository"),
			beego.NSNamespace("/resource",
				beego.NSNamespace("/:resource",
					beego.NSRouter("", &public.Repository{}, "GET:ShowResource"),
					beego.NSRouter("/content", &public.Repository{}, "GET:GetResourceContent"),
				),
			),
		),
	)
	return ns
}
