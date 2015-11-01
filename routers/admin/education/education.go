package education

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/admin/education"
)

// Get returns the routers for the education
func Get() func(*beego.Namespace) {
	program := getProgram()
	ns := beego.NSNamespace("/education",
		beego.NSRouter("", &education.Home{}, "*:Display"),
		program,
	)
	return ns
}
