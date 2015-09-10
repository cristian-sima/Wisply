package controllers

import (
	"html/template"
)

type WisplyController struct {
	MessageController
}

func (this *WisplyController) GenerateXsrf() {
	this.Data["xsrf_input"] = template.HTML(this.XsrfFormHtml())
}

func (c *WisplyController) getUserState() string {
	if c.IsUserConnected() {
		return "userConnected"
	}
	return "userDisconnected"
}

func (c *WisplyController) IsUserConnected() bool {
	v := c.GetSession("user")
	if v == nil {
		return false
	}
	return true
}

func (c *WisplyController) createMenu() {
	menuType := c.getUserState()

	switch menuType {
	case "userDisconnected":
		c.Data["userDisconnected"] = true
	case "userConnected":
		c.Data["userConnected"] = true
	case "adminConnected":
		c.Data["userConnected"] = true
		c.Data["adminConnected"] = true
	}
}

func (c *WisplyController) Prepare() {
	c.createMenu()
}
