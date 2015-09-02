package routers

import (
	"Wisply/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})



    beego.Router("/sample", &controllers.Sample{})
}