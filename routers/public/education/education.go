package education

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/public/education"
	"github.com/cristian-sima/Wisply/controllers/public/education/program"
)

// Get returns the namespace for public education
func Get() *beego.Namespace {
	ns := beego.NewNamespace("education",
		beego.NSRouter("", &education.Home{}, "*:Display"),
		beego.NSNamespace("program/:program",
			beego.NSRouter("", &program.Program{}, "GET:Display"),
		),
	)
	return ns
}
