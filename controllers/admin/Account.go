// Package admin contains all the controllers of the application
package admin

import (
	"strings"

	auth "github.com/cristian-sima/Wisply/models/auth"
)

// Account manages the operations with the accounts
// For instance (delete, modify type, list all)
// It inherits the AdminController, thus an administrator account is required
type Account struct {
	Controller
	model auth.Model
}

// List lists all the Wisply accounts
func (controller *Account) List() {
	accounts := controller.model.GetAllAccounts()
	controller.Data["accounts"] = accounts
	controller.SetCustomTitle("Admin - Accounts")
	controller.TplNames = "site/admin/account/list.tpl"
}

// Modify shows the form to modify the type of an account
// There must be provided a paramater "id" which is the id of the account
func (controller *Account) Modify() {
	var id string
	id = controller.Ctx.Input.Param(":id")
	account, err := auth.NewAccount(id)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.showModifyForm(account)
	}
}

// Update modifies the type of the account given by parameter id
func (controller *Account) Update() {
	accountID := controller.Ctx.Input.Param(":id")
	newType := strings.TrimSpace(controller.GetString("modify-administrator"))
	account, err := auth.NewAccount(accountID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		err := account.ChangeType(newType)
		if err != nil {
			controller.DisplaySimpleError("There was a problem...")
		} else {
			successMessage := "The account has been modified!"
			goTo := "/admin/accounts/"
			controller.DisplaySuccessMessage(successMessage, goTo)
		}
	}
}

// Delete deletes the account given by parameter id
func (controller *Account) Delete() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")
	account, err := auth.NewAccount(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		account.Delete()
		successMessage := "The account [" + account.Email + "] has been deleted. Well done!"
		goTo := "/admin/accounts/"
		controller.DisplaySuccessMessage(successMessage, goTo)
	}
}

// showModifyForm shows the form to modify an account
func (controller *Account) showModifyForm(account *auth.Account) {
	controller.GenerateXSRF()
	controller.Data["accountName"] = account.Name
	if account.IsAdministrator {
		controller.Data["isAdministrator"] = true
	} else {
		controller.Data["isUser"] = true
	}
	controller.TplNames = "site/admin/account/form.tpl"
}
