package settings

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/auth"
)

// Settings is the controller which manages the operation for the settings page
type Settings struct {
	Controller
}

// Display shows the settings page
func (controller *Settings) Display() {
	controller.SetCustomTitle("Account - Settings")
	controller.LoadTemplate("home")
}

// DeleteAccount checks if the supplied password is correct and then
// deletes the account
func (controller *Settings) DeleteAccount() {
	password := strings.TrimSpace(controller.GetString("password"))
	account := controller.Account
	isPasswordValid := auth.VerifyAccount(account, password)
	if isPasswordValid {
		controller.Account.Delete()
	} else {
		controller.Abort("404")
	}
	controller.ShowBlankPage()
}
