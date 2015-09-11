package controllers

import (
	model "github.com/cristian-sima/Wisply/models/admin"
)

type AdminController struct {
	WisplyController
}

func (c *AdminController) ShowDashboard() {
	dashboard := model.GetDashboard()
	c.Data["numberOfUsers"] = dashboard.Users
	c.Data["numberOfSources"] = dashboard.Sources
	c.Layout = "general/admin.tpl"
	c.TplNames = "general/admin/dashboard.tpl"
}

func (controller *AdminController) Prepare() {
	controller.WisplyController.Prepare()
	if !controller.UserConnected || !controller.User.IsAdministrator() {
		controller.redirectUser()
	}
}

func (controller *AdminController) redirectUser() {
	var currentPage string = controller.Ctx.Request.URL.Path
	controller.Redirect("/auth/login?sendMe="+currentPage, 302)
}
