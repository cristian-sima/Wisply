package tools

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/analyse/word"
)

// Digester manages the operations for digester
type Digester struct {
	Controller
}

// Display shows a table with all the tools
func (controller *Digester) Display() {
	controller.GenerateXSRF()
	controller.SetCustomTitle("Admin - Tools")
	controller.LoadTemplate("digester")
}

// Work performs the digests and shows the results
func (controller *Digester) Work() {
	controller.GenerateXSRF()
	text := strings.TrimSpace(controller.GetString("digester-text"))
	if len(text) > maxLenText {
		text = text[0:50000]
	}
	list := word.NewOccurencesList(text)
	list.SortByCounter("DESC")
	result := word.NewGrammarFilter(&list).GetData()
	controller.Data["originalText"] = text
	controller.Data["processed"] = result
	controller.LoadTemplate("digester")
}
