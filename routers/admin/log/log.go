package log

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin/log"
)

// Get returns the namespace for log
func Get() func(*beego.Namespace) {
	ns := beego.NSNamespace("/log",
		beego.NSRouter("", &log.Home{}, "*:Display"),
		getHarvest(),
		beego.NSRouter("/advance-options", &log.Home{}, "*:DisplayAdvanceOptions"),
		beego.NSRouter("/delete", &log.Home{}, "POST:Delete"),
	)
	return ns
}
