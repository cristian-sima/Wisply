package controllers

import (
	model "github.com/cristian-sima/Wisply/models/admin"
)

type AdminController struct {
	WisplyController
}

func (c *AdminController) ShowDashboard() {
	dashboard := model.GetDashboard()
	c.Data["numberOfAccounts"] = dashboard.Accounts
	c.Data["numberOfSources"] = dashboard.Sources
	c.Layout = "site/admin.tpl"
	c.TplNames = "site/admin/dashboard.tpl"
}

func (controller *AdminController) Prepare() {
	controller.WisplyController.Prepare()
	if !controller.AccountConnected || !controller.Account.IsAdministrator() {
		controller.redirectAccount()
	} else {
		controller.initPage()
	}
}

func (controller *AdminController) initPage() {
	controller.Data["isAdminPage"] = true
}

func (controller *AdminController) redirectAccount() {
	var currentPage string = controller.Ctx.Request.URL.Path
	controller.Redirect("/auth/login?sendMe="+currentPage, 302)
}
