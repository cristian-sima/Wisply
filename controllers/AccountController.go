package controllers

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/auth"
)

// AccountController It manages the operations with the accounts (such as delete, modify type, list all)
// It inherits the AdminController, thus an administrator account is required
type AccountController struct {
	AdminController
	model auth.AuthModel
}

// ListAccounts It lists all the Wisply accounts
func (controller *AccountController) ListAccounts() {
	accounts := controller.model.GetAllAccounts()
	controller.Data["accounts"] = accounts
	controller.TplNames = "site/account/list.tpl"
	controller.Layout = "site/admin.tpl"
}

// Modify It shows the form to modify the type of an account
// There must be provided a paramater "id" which is the id of the account
func (controller *AccountController) Modify() {
	var id string
	id = controller.Ctx.Input.Param(":id")
	account, err := auth.NewAccount(id)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.showModifyForm(account)
	}
}

// Update It modifies the type of the account given by parameter id
func (controller *AccountController) Update() {
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
			controller.DisplaySuccessMessage("The account has been modified!", "/admin/accounts/")
		}
	}
}

// Delete It deletes the account given by parameter id
func (controller *AccountController) Delete() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")
	account, err := auth.NewAccount(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		databaseError := account.Delete()
		if databaseError != nil {
			controller.Abort("databaseError")
		} else {
			controller.DisplaySuccessMessage("The account ["+account.Email+"] has been deleted. Well done!", "/admin/accounts/")
		}
	}
}

// It shows the form to modify an account
func (controller *AccountController) showModifyForm(account *auth.Account) {
	controller.GenerateXSRF()
	controller.Data["accountName"] = account.Name
	if account.Administrator {
		controller.Data["isAdministrator"] = true
	} else {
		controller.Data["isUser"] = true
	}
	controller.Layout = "site/admin.tpl"
	controller.TplNames = "site/account/form.tpl"
}
