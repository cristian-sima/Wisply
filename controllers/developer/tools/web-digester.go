package tools

// WebDigester manages the operations for webDigester
type WebDigester struct {
	Controller
}

// Display shows a table with all the tools
func (controller *WebDigester) Display() {
	controller.GenerateXSRF()
	controller.SetCustomTitle("Admin - Tools")
	controller.LoadTemplate("web-digester")
}

// Work performs the digests and shows the results
func (controller *WebDigester) Work() {
	if !controller.IsCaptchaValid("tools") {
		controller.DisplaySimpleError("Please enter a valid code!")
	} else {
		controller.work()
	}
}

func (controller *WebDigester) work() {
	// controller.GenerateXSRF()
	// text := strings.TrimSpace(controller.GetString("digester-text"))
	// if len(text) > maxLenText {
	// 	text = text[0:10000]
	// }
	// list := analyse.NewWebDigester(text)
	// list.SortByCounter("DESC")
	// result := word.NewGrammarFilter(&list.Digester).GetData()
	// controller.Data["originalText"] = text
	// controller.Data["processed"] = result
	// controller.LoadTemplate("web-digester")
}
