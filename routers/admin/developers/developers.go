package developers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin/developers"
	"github.com/cristian-sima/Wisply/controllers/admin/developers/table"
)

// Get returns the Namespace for the developers within administration area
func Get() func(*beego.Namespace) {

	ns := beego.NSNamespace("/developers",
		beego.NSRouter("", &developers.Home{}, "*:Display"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &table.Table{}, "GET:ShowAddForm"),
			beego.NSRouter("", &table.Table{}, "POST:InsertNewTable"),
		),
		beego.NSNamespace("/table/:table",
			beego.NSNamespace("/delete",
				beego.NSRouter("", &table.Table{}, "POST:RemoveAllowedTable"),
			),
			beego.NSNamespace("/modify",
				beego.NSRouter("", &table.Table{}, "GET:ShowModifyForm"),
				beego.NSRouter("", &table.Table{}, "POST:ModifyTable"),
			),
		),
	)
	return ns

}
