package institutions

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/public/institutions"
	"github.com/cristian-sima/Wisply/controllers/public/institutions/institution"
)

// Get returns the namespace for institutions
func Get() *beego.Namespace {
	ns := beego.NewNamespace("/institutions",
		beego.NSRouter("", &institutions.Home{}, "*:Display"),
		beego.NSNamespace("/:institution",
			beego.NSRouter("", &institution.Institution{}, "GET:Display"),
			// program
			beego.NSNamespace("/program",
				beego.NSNamespace("/:program",
					beego.NSRouter("", &institution.Program{}, "GET:Display"),
					// module
					beego.NSNamespace("/module",
						beego.NSNamespace("/:module",
							beego.NSRouter("", &institution.Module{}, "GET:Display"),
						),
					),
				),
			),
		),
	)
	return ns
}
