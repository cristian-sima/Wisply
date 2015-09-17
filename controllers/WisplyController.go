package controllers

import (
	"fmt"
	"html/template"

	auth "github.com/cristian-sima/Wisply/models/auth"
	wisply "github.com/cristian-sima/Wisply/models/wisply"
)

// WisplyController It inherits the MessageController
// Its role is to maintain the connection of the account
type WisplyController struct {
	MessageController
	AccountConnected bool
	Account          *auth.Account
	Model            wisply.Model
}

// GenerateXSRF It generates and sends to template the XSRF code
func (controller *WisplyController) GenerateXSRF() {
	code := controller.XsrfFormHtml()
	controller.Data["xsrf_input"] = template.HTML(code)
}

// Prepare It checks the state of connection and inits the database
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
			controller.deleteConnectionCookie()
			controller.initDisconnectedState()
		}
	} else {
		controller.initDisconnectedState()
	}
}

func (controller *WisplyController) deleteConnectionCookie() {
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
	fmt.Println("aici")
	fmt.Println(account)
	controller.Account = account
	controller.AccountConnected = true
	controller.Data["accountConnected"] = true
	controller.Data["currentAccount"] = controller.Account
}
