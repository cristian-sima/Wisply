package admin

import (
	"github.com/astaxie/beego"
)

type Dashboard struct {
	beego.Controller
}

func (c *Dashboard) Get() {
    c.Layout = "general/admin.tpl"
	c.TplNames = "general/admin/dashboard.tpl"
}