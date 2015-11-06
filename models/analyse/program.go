package analyse

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/analyse/word"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// ProgramAnalyser processes the data about a program
type ProgramAnalyser struct {
	id          int
	parent      InstitutionAnalyser
	program     repository.Program
	keywords    *word.Digester
	formats     *word.Digester
	description *word.Digester
	parentID    string
}

// GetKeywordsDigest returns the digest for the keywords
func (analyser ProgramAnalyser) GetKeywordsDigest() *word.Digester {
	return analyser.keywords
}

// GetGeneral returns the combination of both the digesters
func (analyser ProgramAnalyser) GetGeneral() *word.Digester {
	return analyser.keywords.Combine(analyser.description)
}

// GetParent returns the parent of the program
func (analyser ProgramAnalyser) GetParent() Analyser {
	return NewAnalyser(analyser.parentID)
}

// GetFormatsDigest returns the digest for the formats
func (analyser ProgramAnalyser) GetFormatsDigest() *word.Digester {
	return analyser.formats
}

// GetDescriptionDigest returns the digest for the description
func (analyser ProgramAnalyser) GetDescriptionDigest() *word.Digester {
	return analyser.description
}

// GetProgram returns the program of the analyser
func (analyser ProgramAnalyser) GetProgram() repository.Program {
	return analyser.program
}

// let's start the magic
func (analyser ProgramAnalyser) start() {

	program := analyser.GetProgram()
	modules := program.GetModules()

	ownDescription := analyser.getDescription()

	keywords := word.NewDigester("")
	formats := word.NewDigester("")
	description := word.NewDigester("")

	for _, module := range modules {
		d1, d2, d3 := analyser.getDigesters(module)
		keywords = keywords.Combine(d1)
		formats = formats.Combine(d2)
		description = description.Combine(d3)
	}

	description = description.Combine(ownDescription)

	analyser.keywords = keywords
	analyser.formats = formats
	analyser.description = description

	analyser.parent.insertProgramData(analyser)
}

func (analyser ProgramAnalyser) getDescription() *word.Digester {
	program := analyser.program
	return analyser.digest(program.GetContent() + " " + program.GetTitle())
}

func (analyser ProgramAnalyser) getDigesters(module repository.Module) (*word.Digester, *word.Digester, *word.Digester) {

	fieldList := "`keywords`, `formats`, `description`"
	sql := "SELECT " + fieldList + " FROM `digest_module` WHERE module=? AND analyse=? LIMIT 0,1"
	rows, err := database.Connection.Query(sql, module.GetID(), analyser.parent.parent.id)

	if err != nil {
		panic(err)
	}

	keywords := word.NewDigester("")
	formats := word.NewDigester("")
	description := word.NewDigester("")

	for rows.Next() {
		var d1, d2, d3 string
		rows.Scan(&d1, &d2, &d3)
		keywords = word.NewDigesterFromJSON(d1)
		formats = word.NewDigesterFromJSON(d2)
		description = word.NewDigesterFromJSON(d3)
	}
	return keywords, formats, description
}

// GetProgramAnalysersByProgram gets all the program analysers for a program
func GetProgramAnalysersByProgram(programID int) []ProgramAnalyser {
	var list []ProgramAnalyser
	fieldList := "`id`, `description`, `formats`, `keywords`, `analyse`"
	sql := "SELECT " + fieldList + " FROM `digest_program` WHERE program=? "
	rows, _ := database.Connection.Query(sql, programID)
	for rows.Next() {
		var d1, d2, d3 string
		analyser := ProgramAnalyser{}
		rows.Scan(&analyser.id, &d1, &d2, &d3, &analyser.parentID)
		analyser.description = word.NewDigesterFromJSON(d1)
		analyser.formats = word.NewDigesterFromJSON(d2)
		analyser.keywords = word.NewDigesterFromJSON(d3)
		fmt.Println(analyser.program)
		list = append(list, analyser)
	}
	return list
}

func (analyser ProgramAnalyser) digest(text string) *word.Digester {
	digester := word.NewDigester(text)
	digester.SortByCounter("DESC")
	result := word.NewGrammarFilter(digester).GetData()
	result.RemoveOccurence(analyser.program.GetCode())
	return result
}
