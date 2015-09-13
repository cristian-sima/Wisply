package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/cristian-sima/Wisply/models/adapter"
)

type MessageController struct {
	beego.Controller
}

func (controller *MessageController) DisplaySimpleError(msg string) {
	err := WisplyError{
		Message: msg,
	}
	controller.DisplayError(err)
}

func (controller *MessageController) DisplayError(err WisplyError) {
	content := err.GetMessage()
	if len(err.Data) != 0 {
		controller.Data["validationFailed"] = true
		controller.Data["validationErrors"] = err.Data
	}
	controller.DisplayMessage("error", content)
}

func (controller *MessageController) DisplaySuccessMessage(content string, backLink string) {
	controller.Data["backLink"] = backLink
	controller.DisplayMessage("success", content)
}

func (controller *MessageController) DisplayMessage(typeOfMessage string, content string) {
	controller.Data["messageContent"] = content
	controller.Data["displayMessage"] = true
	controller.TplNames = "site/message/" + typeOfMessage + ".tpl"
	controller.Layout = "site/message.tpl"
}
