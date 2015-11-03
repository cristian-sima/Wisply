package tools

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/developer/tools"
)

// Get returns the Namespace for data
func Get() func(*beego.Namespace) {
	ns := beego.NSNamespace("/tools",
		beego.NSNamespace("/digester",
			beego.NSRouter("", &tools.Digester{}, "GET:Display"),
			beego.NSRouter("", &tools.Digester{}, "POST:Work"),
		),
		beego.NSNamespace("/web-digester",
			beego.NSRouter("", &tools.WebDigester{}, "GET:Display"),
			beego.NSRouter("", &tools.WebDigester{}, "POST:Work"),
		),
	)
	return ns
}
