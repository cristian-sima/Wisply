package controllers

import (
	"github.com/astaxie/beego"
)

type Webscience struct {
	beego.Controller
}

func (c *Webscience) Get() {
    c.Layout = "main/layout.tpl"
	c.TplNames = "main/page/webscience.tpl"
}
