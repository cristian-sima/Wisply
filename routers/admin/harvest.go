package admin

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin"
)

func getHarvest() func(*beego.Namespace) {
	ns := beego.NSNamespace("/harvest",
		beego.NSNamespace("/init",
			beego.NSRouter(":id", &admin.Harvest{}, "POST:ShowPanel"),
			beego.NSRouter("/ws", &admin.Harvest{}, "GET:InitWebsocketConnection"),
		),
		beego.NSNamespace("/recover",
			beego.NSRouter(":id", &admin.Harvest{}, "POST:RecoverProcess"),
		),
		beego.NSNamespace("/finish",
			beego.NSRouter(":id", &admin.Harvest{}, "POST:ForceFinishProcess"),
		),
	)
	return ns
}
