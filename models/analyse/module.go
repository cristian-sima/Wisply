package analyse

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/analyse/word"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// ModuleAnalyser processes the data about the module
type ModuleAnalyser struct {
	parent InstitutionAnalyser
	module repository.Module
}

// it starts the magic
func (analyser ModuleAnalyser) start() {
	module := analyser.module
	digesters := make(map[string]*word.Digester)
	identifiers := analyser.getIdentifiers()
	d1, d2 := analyser.getDigests(identifiers)
	digesters["keywords"] = d1
	digesters["formats"] = d2
	digesters["description"] = analyser.getDescription()

	analyser.parent.insertModuleData(module, digesters)
}

func (analyser ModuleAnalyser) getDescription() *word.Digester {
	module := analyser.module
	return analyser.digest(module.GetContent() + " " + module.GetTitle())
}

func (analyser ModuleAnalyser) getDigests(identifiers []string) (*word.Digester, *word.Digester) {

	keywordsBuffer, formatsBuffer := "", ""

	for _, identifier := range identifiers {

		sql := "SELECT `resource_key`, `value` FROM `resource_key` WHERE `resource` = ?"
		rows, err := database.Connection.Query(sql, identifier)

		if err != nil {
			fmt.Println("Error #1 module analysing data")
			fmt.Println(err)
		}

		for rows.Next() {
			var key, value string
			rows.Scan(&key, &value)
			switch key {
			case "title", "identifier":
				keywordsBuffer += " " + value
				break
			case "format":
				formatsBuffer += " " + value
				break
			}
		}
	}
	d1 := analyser.digest(keywordsBuffer)
	d2 := analyser.digest(formatsBuffer)
	return d1, d2
}

func (analyser ModuleAnalyser) getIdentifiers() []string {
	identifiers := []string{}
	sql := "SELECT  DISTINCT `resource` FROM `resource_key` WHERE `value` LIKE ?"
	rows, err := database.Connection.Query(sql, "%"+analyser.module.GetCode()+"%")

	if err != nil {
		fmt.Println("Error #1 module analysing data")
		fmt.Println(err)
	}

	for rows.Next() {
		identifier := ""
		rows.Scan(&identifier)
		identifiers = append(identifiers, identifier)
	}
	return identifiers
}

func (analyser ModuleAnalyser) digest(text string) *word.Digester {
	digester := word.NewDigester(text)
	digester.SortByCounter("DESC")
	result := word.NewGrammarFilter(&digester).GetData()
	result.RemoveOccurence(analyser.module.GetCode())
	return result
}
