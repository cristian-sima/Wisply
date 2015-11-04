package education

import (
	"github.com/astaxie/beego"

	subject "github.com/cristian-sima/Wisply/controllers/admin/education/subject"
)

func getSubject() func(*beego.Namespace) {
	ns := beego.NSNamespace("/subjects",
		beego.NSNamespace("/:subject",
			beego.NSRouter("", &subject.Subject{}, "*:Display"),
			beego.NSNamespace("/advance-options",
				beego.NSRouter("", &subject.Subject{}, "*:ShowAdvanceOptions"),
				beego.NSNamespace("/modify",
					beego.NSRouter("", &subject.Subject{}, "GET:ShowModifyForm"),
					beego.NSRouter("", &subject.Subject{}, "POST:Update"),
				),
				beego.NSNamespace("/modify-static-description",
					beego.NSRouter("", &subject.Subject{}, "GET:ShowModifyDescription"),
					beego.NSRouter("", &subject.Subject{}, "POST:UpdateDescription"),
				),
			),
			beego.NSNamespace("/delete",
				beego.NSRouter("", &subject.Subject{}, "POST:Delete"),
			),
			// definition
			beego.NSNamespace("/definition",
				beego.NSNamespace("/add",
					beego.NSRouter("", &subject.Definition{}, "GET:ShowAddForm"),
					beego.NSRouter("", &subject.Definition{}, "POST:CreateDefinition"),
				),
				beego.NSNamespace("/:definition",
					beego.NSNamespace("/modify",
						beego.NSRouter("", &subject.Definition{}, "GET:ShowModifyForm"),
						beego.NSRouter("", &subject.Definition{}, "POST:Update"),
					),
					beego.NSNamespace("/delete",
						beego.NSRouter("", &subject.Definition{}, "POST:Delete"),
					),
				),
			),
			// ka
			beego.NSNamespace("/ka",
				beego.NSNamespace("/add",
					beego.NSRouter("", &subject.KA{}, "GET:ShowAddForm"),
					beego.NSRouter("", &subject.KA{}, "POST:CreateKA"),
				),
				beego.NSNamespace("/:ka",
					beego.NSNamespace("/modify",
						beego.NSRouter("", &subject.KA{}, "GET:ShowModifyForm"),
						beego.NSRouter("", &subject.KA{}, "POST:Update"),
					),
					beego.NSNamespace("/delete",
						beego.NSRouter("", &subject.KA{}, "POST:Delete"),
					),
				),
			),
		),
		beego.NSNamespace("add",
			beego.NSRouter("", &subject.Subject{}, "GET:ShowAddForm"),
			beego.NSRouter("", &subject.Subject{}, "POST:CreateSubject"),
		),
	)
	return ns
}
