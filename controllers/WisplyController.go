package controllers

import (
	. "github.com/cristian-sima/Wisply/models/auth"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"html/template"
)

type WisplyController struct {
	MessageController
	AccountConnected bool
	Account          *Account
	Model            WisplyModel
}

func (this *WisplyController) GenerateXsrf() {
	this.Data["xsrf_input"] = template.HTML(this.XsrfFormHtml())
}

func (controller *WisplyController) Prepare() {
	controller.initState()
	InitDatabase()
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
	cookieName := Settings["cookieName"].(string)
	cookie := controller.Ctx.GetCookie(cookieName)
	if cookie != "" {
		idUser, err := ReConnect(cookie)
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
	cookieName := Settings["cookieName"].(string)
	cookiePath := Settings["cookiePath"].(string)
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
	account, _ := NewAccount(id)
	controller.Account = account
	controller.AccountConnected = true
	controller.Data["accountConnected"] = true
	controller.Data["currentAccount"] = controller.Account
}
