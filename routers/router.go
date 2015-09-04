package routers

import (
	"Wisply/controllers/general"
	"Wisply/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &general.MainController{})

    beego.Router("/about", &general.About{})

    beego.Router("/contact", &general.Contact{})
    beego.Router("/webscience", &general.Webscience{})

    beego.Router("/sample", &general.Sample{})


    beego.Router("/admin", &admin.Dashboard{})

    // source
    beego.Router("/admin/sources", &admin.SourceController{}, "*:List")
    beego.Router("/admin/sources/add", &admin.SourceController{}, "Get:AddNewSource")
    beego.Router("/admin/sources/add", &admin.SourceController{}, "Post:Insert")
    beego.Router("/admin/sources/modify/:id", &admin.SourceController{}, "Get:Edit")
    beego.Router("/admin/sources/modify/:id", &admin.SourceController{}, "Post:Update")

    beego.Router("/admin/sources/delete/:id", &admin.SourceController{}, "POST:Delete")

 }