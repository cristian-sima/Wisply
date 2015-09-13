package controllers

import (
	. "github.com/cristian-sima/Wisply/models/auth"
	"strconv"
	"strings"
	"fmt"
)

type AuthController struct {
	WisplyController
	Model AuthModel
}

func (controller *AuthController) Prepare() {
	controller.WisplyController.Prepare()
}

func (controller *AuthController) ShowLoginForm() {
	if controller.AccountConnected {
		controller.Redirect("/", 302)
	} else {
		controller.Data["sendMe"] = strings.TrimSpace(controller.GetString("sendMe"))
		controller.showForm("login")
	}
}

func (controller *AuthController) ShowRegisterForm() {
	controller.showForm("register")
}

func (controller *AuthController) showForm(name string) {
	controller.GenerateXsrf()
	controller.TplNames = "site/auth/" + name + ".tpl"
	controller.Layout = "site/layout.tpl"
}

func (controller *AuthController) CreateNewAccount() {
	var name, password, confirmPassowrd, email string

	name = strings.TrimSpace(controller.GetString("register-name"))
	password = strings.TrimSpace(controller.GetString("register-password"))
	email = strings.TrimSpace(controller.GetString("register-email"))
	confirmPassowrd = strings.TrimSpace(controller.GetString("register-password-confirm"))

	userDetails := make(map[string]interface{})
	userDetails["name"] = name
	userDetails["password"] = password
	userDetails["email"] = email

	if confirmPassowrd != password {
		controller.DisplaySimpleError("The passwords do not match!")
	} else {
		register := Register{}
		problem, err := register.Try(userDetails)
		if err != nil {
			controller.DisplayError(problem)
		} else {
			controller.DisplaySuccessMessage("Your account is ready!", "/auth/login/")
		}
	}
}

func (controller *AuthController) LoginAccount() {
	var sendMeAddress string = strings.TrimSpace(controller.GetString("login-send-me"))
	loginDetails := make(map[string]interface{})
	loginDetails["email"] = strings.TrimSpace(controller.GetString("login-email"))
	loginDetails["password"] = strings.TrimSpace(controller.GetString("login-password"))

	login := Login{}
	problems, err := login.Try(loginDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		account, _ := GetAccountByEmail(loginDetails["email"].(string))
		controller.connectAccount(account, sendMeAddress)
	}
}

func (controller *AuthController) connectAccount(account *Account, sendMeAddress string) {
	controller.saveLoginDetails(account)
	controller.safeRedilectAccount(sendMeAddress)
}

func (controller *AuthController) saveLoginDetails(account *Account) {
	accountId := strconv.Itoa(account.Id)
	controller.SetSession("account-id", accountId)
	cookie := account.GenerateConnectionCookie()
	fmt.Println("asta e: ")
	fmt.Println(cookie)
	controller.Ctx.SetCookie(cookie.Name, cookie.GetValue(), strconv.Itoa(cookie.Duration), cookie.Path)
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
	controller.distroySession()
	controller.deleteConnectionCookie()
	controller.Redirect("/", 200)
}

func (controller *AuthController) distroySession() {
	controller.DelSession("session")
	controller.DestroySession()
}

func (controller *AuthController) deleteConnectionCookie() {
	cookieName := Settings["cookieName"].(string)
	cookiePath := Settings["cookiePath"].(string)
	cookie := controller.Ctx.GetCookie(cookieName)
	if cookie != "" {
		controller.Ctx.SetCookie(cookieName, "", "-1", cookiePath)
	}
}
