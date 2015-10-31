// Package developer contains all the addresses for the developers area
package developer

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/developer"
)

// Load tells the framework to load the addresses for the router
func Load() {

	search := getSearch()
	repository := getRepository()

	developer :=
		beego.NewNamespace("/api",
			beego.NSRouter("", &developer.Static{}, "GET:ShowHomePage"),
			beego.NSNamespace("/table/",
				beego.NSRouter("list", &developer.Table{}, "GET:ShowList"),
				beego.NSRouter("generate/:name", &developer.Table{}, "*:GenerateTable"),
				beego.NSRouter("download/:name", &developer.Table{}, "*:DownloadTable"),
			),
			search,
			repository,
		)

	beego.AddNamespace(developer)
}
