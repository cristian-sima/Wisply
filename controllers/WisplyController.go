package controllers

import (
	"fmt"
	. "github.com/cristian-sima/Wisply/models/auth"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"html/template"
)

type WisplyController struct {
	MessageController
	AccountConnected bool
	Account          Account
	Model            WisplyModel
}

func (this *WisplyController) GenerateXsrf() {
	this.Data["xsrf_input"] = template.HTML(this.XsrfFormHtml())
}

func (controller *WisplyController) Prepare() {
	controller.updateAccountConnection()
	InitDatabase()
}

func (controller *WisplyController) updateAccountConnection() {
	session := controller.GetSession("account-id")
	if session == nil {
		controller.AccountConnected = false
		controller.Data["accountDisconnected"] = true
	} else {
		id := (session).(string)
		controller.Account = NewAccount(id)
		fmt.Println(controller.Account)
		controller.AccountConnected = true
		controller.Data["accountConnected"] = true
		controller.Data["currentAccount"] = controller.Account
	}
}
