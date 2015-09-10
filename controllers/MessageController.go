package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type MessageController struct {
	beego.Controller
}

func (controller *MessageController) DisplayErrorMessage(errorMessage string) {

	content := errorMessage
	controller.DisplayMessage("error", content)
}

func (controller *MessageController) DisplayErrorMessages(errors map[string][]string) {
	var (
		number  int    = len(errors)
		message string = getMessage(number)
	)
	content := "Your request was not successful. " + message
	controller.Data["validationFailed"] = true
	controller.Data["validationErrors"] = errors
	controller.DisplayMessage("error", content)
}

func getMessage(number int) string {
	problemsMessage := ""
	if number == 1 {
		problemsMessage = "one field"
	} else {
		problemsMessage = strconv.Itoa(number) + " fields"
	}
	return "There were problems with " + problemsMessage + ":"
}

func (controller *MessageController) DisplaySuccessMessage(content string, backLink string) {
	controller.Data["backLink"] = backLink
	controller.DisplayMessage("success", content)
}

func (controller *MessageController) DisplayMessage(typeOfMessage string, content string) {
	controller.Data["messageContent"] = content
	controller.Data["displayMessage"] = true
	controller.TplNames = "general/message/" + typeOfMessage + ".tpl"
	controller.Layout = "general/message.tpl"
}
