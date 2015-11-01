// Package admin contains all the addresses for the administration area
package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
	"github.com/cristian-sima/Wisply/routers/admin/accounts"
	"github.com/cristian-sima/Wisply/routers/admin/education"
)

// Load tells the framework to load the addresses for the router
func Load() {

	account := accounts.Get()
	curriculum := curriculum.Get()
	developer := getDeveloper()
	institution := getInstitution()
	harvest := getHarvest()
	log := getLog()
	repository := getRepository()

	ns := beego.NewNamespace("/admin",
		beego.NSRouter("", &admin.Home{}, "*:DisplayDashboard"),
		account,
		curriculum,
		developer,
		institution,
		harvest,
		log,
		repository,
	)
	beego.AddNamespace(ns)
}
