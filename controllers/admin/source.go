package admin

import (
	"github.com/astaxie/beego"
)

type Source struct {
	beego.Controller
}

func (c *Source) List() {
    c.Layout = "general/admin.tpl"
	c.TplNames = "general/source/list.tpl"
}

func (c *Source) Get() {
    c.Layout = "general/admin.tpl"
	c.TplNames = "general/source/add.tpl"
}

func (c *Source) Post() {

    c.Data["messageContent"] = "The source has been added!";
    c.Data["messageLink"] = "/admin/source";

    c.Layout = "general/status.tpl"
	c.TplNames = "general/message/success.tpl"
}