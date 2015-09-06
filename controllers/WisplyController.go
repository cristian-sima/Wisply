package controllers

import (
  "html/template"
)

type WisplyController struct {
	MessageController
}

func (this *WisplyController) GenerateXsrf(){
    this.Data["xsrf_input"]= template.HTML(this.XsrfFormHtml())
}
