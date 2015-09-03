package general

import (
	"github.com/astaxie/beego"
)

type Contact struct {
	beego.Controller
}

func (c *Contact) Get() {
    c.Layout = "general/layout.tpl"
	c.TplNames = "general/page/contact.tpl"
}