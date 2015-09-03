package general

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Sample struct {
	beego.Controller
}

func (c *MainController) Get() {
    c.Layout = "general/layout.tpl"
	c.TplNames = "general/page/index.tpl"
}


// sample
func (c *Sample) Get() {
	c.TplNames = "sample.tpl"
}

