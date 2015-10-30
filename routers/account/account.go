package account

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/account"
)

// Load tells the framework to load the addresses for the router
func Load() {
	accountNS :=
		beego.NewNamespace("/account",
			beego.NSRouter("", &account.Home{}, "GET:Show"),
			beego.NSNamespace("/search/",
				beego.NSRouter("", &account.Search{}, "GET:DisplayHistory"),
				beego.NSRouter("clear", &account.Search{}, "POST:ClearHistory"),
			),
			beego.NSNamespace("/settings",
				beego.NSRouter("", &account.Settings{}, "GET:DisplayPage"),
				beego.NSRouter("/delete", &account.Settings{}, "POST:DeleteAccount"),
			),
		)
	beego.AddNamespace(accountNS)
}
