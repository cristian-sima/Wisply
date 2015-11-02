package data

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/developer/data"
)

// Get returns the Namespace for data
func Get() func(*beego.Namespace) {
	ns := beego.NSNamespace("/data/",
		beego.NSNamespace("/table",
			beego.NSRouter("", &data.Table{}, "GET:ShowList"),
			beego.NSRouter("generate/:name", &data.Table{}, "*:Generate"),
			beego.NSRouter("download/:name", &data.Table{}, "*:Download"),
		),
	)
	return ns
}
