package general

import (
	"github.com/astaxie/beego"
)

type Webscience struct {
	beego.Controller
}

func (c *Webscience) Get() {
    c.Layout = "general/layout.tpl"
	c.TplNames = "general/page/webscience.tpl"
}