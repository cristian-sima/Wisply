// Package admin contains all the addresses for the administration area
package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
	"github.com/cristian-sima/Wisply/routers/admin/accounts"
	"github.com/cristian-sima/Wisply/routers/admin/developers"
	"github.com/cristian-sima/Wisply/routers/admin/education"
	"github.com/cristian-sima/Wisply/routers/admin/institutions"
	"github.com/cristian-sima/Wisply/routers/admin/log"
	"github.com/cristian-sima/Wisply/routers/admin/repositories"
)

// Load tells the framework to load the addresses for the router
func Load() {

	account := accounts.Get()
	curriculum := education.Get()
	developer := developers.Get()
	institutions := institutions.Get()
	harvest := getHarvest()
	log := log.Get()
	repository := repositories.Get()

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("", &admin.Home{}, "*:Display"),
		account,
		curriculum,
		developer,
		institutions,
		harvest,
		log,
		repository,
	)
	beego.AddNamespace(ns)
}
