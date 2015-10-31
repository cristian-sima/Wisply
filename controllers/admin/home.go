package admin

import admin "github.com/cristian-sima/Wisply/models/admin"

// Home manages the operations for the dashboard/home page for admin
type Home struct {
	Controller
}

// DisplayDashboard shows the administrator dashboard
func (controller *Home) DisplayDashboard() {
	dashboard := admin.NewDashboard()
	controller.Data["numberOfAccounts"] = dashboard.Accounts
	controller.Data["numberOfRepositories"] = dashboard.Repositories
	controller.TplNames = "site/admin/admin/dashboard.tpl"
	controller.SetCustomTitle("Admin - Dashboard")
}
