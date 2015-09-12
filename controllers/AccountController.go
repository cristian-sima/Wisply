package controllers

import (
	"fmt"
	. "github.com/cristian-sima/Wisply/models/auth"
	"strings"
)

type AccountController struct {
	AdminController
	model AuthModel
}

func (controller *AccountController) ListAccounts() {

	var exists bool = false

	accounts := controller.model.GetAll()

	exists = (len(accounts) != 0)

	controller.Data["anything"] = exists
	controller.Data["accounts"] = accounts
	controller.TplNames = "site/account/list.tpl"
	controller.Layout = "site/admin.tpl"
}

func (controller *AccountController) Modify() {

	var id string

	id = controller.Ctx.Input.Param(":id")

	account, err := controller.model.GetAccountById(id)

	if err != nil {
		fmt.Println(err)
		controller.Abort("databaseError")
	} else {
		controller.showModifyForm(account)
	}
}

func (controller *AccountController) Update() {

	var accountId, isAdministrator string
	rawData := make(map[string]interface{})

	accountId = controller.Ctx.Input.Param(":id")

	isAdministrator = strings.TrimSpace(controller.GetString("modify-administrator"))
	rawData["administrator"] = isAdministrator

	_, err := controller.model.GetAccountById(accountId)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		problems, err := controller.model.ValidateModifyAccount(rawData)
		if err != nil {
			controller.DisplayErrorMessages(problems)
		} else {
			databaseError := controller.model.UpdateAccountType(accountId, isAdministrator)
			if databaseError != nil {
				controller.Abort("databaseError")
			} else {
				controller.DisplaySuccessMessage("The account has been modified!", "/admin/accounts/")
			}
		}
	}
}

func (controller *AccountController) Delete() {
	var id string
	id = controller.Ctx.Input.Param(":id")
	account, err := controller.model.GetAccountById(id)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		databaseError := controller.model.DeleteAccountById(id)
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
