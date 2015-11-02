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
		beego.NSRouter("/:institution", &institution.Institution{}, "GET:Display"),
	)
	return ns
}
