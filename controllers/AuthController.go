package controllers

import (
	. "github.com/cristian-sima/Wisply/models/auth"
	"strconv"
	"strings"
)

type AuthController struct {
	DefaultController
	Model AuthModel
}

func (controller *AuthController) Prepare() {
	controller.WisplyController.Prepare()
}

func (controller *AuthController) ShowLoginForm() {
	if controller.AccountConnected {
		controller.Redirect("/", 302)
	} else {
		controller.GenerateXsrf()
		controller.Data["sendMe"] = strings.TrimSpace(controller.GetString("sendMe"))
		controller.TplNames = "site/auth/login.tpl"
		controller.Layout = "site/layout.tpl"
	}
}

func (controller *AuthController) ShowRegisterForm() {
	controller.GenerateXsrf()
	controller.TplNames = "site/auth/register.tpl"
	controller.Layout = "site/layout.tpl"
}

func (controller *AuthController) CreateNewAccount() {
	var username, password, confirmPassowrd, email string

	username = strings.TrimSpace(controller.GetString("register-username"))
	password = strings.TrimSpace(controller.GetString("register-password"))
	email = strings.TrimSpace(controller.GetString("register-email"))
	confirmPassowrd = strings.TrimSpace(controller.GetString("register-password-confirm"))

	rawData := make(map[string]interface{})
	rawData["username"] = username
	rawData["password"] = password
	rawData["email"] = email

	if confirmPassowrd != password {
		controller.DisplayErrorMessage("The passwords do not match!")
	} else {
		problems, err := controller.Model.ValidateRegisterDetails(rawData)
		if err != nil {
			controller.DisplayErrorMessages(problems)
		} else {
			usernameAlreadyExists := controller.Model.CheckAccountnameExists(username)
			if usernameAlreadyExists {
				controller.DisplayErrorMessage("The username is already used. Try another")
			} else {
				databaseError := controller.Model.CreateNewAccount(rawData)
				if databaseError != nil {
					controller.Abort("databaseError")
				} else {
					controller.DisplaySuccessMessage("Your account is ready!", "/auth/login/")
				}
			}
		}
	}
}

func (controller *AuthController) LoginAccount() {
	var sendMeAddress string = strings.TrimSpace(controller.GetString("login-send-me"))
	rawData := make(map[string]interface{})
	rawData["username"] = strings.TrimSpace(controller.GetString("login-username"))
	rawData["password"] = strings.TrimSpace(controller.GetString("login-password"))

	problems, err := controller.Model.ValidateLoginDetails(rawData)
	if err != nil {
		controller.DisplayErrorMessages(problems)
	} else {
		account, err := controller.Model.TryLoginAccount(rawData)
		if err != nil {
			controller.DisplayErrorMessage("There was a problem while login. We think the username or the password were not good.")
		} else {
			controller.connectAccount(account, sendMeAddress)
		}
	}
}

func (controller *AuthController) connectAccount(account *Account, sendMeAddress string) {
	controller.saveLoginDetails(account)
	controller.safeRedilectAccount(sendMeAddress)
}

func (controller *AuthController) saveLoginDetails(account *Account) {
	var accountId string = strconv.Itoa(account.Id)
	controller.Model.UpdateAccountLoginToken(accountId)
	controller.SetSession("account-id", accountId)
	controller.Ctx.SetCookie("wisply-connection", accountId, 1<<31-1, "/")
}

func (controller *AuthController) safeRedilectAccount(sendMe string) {
	var safeAddress string
	safeAddress = controller.getSafeURL(sendMe)
	controller.Redirect(safeAddress, 302)
}

func (controller *AuthController) getSafeURL(urlToTest string) string {
	var safeURL string = ""
	if urlToTest == "" || urlToTest == "/auth/login/" || urlToTest == "/auth/login" {
		safeURL = "/"
	} else {
		if controller.isSafeRedirection(urlToTest) {
			safeURL = urlToTest
		} else {
			safeURL = "/"
		}
	}
	return safeURL
}

func (controller *AuthController) isSafeRedirection(urlToTest string) bool {
	var isSafe bool
	isSafe = strings.HasPrefix(urlToTest, "/")
	return isSafe
}

func (controller *AuthController) Logout() {
	controller.DelSession("account")
	controller.DestroySession()
	controller.Redirect("/", 200)
}
