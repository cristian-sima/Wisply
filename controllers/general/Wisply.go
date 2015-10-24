// Package general is not used on any page, but it is inherited by the other controllers
// Contains basic functions such as generating XSRF token or connecting the account
package general

import (
	"html/template"
	"time"

	"github.com/cristian-sima/Wisply/models/auth"
	"github.com/cristian-sima/Wisply/models/curriculum"
)

// WisplyController inherits the MessageController
// Its role is to maintain the connection of the account
type WisplyController struct {
	MessageController
	AccountConnected bool
	Account          *auth.Account
}

// GenerateXSRF generates and sends to template the XSRF code
func (controller *WisplyController) GenerateXSRF() {
	code := controller.XsrfFormHtml()
	controller.Data["xsrf_input"] = template.HTML(code)
}

// Prepare checks the state of connection and inits the database
func (controller *WisplyController) Prepare() {
	controller.initState()
}

func (controller *WisplyController) initState() {
	session := controller.GetSession("account-id")
	if session != nil {
		id := (session).(string)
		controller.initConnectedState(id)
	} else {
		controller.checkConnectionCookie()
	}
	controller.loadPrograms()
}

func (controller *WisplyController) loadPrograms() {
	controller.Data["programs"] = curriculum.GetAllPrograms()
}

func (controller *WisplyController) checkConnectionCookie() {
	cookieName := auth.Settings["cookieName"].(string)
	cookie := controller.Ctx.GetCookie(cookieName)
	if cookie != "" {
		idUser, err := auth.ReconnectUsingCookie(cookie)
		if err == nil {
			controller.initConnectedState(idUser)
		} else {
			controller.DeleteConnectionCookie()
			controller.initDisconnectedState()
		}
	} else {
		controller.initDisconnectedState()
	}
}

// DeleteConnectionCookie deletes the cookies
func (controller *WisplyController) DeleteConnectionCookie() {
	cookieName := auth.Settings["cookieName"].(string)
	cookiePath := auth.Settings["cookiePath"].(string)
	cookie := controller.Ctx.GetCookie(cookieName)
	if cookie != "" {
		controller.Ctx.SetCookie(cookieName, "", -1, cookiePath)
	}
}

func (controller *WisplyController) initDisconnectedState() {
	controller.AccountConnected = false
	controller.Data["accountDisconnected"] = true
}

func (controller *WisplyController) initConnectedState(id string) {
	account, _ := auth.NewAccount(id)
	controller.Account = account
	controller.AccountConnected = true
	controller.Data["accountConnected"] = true
	controller.Data["currentAccount"] = controller.Account
}

// IsAccountConnected checks if there is any account connected
func (controller *WisplyController) IsAccountConnected() bool {
	return controller.AccountConnected
}

// IndicateLastModification shows in the footer the data when the page has been last modified
// Please use http://www.timestampgenerator.com/
func (controller *WisplyController) IndicateLastModification(timestamp int64) {
	formatedString := time.Unix(timestamp, 0).Format(time.ANSIC)
	controller.Data["indicateLastModification"] = true
	controller.Data["lastModification"] = formatedString
}

// SetCustomTitle sets a custom title for the page.
// If the function is not called, it sets the title "Wisply"
func (controller *WisplyController) SetCustomTitle(title string) {
	controller.Data["customTitle"] = title
}

// ShowBlankPage displays a blank page
func (controller *WisplyController) ShowBlankPage() {
	controller.Layout = "site/blank-layout.tpl"
	controller.TplNames = "site/blank.tpl"
}

// RedirectToLoginPage redirects the account to the login page
func (controller *WisplyController) RedirectToLoginPage() {

	loginPath := "/auth/login"
	addressParameter := "sendMe"

	currentPage := controller.Ctx.Request.URL.Path
	controller.Redirect(loginPath+"?"+addressParameter+"="+currentPage, 302)
}
