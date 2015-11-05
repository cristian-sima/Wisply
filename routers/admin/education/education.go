package education

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin/education"
)

// Get returns the routers for the education
func Get() func(*beego.Namespace) {
	subject := getSubject()
	ns := beego.NSNamespace("/education",
		beego.NSRouter("/analyse", &education.Home{}, "GET:Analyse"),
		beego.NSRouter("/analyser/:analyser/delete", &education.Home{}, "POST:DeleteAnalyser"),
		beego.NSRouter("", &education.Home{}, "*:Display"),
		subject,
	)
	return ns
}
