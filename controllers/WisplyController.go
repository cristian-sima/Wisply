package controllers

import (
	"fmt"
	. "github.com/cristian-sima/Wisply/models/auth"
	. "github.com/cristian-sima/Wisply/models/wisply"
	"html/template"
)

type WisplyController struct {
	MessageController
	UserConnected bool
	User          User
	Model         WisplyModel
}

func (this *WisplyController) GenerateXsrf() {
	this.Data["xsrf_input"] = template.HTML(this.XsrfFormHtml())
}

func (controller *WisplyController) Prepare() {
	controller.updateUserConnection()
	InitDatabase()
}

func (controller *WisplyController) updateUserConnection() {
	session := controller.GetSession("user-id")
	if session == nil {
		controller.UserConnected = false
		controller.Data["userDisconnected"] = true
	} else {
		id := (session).(string)
		controller.User = NewUser(id)
		fmt.Println(controller.User)
		controller.UserConnected = true
		controller.Data["userConnected"] = true
		controller.Data["currentUser"] = controller.User
	}
}
