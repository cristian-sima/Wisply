package controllers

import (
	"github.com/astaxie/beego"
)

type About struct {
	beego.Controller
}

func (c *About) Get() {
    c.Layout = "main/layout.tpl"
	c.TplNames = "main/page/about.tpl"
}