package education

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/public/education"
	"github.com/cristian-sima/Wisply/controllers/public/education/subject"
)

// Get returns the namespace for public education
func Get() *beego.Namespace {
	ns := beego.NewNamespace("education",
		beego.NSRouter("", &education.Home{}, "*:Display"),
		beego.NSNamespace("subjects/:subject",
			beego.NSRouter("", &subject.Subject{}, "GET:Display"),
		),
	)
	return ns
}
