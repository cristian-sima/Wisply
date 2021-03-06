package account

// Home is the controller which manages the account dashboard
type Home struct {
	Controller
}

// Show displays the dashboard of an account
func (controller *Home) Show() {
	controller.SetCustomTitle("Account - Dashboard")
	controller.LoadTemplate("home")
}
