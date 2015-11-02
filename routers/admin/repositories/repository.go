package repositories

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin/repositories"
	"github.com/cristian-sima/Wisply/controllers/admin/repositories/repository"
)

// Get returns the namespace for institutions
func Get() func(*beego.Namespace) {

	ns := beego.NSNamespace("/repositories",
		beego.NSRouter("", &repositories.Home{}, "*:Display"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &repository.Repository{}, "GET:ShowChooseCategory"),
		),
		beego.NSNamespace("/newRepository",
			beego.NSRouter("", &repository.Repository{}, "GET:ShowInsertForm"),
			beego.NSRouter("", &repository.Repository{}, "POST:Insert"),
		),
		beego.NSNamespace("/:repository",
			beego.NSRouter("", &repository.Repository{}, "GET:Display"),
			beego.NSNamespace("/advance-options",
				beego.NSRouter("", &repository.Repository{}, "GET:ShowAdvanceOptions"),
				beego.NSNamespace("/modify",
					beego.NSRouter("", &repository.Repository{}, "GET:ShowModifyForm"),
					beego.NSRouter("", &repository.Repository{}, "POST:Modify"),
					beego.NSNamespace("/filter",
						beego.NSRouter("", &repository.Repository{}, "GET:ShowFilterForm"),
						beego.NSRouter("", &repository.Repository{}, "POST:ModifyFilter"),
					),
				),
			),
			beego.NSNamespace("/clear",
				beego.NSRouter("", &repository.Repository{}, "POST:ClearRepository"),
			),
			beego.NSNamespace("/delete",
				beego.NSRouter("", &repository.Repository{}, "POST:Delete"),
			),
		),
	)
	return ns
}
