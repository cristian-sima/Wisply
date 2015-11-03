package analyse

import (
	"fmt"
	"net/url"

	"github.com/cristian-sima/Wisply/models/analyse/word"
)

// WebDigester is an Digester struct which is able to gather data by
// sending http requests and then collecting the source data
type WebDigester struct {
	word.Digester
}

// AnalyseText gets a text and finds the web occurences
func (digester *WebDigester) AnalyseText(text string) {
	digester.Digester = word.NewDigester(text)
	digester.perform()
}

func (digester *WebDigester) perform() {
	links := digester.getLinks()
	for _, link := range links {
		fmt.Println(link)
	}
}

func (digester *WebDigester) getLinks() []string {
	list := []string{}
	for _, occurence := range digester.GetData() {
		if isLink(occurence.GetWord()) {
			list = append(list, occurence.GetWord())
		}
	}
	return list
}

// NewWebDigester creates a new web digester
func NewWebDigester(text string) WebDigester {
	digester := WebDigester{}
	digester.AnalyseText(text)
	return digester
}

func isLink(word string) bool {
	_, err := url.ParseRequestURI(word)
	return err == nil
}
