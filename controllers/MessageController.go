package controllers

import (
  "github.com/astaxie/beego"
  "strconv"
)

type MessageController struct {
  beego.Controller
}

func (c *MessageController) DisplayErrorMessage (errors map[string][]string) {
  var (
    number int = len(errors)
    message string = getMessage(number)
  )
  content := "Your request was not successful. " + message
  c.Data["validationFailed"] = true
  c.Data["validationErrors"] = errors
  c.DisplayMessage("error", content)
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

func (c *MessageController) DisplaySuccessMessage (content string, backLink string) {
  c.Data["backLink"] = backLink
  c.DisplayMessage("success", content)
}

func (c *MessageController) DisplayMessage (typeOfMessage string, content string) {
  c.Data["messageContent"] = content
  c.Data["displayMessage"] = true
  c.TplNames = "general/message/" + typeOfMessage + ".tpl"
  c.Layout = "general/message.tpl"
}
