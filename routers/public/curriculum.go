package public

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/public"
)

func getCurriculum() *beego.Namespace {
	ns := beego.NewNamespace("curriculum",
		beego.NSRouter("", &public.Curriculum{}, "*:ShowCurricula"),
		beego.NSNamespace("/:id",
			beego.NSRouter("", &public.Curriculum{}, "GET:ShowProgram"),
		),
	)
	return ns
}
