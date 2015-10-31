package public

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/public"
)

func getInstitution() *beego.Namespace {
	ns := beego.NewNamespace("/institutions",
		beego.NSRouter("", &public.Institution{}, "*:List"),
		beego.NSRouter("/:id", &public.Institution{}, "GET:ShowInstitution"),
	)
	return ns
}
