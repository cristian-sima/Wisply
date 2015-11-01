package accounts

import (
	"github.com/astaxie/beego"
	accountsController "github.com/cristian-sima/Wisply/controllers/admin/accounts"
	accountController "github.com/cristian-sima/Wisply/controllers/admin/accounts/account"
)

// Get returns the Namespace for the accounts within administration area
func Get() func(*beego.Namespace) {
	ns := beego.NSNamespace("/accounts",
		beego.NSRouter("", &accountsController.Home{}, "*:Display"),
		beego.NSNamespace("/:account/:id",
			beego.NSRouter("", &accountController.Account{}, "GET:ShowModifyForm"),
			beego.NSRouter("", &accountController.Account{}, "POST:Modify"),
			beego.NSRouter("/delete", &accountController.Account{}, "POST:Delete"),
		),
	)
	return ns
}
