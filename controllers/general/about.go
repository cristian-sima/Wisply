package general

import (
	"github.com/astaxie/beego"
)

type About struct {
	beego.Controller
}

func (c *About) Get() {
    c.Layout = "general/layout.tpl"
	c.TplNames = "general/page/about.tpl"
}