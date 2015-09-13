package controllers

import (
	. "github.com/cristian-sima/Wisply/models/auth"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"html/template"
	"fmt"
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
		fmt.Println("exista session cookie")
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
		fmt.Println("Coookie este " + cookie)
		fmt.Println("try to reconnect")
		idUser, err := ReConnect(cookie)
		if err == nil {
			fmt.Println("a mers")
			controller.initConnectedState(idUser)
		} else {
			fmt.Println("nu am putut pentru ca:")
			fmt.Println(err)
			controller.deleteConnectionCookie()
			controller.initDisconnectedState()
		}
	} else {
		controller.initDisconnectedState()
	}
}

func (controller *WisplyController) deleteConnectionCookie() {
	fmt.Println("sterg connection cookie")
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
