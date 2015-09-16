package controllers

import (
	beego "github.com/astaxie/beego"
	adapter "github.com/cristian-sima/Wisply/models/adapter"
)

// MessageController It encapsulates the operations for showing messages
type MessageController struct {
	beego.Controller
}

// DisplaySimpleError It shows an simple error message (string)
func (controller *MessageController) DisplaySimpleError(msg string) {
	err := adapter.WisplyError{
		Message: msg,
	}
	controller.DisplayError(err)
}

// DisplayError It shows a complex error message (ussually after the validation of fields)
func (controller *MessageController) DisplayError(err adapter.WisplyError) {
	content := err.GetMessage()
	if len(err.Data) != 0 {
		controller.Data["validationFailed"] = true
		controller.Data["validationErrors"] = err.Data
	}
	controller.displayMessage("error", content)
}

// DisplaySuccessMessage It displays a success message and provides a link to go back
func (controller *MessageController) DisplaySuccessMessage(content string, backLink string) {
	controller.Data["backLink"] = backLink
	controller.displayMessage("success", content)
}

// It is used by DisplayError and DisplaySuccessMessage
func (controller *MessageController) displayMessage(typeOfMessage string, content string) {
	controller.Data["messageContent"] = content
	controller.Data["displayMessage"] = true
	controller.TplNames = "site/message/" + typeOfMessage + ".tpl"
	controller.Layout = "site/message.tpl"
}
