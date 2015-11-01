// Package developer contains all the addresses for the developers area
package developer

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/developer"
	"github.com/cristian-sima/Wisply/routers/developer/api"
	"github.com/cristian-sima/Wisply/routers/developer/data"
)

// Load tells the framework to load the addresses for the router
func Load() {

	developer :=
		beego.NewNamespace("/developer",
			beego.NSRouter("", &developer.Home{}, "GET:Display"),
			data.Get(),
			api.Get(),
		)

	beego.AddNamespace(developer)
}
