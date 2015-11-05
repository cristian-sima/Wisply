package analyse

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/analyse/word"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// ModuleAnalyser processes the data about the module
type ModuleAnalyser struct {
	id          int
	parent      InstitutionAnalyser
	module      repository.Module
	keywords    *word.Digester
	formats     *word.Digester
	description *word.Digester
	parentID    string
}

// GetKeywordsDigest returns the digest for the keywords
func (analyser ModuleAnalyser) GetKeywordsDigest() *word.Digester {
	return analyser.keywords
}

// GetGeneral returns the combination of both the digesters
func (analyser ModuleAnalyser) GetGeneral() *word.Digester {
	return analyser.keywords.Combine(analyser.description)
}

// GetParent returns the parent of the module
func (analyser ModuleAnalyser) GetParent() Analyser {
	return NewAnalyser(analyser.parentID)
}

// GetFormatsDigest returns the digest for the formats
func (analyser ModuleAnalyser) GetFormatsDigest() *word.Digester {
	return analyser.formats
}

// GetDescriptionDigest returns the digest for the description
func (analyser ModuleAnalyser) GetDescriptionDigest() *word.Digester {
	return analyser.description
}

// GetModule returns the module of the analyser
func (analyser ModuleAnalyser) GetModule() repository.Module {
	return analyser.module
}

// it starts the magic
func (analyser ModuleAnalyser) start() {
	identifiers := analyser.getIdentifiers()
	d1, d2 := analyser.getDigests(identifiers)
	analyser.keywords = d1
	analyser.formats = d2
	analyser.description = analyser.getDescription()

	analyser.parent.insertModuleData(analyser)
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

// GetModuleAnalysersByModule gets all the module analysers for a module
func GetModuleAnalysersByModule(moduleID int) []ModuleAnalyser {
	var list []ModuleAnalyser
	fieldList := "`id`, `description`, `formats`, `keywords`, `analyse`"
	sql := "SELECT " + fieldList + " FROM `digest_module` WHERE module=? "
	rows, _ := database.Connection.Query(sql, moduleID)
	for rows.Next() {
		var d1, d2, d3 string
		analyser := ModuleAnalyser{}
		rows.Scan(&analyser.id, &d1, &d2, &d3, &analyser.parentID)
		analyser.description = word.NewDigesterFromJSON(d1)
		analyser.formats = word.NewDigesterFromJSON(d2)
		analyser.keywords = word.NewDigesterFromJSON(d3)
		fmt.Println(analyser.module)
		list = append(list, analyser)
	}
	return list
}

//
// // NewModuleAnalyser gets the ModuleAnalyser
// func NewModuleAnalyser(id string) ModuleAnalyser {
// 	analyser := ModuleAnalyser{}
// 	fieldList := "`id`, `module`, `description`, `formats`, `keywords`, `analyse`"
// 	sql := "SELECT " + fieldList + " FROM `digest_module` WHERE id=? "
// 	query, err := database.Connection.Prepare(sql)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	query.QueryRow(id).Scan(&analyser.id, &analyser.module, &analyser.description, &analyser.formats, &analyser.keywords, &analyser.parentID)
// 	return analyser
// }

func (analyser ModuleAnalyser) digest(text string) *word.Digester {
	digester := word.NewDigester(text)
	digester.SortByCounter("DESC")
	result := word.NewGrammarFilter(&digester).GetData()
	result.RemoveOccurence(analyser.module.GetCode())
	return result
}
