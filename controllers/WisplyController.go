package controllers

import (
	"github.com/astaxie/beego"
)

type WisplyController struct {
	beego.Controller
}

func (c *WisplyController) DisplayErrorMessage (errors map[string][]string) {
    content := "The request was not successful. There were problems with the fields."
    c.Data["validationFailed"] = true
    c.Data["validationErrors"] = errors
    c.DisplayMessage("error", content)
}

func (c *WisplyController) DisplaySuccessMessage (content string, backLink string) {
	c.Data["backLink"] = backLink
    c.DisplayMessage("success", content)
}

func (c *WisplyController) DisplayMessage (typeOfMessage string, content string) {

    c.Data["messageContent"] = content
	c.Data["displayMessage"] = true
	c.TplNames = "general/message/" + typeOfMessage + ".tpl"
    c.Layout = "general/message.tpl"
}
