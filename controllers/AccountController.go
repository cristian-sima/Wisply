package controllers

import (
	. "github.com/cristian-sima/Wisply/models/auth"
	"strings"
)

type AccountController struct {
	AdminController
	model AuthModel
}

func (controller *AccountController) ListAccounts() {

	var exists bool = false

	accounts := controller.model.GetAllAccounts()

	exists = (len(accounts) != 0)

	controller.Data["anything"] = exists
	controller.Data["accounts"] = accounts
	controller.TplNames = "site/account/list.tpl"
	controller.Layout = "site/admin.tpl"
}

func (controller *AccountController) Modify() {

	var id string

	id = controller.Ctx.Input.Param(":id")

	account, err := NewAccount(id)

	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.showModifyForm(account)
	}
}

func (controller *AccountController) Update() {

	accountId := controller.Ctx.Input.Param(":id")

	newType := strings.TrimSpace(controller.GetString("modify-administrator"))

	account, err := NewAccount(accountId)
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

func (controller *AccountController) Delete() {
	var id string
	id = controller.Ctx.Input.Param(":id")
	account, err := NewAccount(id)
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

func (controller *AccountController) showModifyForm(account *Account) {
	controller.GenerateXsrf()
	controller.Data["accountName"] = account.Name

	if account.Administrator {
		controller.Data["isAdministrator"] = true
	} else {
		controller.Data["isUser"] = true
	}
	controller.Layout = "site/admin.tpl"
	controller.TplNames = "site/account/form.tpl"
}
