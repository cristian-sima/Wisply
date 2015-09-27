package general

import (
	"html/template"
	"time"

	auth "github.com/cristian-sima/Wisply/models/auth"
	wisply "github.com/cristian-sima/Wisply/models/wisply"
)

// WisplyController inherits the MessageController
// Its role is to maintain the connection of the account
type WisplyController struct {
	MessageController
	AccountConnected bool
	Account          *auth.Account
	Model            wisply.Model
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
}

func (controller *WisplyController) checkConnectionCookie() {
	cookieName := auth.Settings["cookieName"].(string)
	cookie := controller.Ctx.GetCookie(cookieName)
	if cookie != "" {
		idUser, err := auth.ReConnect(cookie)
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

// IndicateLastModification shows in the footer the data when the page has been last modified
// Please use http://www.timestampgenerator.com/
func (controller *WisplyController) IndicateLastModification(timestamp int64) {
	formatedString := time.Unix(timestamp, 0).Format(time.ANSIC)
	controller.Data["indicateLastModification"] = true
	controller.Data["lastModification"] = formatedString
}

// SetCustomTitle sets a custom title for the page. In case the function is not called, it sets the title "Wisply"
func (controller *WisplyController) SetCustomTitle(title string) {
	controller.Data["customTitle"] = title
}
