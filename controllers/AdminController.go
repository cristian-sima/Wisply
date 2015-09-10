package controllers

import (
	model "github.com/cristian-sima/Wisply/models/admin"
)

type AdminController struct {
	WisplyController
}

func (c *AdminController) ShowDashboard() {

	var numberOfUsers, numberOfSources int

	dashboard := model.GetDashboard()

	numberOfUsers = dashboard.Users
	numberOfSources = dashboard.Sources

	c.Data["numberOfUsers"] = numberOfUsers
	c.Data["numberOfSources"] = numberOfSources

	c.Layout = "general/admin.tpl"
	c.TplNames = "general/admin/dashboard.tpl"
}

func (controller *AdminController) Prepare() {
	controller.WisplyController.Prepare()
	if !controller.IsUserConnected() {
		controller.UserIsNotConnected()
	}
}

func (controller *AdminController) UserIsNotConnected() {
	var currentPage string = controller.Ctx.Request.URL.Path
	controller.Redirect("/auth/login?sendMe="+currentPage, 302)
}
