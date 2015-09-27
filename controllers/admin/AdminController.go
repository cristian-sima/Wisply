package admin

import (
	general "github.com/cristian-sima/Wisply/controllers/general"
	model "github.com/cristian-sima/Wisply/models/admin"
)

// Controller must be inherited by all the pages that are for administrators
// It ensures that an account is connected when accesing the page
type Controller struct {
	general.WisplyController
}

// DisplayDashboard shows the administrator dashboard
func (controller *Controller) DisplayDashboard() {
	dashboard := model.NewDashboard()
	controller.Data["numberOfAccounts"] = dashboard.Accounts
	controller.Data["numberOfRepositories"] = dashboard.Repositories
	controller.Layout = "site/admin.tpl"
	controller.TplNames = "site/admin/dashboard.tpl"
}

// Prepare redirects to a login page in case the account is not connected, else it loads the page
func (controller *Controller) Prepare() {
	controller.WisplyController.Prepare()
	if !controller.AccountConnected || !controller.Account.IsAdministrator {
		controller.redirectAccount()
	} else {
		controller.initPage()
	}
}

// initPage is called when an administrator is connected
func (controller *Controller) initPage() {
	controller.Data["isAdminPage"] = true
}

// redirectAccount redirects the account to the login page
func (controller *Controller) redirectAccount() {

	loginPath := "/auth/login"
	addressParameter := "sendMe"

	currentPage := controller.Ctx.Request.URL.Path
	controller.Redirect(loginPath+"?"+addressParameter+"="+currentPage, 302)
}
