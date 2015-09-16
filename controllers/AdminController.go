package controllers

import (
	model "github.com/cristian-sima/Wisply/models/admin"
)

// AdminController This controller must be inherited by all the pages that are for administrators
// It ensures that an account is connected when accesing the page
type AdminController struct {
	WisplyController
}

// ShowDashboard It shows the administrator dashboard
func (controller *AdminController) ShowDashboard() {
	dashboard := model.GetDashboard()
	controller.Data["numberOfAccounts"] = dashboard.Accounts
	controller.Data["numberOfSources"] = dashboard.Sources
	controller.Layout = "site/admin.tpl"
	controller.TplNames = "site/admin/dashboard.tpl"
}

// Prepare If the account is not connect it redirects to a login page, else it loads the page
func (controller *AdminController) Prepare() {
	controller.WisplyController.Prepare()
	if !controller.AccountConnected || !controller.Account.IsAdministrator() {
		controller.redirectAccount()
	} else {
		controller.initPage()
	}
}

// initPage it is called when an administrator is connected
func (controller *AdminController) initPage() {
	controller.Data["isAdminPage"] = true
}

// redirectAccount It redirects the account to the login page
func (controller *AdminController) redirectAccount() {

	loginPath := "/auth/login"
	addressParameter := "sendMe"

	currentPage := controller.Ctx.Request.URL.Path
	controller.Redirect(loginPath+"?"+addressParameter+"="+currentPage, 302)
}
