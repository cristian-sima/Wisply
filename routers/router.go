package routers

import (
	"Wisply/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/about", &controllers.About{})
    beego.Router("/contact", &controllers.Contact{})
    beego.Router("/webscience", &controllers.Webscience{})

    beego.Router("/sample", &controllers.Sample{})
}