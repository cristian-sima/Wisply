package account

import (
	"github.com/cristian-sima/Wisply/controllers/admin/accounts"
	"github.com/cristian-sima/Wisply/models/auth"
)

// Controller manages the operations with the accounts
type Controller struct {
	accounts.Controller
	account *auth.Account
}

// Prepare loads the account
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/accounts/account")
	controller.loadAccount()
}

// GetAccount returns the reference to the account
func (controller *Controller) GetAccount() *auth.Account {
	return controller.account
}

func (controller *Controller) loadAccount() {
	ID := controller.Ctx.Input.Param(":account")
	account, err := auth.NewAccount(ID)
	if err != nil {
		controller.Abort("show-database-error")
	}
	controller.Data["account"] = account
	controller.account = account
	controller.SetCustomTitle("Admin - " + account.Name)
	controller.account = account
}
