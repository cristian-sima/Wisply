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

    beego.Router("/admin/source", &admin.Source{}, "*:List")
    beego.Router("/admin/source/add", &admin.Source{}, "Get:Get")
    beego.Router("/admin/source/add", &admin.Source{}, "Post:Post")

 }
