package controllers

import (
	"github.com/astaxie/beego"
)

type Contact struct {
	beego.Controller
}

func (c *Contact) Get() {
    c.Layout = "main/layout.tpl"
	c.TplNames = "main/page/contact.tpl"
}