package routers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers"
)

func init() {


    beego.Router("/",               &controllers.DefaultController{}, "*:ShowIndexPage")
    beego.Router("/about",          &controllers.DefaultController{}, "*:ShowAboutPage")
    beego.Router("/contact",        &controllers.DefaultController{}, "*:ShowContactPage")
    beego.Router("/webscience",     &controllers.DefaultController{}, "*:ShowWebsciencePage")
    beego.Router("/sample",         &controllers.DefaultController{}, "*:ShowSamplePage")

    beego.Router("/admin", &controllers.AdminController{}, "*:ShowDashboard")

    // source
    beego.Router("/admin/sources", &controllers.SourceController{}, "*:ListSources")
    beego.Router("/admin/sources/add", &controllers.SourceController{}, "Get:AddNewSource")
    beego.Router("/admin/sources/add", &controllers.SourceController{}, "Post:InsertSource")
    beego.Router("/admin/sources/modify/:id", &controllers.SourceController{}, "Get:Modify")
    beego.Router("/admin/sources/modify/:id", &controllers.SourceController{}, "Post:Update")

    beego.Router("/admin/sources/delete/:id", &controllers.SourceController{}, "POST:Delete")

 }
