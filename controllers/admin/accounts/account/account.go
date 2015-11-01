package account

import "strings"

// Account manages the operations with an account
type Account struct {
	Controller
}

// ShowModifyForm shows the form to modify the type of an account
func (controller *Account) ShowModifyForm() {
	controller.GenerateXSRF()
	controller.LoadTemplate("form")
}

// Modify modifies the type of the account
func (controller *Account) Modify() {
	newType := strings.TrimSpace(controller.GetString("modify-administrator"))
	err := controller.GetAccount().ChangeType(newType)
	if err != nil {
		controller.DisplaySimpleError("There was a problem...")
	} else {
		successMessage := "The account has been modified."
		goTo := "/admin/accounts/"
		controller.DisplaySuccessMessage(successMessage, goTo)
	}
}

// Delete deletes the account
func (controller *Account) Delete() {
	account := controller.GetAccount()
	err := account.Delete()
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		successMessage := "The account [" + account.Email + "] has been deleted."
		goTo := "/admin/accounts/"
		controller.DisplaySuccessMessage(successMessage, goTo)
	}
}
