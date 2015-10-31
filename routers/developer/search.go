package developer

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/developer"
)

func getSearch() func(*beego.Namespace) {
	ns := beego.NSNamespace("/search",
		beego.NSNamespace("/anything/:query",
			beego.NSRouter("", &api.Search{}, "*:SearchAnything"),
		),
		beego.NSNamespace("/save/:query",
			beego.NSRouter("", &api.Search{}, "POST:JustSaveAccountQuery"),
		),
	)
	return ns
}
