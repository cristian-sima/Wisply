package controllers

import (
	"fmt"
	"html/template"
)

type WisplyController struct {
	MessageController
}

func (this *WisplyController) GenerateXsrf() {
	this.Data["xsrf_input"] = template.HTML(this.XsrfFormHtml())
}

func (c *WisplyController) getMenu() map[string]string {
	items := make(map[string]string)

	if c.isUserLogged() {
		items["Log out"] = "/auth/logout"
	} else {
		items["Login"] = "/auth/login"
		items["Register"] = "/auth/register"
	}
	return items
}

func (c *WisplyController) createMenu() {
	items := c.getMenu()
	c.Data["TopLeftMenuItems"] = items
}

func (c *WisplyController) Prepare() {
	c.createMenu()
}

func (c *WisplyController) isUserLogged() bool {
	v := c.GetSession("user")
	fmt.Println(v)
	if v == nil {
		return false
	}
	return true
}
