package general

import adapter "github.com/cristian-sima/Wisply/models/adapter"

// MessageController encapsulates the operations for showing messages
type Message struct {
	Controller
}

// DisplaySimpleError shows an simple error message (string)
func (controller *Message) DisplaySimpleError(msg string) {
	err := adapter.WisplyError{
		Message: msg,
	}
	controller.DisplayError(err)
}

// DisplayError shows a complex error message
// (ussually after the validation of fields)
func (controller *Message) DisplayError(err adapter.WisplyError) {
	content := err.GetMessage()
	if len(err.Data) != 0 {
		controller.Data["validationFailed"] = true
		controller.Data["validationErrors"] = err.Data
	}
	controller.displayMessage("error", content)
}

// DisplaySuccessMessage displays a success message
// It provides a link to go back
func (controller *Message) DisplaySuccessMessage(content string, backLink string) {
	controller.Data["backLink"] = backLink
	controller.displayMessage("success", content)
}

// displayMessage is used by DisplayError and DisplaySuccessMessage
func (controller *Message) displayMessage(typeOfMessage string, content string) {
	controller.Data["messageContent"] = content
	controller.Data["displayMessage"] = true
	filename := "site/general/message/" + typeOfMessage
	controller.TplNames = filename + ".tpl"
	controller.Layout = "site/message-layout.tpl"
}
