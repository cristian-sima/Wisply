package admin

import admin "github.com/cristian-sima/Wisply/models/admin"

// Home manages the operations for the dashboard/home page for admin
type Home struct {
	Controller
}

// Display shows the administrator dashboard
func (controller *Home) Display() {
	controller.SetCustomTitle("Admin - Dashboard")
	dashboard := admin.NewDashboard()
	controller.Data["numberOfAccounts"] = dashboard.Accounts
	controller.Data["numberOfRepositories"] = dashboard.Repositories
	controller.LoadTemplate("home")
}
