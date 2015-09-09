package routers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers"
)

func init() {

	beego.Router("/", &controllers.DefaultController{}, "*:ShowIndexPage")
	beego.Router("/about", &controllers.DefaultController{}, "*:ShowAboutPage")
	beego.Router("/contact", &controllers.DefaultController{}, "*:ShowContactPage")
	beego.Router("/webscience", &controllers.DefaultController{}, "*:ShowWebsciencePage")
	beego.Router("/sample", &controllers.DefaultController{}, "*:ShowSamplePage")

	// ----------------------------- Admin --------------------------------------


	sourcesNamespace := beego.NSNamespace("/sources",
		beego.NSRouter("", &controllers.SourceController{}, "*:ListSources"),
		beego.NSNamespace("/add",
				beego.NSRouter("", &controllers.SourceController{}, "Get:AddNewSource"),
				beego.NSRouter("", &controllers.SourceController{}, "Post:InsertSource"),
		),
		beego.NSNamespace("/modify",
				beego.NSRouter(":id", &controllers.SourceController{}, "Get:Modify"),
				beego.NSRouter(":id", &controllers.SourceController{}, "Post:Update"),
		),
		beego.NSNamespace("/delete",
				beego.NSRouter(":id", &controllers.SourceController{}, "Post:Delete"),
		),
	)

	adminNamespace :=
		beego.NewNamespace("/admin",
			beego.NSRouter("",  &controllers.AdminController{}, "*:ShowDashboard"),
			sourcesNamespace,
		)

	// register namespace
	beego.AddNamespace(adminNamespace)

}
