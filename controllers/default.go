package controllers

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
    c.Layout = "main/layout.tpl"
	c.TplNames = "main/index.tpl"
}


// sample
func (c *Sample) Get() {
	c.TplNames = "sample.tpl"
}
