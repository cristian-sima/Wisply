package account

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/account"
	"github.com/cristian-sima/Wisply/controllers/account/searches"
	"github.com/cristian-sima/Wisply/controllers/account/settings"
)

// Load tells the framework to load the addresses for the router
func Load() {
	accountNS :=
		beego.NewNamespace("/account",
			beego.NSRouter("", &account.Home{}, "GET:Show"),
			beego.NSNamespace("/searches/",
				beego.NSRouter("", &searches.List{}, "GET:Display"),
				beego.NSRouter("clear", &searches.List{}, "POST:Clear"),
			),
			beego.NSNamespace("/settings",
				beego.NSRouter("", &settings.Settings{}, "GET:Display"),
				beego.NSRouter("/delete", &settings.Settings{}, "POST:DeleteAccount"),
			),
		)
	beego.AddNamespace(accountNS)
}
