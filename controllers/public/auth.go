package public

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/auth"
)

// Auth manages the operations with the authentication
type Auth struct {
	Controller
}

// Prepare sets the template path
func (controller *Auth) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("public/auth")
}

// ShowLoginPage shows the login form in case user is not connected or otherwise
// it redirects to "/"
func (controller *Auth) ShowLoginPage() {
	if controller.AccountConnected {
		controller.Redirect("/", 302)
	} else {
		controller.showLoginForm()
	}
}

// ShowRegisterForm shows the form to register a new account
func (controller *Auth) ShowRegisterForm() {
	controller.SetCustomTitle("Create a new account")
	controller.showForm("register")
	controller.LoadCaptcha("register-form-page")
}

// CreateNewAccount checks if the password and the confirmation are the same
// If so it sends the details of the user to processRegisterRequest
// The parameters should be: register-name, register-password,
// register-email and register-password-confirm
func (controller *Auth) CreateNewAccount() {
	page := "register-form-page"
	controller.RegisterCaptchaAction(page)

	if !controller.IsCaptchaValid(page) {
		controller.DisplaySimpleError("Please enter a valid code!")
	} else {
		controller.createNewAccount()
	}
}

// LoginAccount checks if the details provided are good and it logins the account
func (controller *Auth) LoginAccount() {
	page := "login-form-page"

	controller.RegisterCaptchaAction(page)

	if !controller.IsCaptchaValid(page) {
		controller.DisplaySimpleError("Please enter a valid code!")
	} else {
		controller.loginAccount()
	}
}

// Logout It logs out the account
func (controller *Auth) Logout() {
	controller.distroySession()
	controller.DeleteConnectionCookie()
	controller.Redirect("/", 200)
}

// showLoginForm shows the login form
func (controller *Auth) showLoginForm() {
	rawSendMe := controller.GetString("sendMe")
	controller.Data["sendMe"] = strings.TrimSpace(rawSendMe)
	controller.SetCustomTitle("Login to Wisply")
	controller.showForm("login")
	controller.LoadCaptcha("login-form-page")
}

// showForm shows a form indicated by the parameter name.
// It can be "login" or "register"
func (controller *Auth) showForm(name string) {
	controller.GenerateXSRF()
	controller.LoadTemplate(name)
}

func (controller *Auth) createNewAccount() {

	confirmPassowrd := strings.TrimSpace(controller.GetString("register-password-confirm"))
	password := strings.TrimSpace(controller.GetString("register-password"))
	userDetails := make(map[string]interface{})
	userDetails["name"] = strings.TrimSpace(controller.GetString("register-name"))
	userDetails["password"] = password
	userDetails["email"] = strings.TrimSpace(controller.GetString("register-email"))

	if confirmPassowrd != password {
		controller.DisplaySimpleError("The passwords do not match!")
	} else {
		controller.processRegisterRequest(userDetails)
	}
}

func (controller *Auth) processRegisterRequest(userDetails map[string]interface{}) {

	register := auth.Register{}
	problem, err := register.Try(userDetails)
	if err != nil {
		controller.DisplayError(problem)
	} else {
		message := "Your account is ready!"
		goTo := "/auth/login/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// LoginAccount checks if the details provided are good and it logins the account
func (controller *Auth) loginAccount() {

	sendMeAddress := strings.TrimSpace(controller.GetString("login-send-me"))
	rememberMe := strings.TrimSpace(controller.GetString("login-remember-me"))

	loginDetails := make(map[string]interface{})
	loginDetails["email"] = strings.TrimSpace(controller.GetString("login-email"))
	loginDetails["password"] = strings.TrimSpace(controller.GetString("login-password"))

	login := auth.Login{}
	problems, err := login.Try(loginDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		account, _ := auth.GetAccountByEmail(loginDetails["email"].(string))
		if rememberMe == "on" {
			controller.rememberConnection(account)
		}
		controller.connectAccount(account, sendMeAddress)
	}
}

// connectAccount creates a session for the account and redirects
func (controller *Auth) connectAccount(account *auth.Account, sendMeAddress string) {
	controller.saveLoginDetails(account)
	controller.safeRedilectAccount(sendMeAddress)
}

// saveLoginDetails creates a new session for the account
func (controller *Auth) saveLoginDetails(account *auth.Account) {
	accountID := strconv.Itoa(account.ID)
	controller.SetSession("account-id", accountID)
}

// rememberConnection remembers the account by using a connection cookie
func (controller *Auth) rememberConnection(account *auth.Account) {
	cookieName := auth.Settings["cookieName"].(string)
	cookie := account.GenerateConnectionCookie()
	controller.DeleteConnectionCookie()
	controller.Ctx.SetCookie(cookieName, cookie.GetValue(), cookie.Duration, cookie.Path)
}

// safeRedilectAccount gets the safe address where to redirects the account
// It redirects the account
func (controller *Auth) safeRedilectAccount(sendMe string) {
	var safeAddress string
	safeAddress = controller.getSafeURL(sendMe)
	controller.Redirect(safeAddress, 302)
}

func (controller *Auth) getSafeURL(urlToTest string) string {
	var safeURL string
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

func (controller *Auth) isSafeRedirection(urlToTest string) bool {
	var isSafe bool
	isSafe = strings.HasPrefix(urlToTest, "/")
	return isSafe
}

func (controller *Auth) distroySession() {
	controller.DelSession("session")
	controller.DestroySession()
}
