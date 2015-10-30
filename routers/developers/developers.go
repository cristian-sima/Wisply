package developers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/api"
)

func Load() {

	// api
	// ----------------------------- Repository ----------------------------------

	apiRepositoryNS := beego.NSNamespace("/repository",
		beego.NSNamespace("/resources/:id",
			beego.NSNamespace("/get",
				beego.NSRouter("/:min/:number", &api.Repository{}, "GET:GetResources"),
			),
		),
	)

	// api
	// ----------------------------- Search ----------------------------------

	apiSearchNS := beego.NSNamespace("/search",
		beego.NSNamespace("/anything/:query",
			beego.NSRouter("", &api.Search{}, "*:SearchAnything"),
		),
		beego.NSNamespace("/save/:query",
			beego.NSRouter("", &api.Search{}, "POST:JustSaveAccountQuery"),
		),
	)

	// api
	// ----------------------------- API -------------------------------

	// api
	// ----------------------------- ACCOUNT -------------------------------

	apiNS :=
		beego.NewNamespace("/api",
			beego.NSRouter("", &api.Static{}, "GET:ShowHomePage"),
			beego.NSNamespace("/table/",
				beego.NSRouter("list", &api.Table{}, "GET:ShowList"),
				beego.NSRouter("generate/:name", &api.Table{}, "*:GenerateTable"),
				beego.NSRouter("download/:name", &api.Table{}, "*:DownloadTable"),
			),
			apiRepositoryNS,
			apiSearchNS,
		)

	beego.AddNamespace(apiNS)
}
