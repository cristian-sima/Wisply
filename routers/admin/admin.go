package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

// Load tells the framework to load the addresses for the router
func Load() {

	account := getAccount()
	curriculum := getCurriculum()
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
