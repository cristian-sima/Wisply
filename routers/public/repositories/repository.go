package repositories

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/public/repositories/repository"
)

// Get returns the Namespace for the repositories
func Get() *beego.Namespace {
	ns := beego.NewNamespace("/repositories",
		beego.NSNamespace("/:repository",
			beego.NSRouter("", &repository.Repository{}, "GET:Display"),
			beego.NSNamespace("/resource",
				beego.NSNamespace("/:resource",
					beego.NSRouter("", &repository.Repository{}, "GET:Display"),
				),
			),
		),
	)
	return ns
}
