package analyse

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/analyse/word"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/education"
	"github.com/cristian-sima/Wisply/models/repository"
)

// SubjectAnalyser processes the data about a subject
type SubjectAnalyser struct {
	id          int
	parent      Analyser
	subject     education.Subject
	keywords    *word.Digester
	formats     *word.Digester
	description *word.Digester
	parentID    string
}

// GetKeywordsDigest returns the digest for the keywords
func (analyser SubjectAnalyser) GetKeywordsDigest() *word.Digester {
	return analyser.keywords
}

// GetGeneral returns the combination of both the digesters
func (analyser SubjectAnalyser) GetGeneral() *word.Digester {
	return analyser.keywords.Combine(analyser.description)
}

// GetParent returns the parent of the subject
func (analyser SubjectAnalyser) GetParent() Analyser {
	return NewAnalyser(analyser.parentID)
}

// GetFormatsDigest returns the digest for the formats
func (analyser SubjectAnalyser) GetFormatsDigest() *word.Digester {
	return analyser.formats
}

// GetDescriptionDigest returns the digest for the description
func (analyser SubjectAnalyser) GetDescriptionDigest() *word.Digester {
	return analyser.description
}

// GetSubject returns the subject of the analyser
func (analyser SubjectAnalyser) GetSubject() education.Subject {
	return analyser.subject
}

// let's start the magic
func (analyser SubjectAnalyser) start() {

	subject := analyser.GetSubject()
	exists := false

	institutions := repository.GetAllInstitutions()

	ownDescription := analyser.getDescription()

	keywords := word.NewDigester("")
	formats := word.NewDigester("")
	description := word.NewDigester("")

	for _, institution := range institutions {
		exists = true
		d1, d2, d3 := analyser.getDigesters(institution)
		keywords = keywords.Combine(d1)
		formats = formats.Combine(d2)
		description = description.Combine(d3)
	}
	if exists || len(ownDescription.GetData()) != 0 {

		description = description.Combine(ownDescription)

		description.SortByCounter("DESC")
		formats.SortByCounter("DESC")
		keywords.SortByCounter("DESC")

		columns := "`subject`, `description`, `formats`, `keywords`, `analyse`"
		table := "digest_subject"

		buffer := database.NewSQLBuffer(table, columns)
		buffer.AddRow(subject.GetID(), description.GetPlainJSON(), formats.GetPlainJSON(), keywords.GetPlainJSON(), analyser.parent.id)

		err := buffer.Exec()

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (analyser SubjectAnalyser) getDigesters(institution repository.Institution) (*word.Digester, *word.Digester, *word.Digester) {
	subject := analyser.subject
	fieldList := "`keywords`, `formats`, `description`"
	sql := "SELECT " + fieldList + " FROM `digest_institution` WHERE institution=? AND analyse=? AND subject = ? LIMIT 0,1"
	rows, err := database.Connection.Query(sql, institution.ID, analyser.parent.id, subject.GetID())

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

func (analyser SubjectAnalyser) getDescription() *word.Digester {
	subject := analyser.subject
	kas := subject.GetKAs()
	buffer := ""

	for _, ka := range kas {
		buffer += " " + ka.GetContent()
		buffer += " " + ka.GetTitle()
	}

	definitions := subject.GetDefinitions()

	for _, definition := range definitions {
		buffer += " " + definition.GetContent()
	}

	return analyser.digest(buffer)
}

// func (analyser SubjectAnalyser) getDigesters(module repository.Module) (*word.Digester, *word.Digester, *word.Digester) {
//
// 	fieldList := "`keywords`, `formats`, `description`"
// 	sql := "SELECT " + fieldList + " FROM `digest_module` WHERE module=? AND analyse=? LIMIT 0,1"
// 	rows, err := database.Connection.Query(sql, module.GetID(), analyser.parent.parent.id)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	keywords := word.NewDigester("")
// 	formats := word.NewDigester("")
// 	description := word.NewDigester("")
//
// 	for rows.Next() {
// 		var d1, d2, d3 string
// 		rows.Scan(&d1, &d2, &d3)
// 		keywords = word.NewDigesterFromJSON(d1)
// 		formats = word.NewDigesterFromJSON(d2)
// 		description = word.NewDigesterFromJSON(d3)
// 	}
// 	return keywords, formats, description
// }

// GetSubjectAnalysersBySubject gets all the subject analysers for a subject
func GetSubjectAnalysersBySubject(subjectID int) []SubjectAnalyser {
	var list []SubjectAnalyser
	fieldList := "`id`, `description`, `formats`, `keywords`, `analyse`"
	sql := "SELECT " + fieldList + " FROM `digest_subject` WHERE subject=? "
	rows, _ := database.Connection.Query(sql, subjectID)
	for rows.Next() {
		var d1, d2, d3 string
		analyser := SubjectAnalyser{}
		rows.Scan(&analyser.id, &d1, &d2, &d3, &analyser.parentID)
		analyser.description = word.NewDigesterFromJSON(d1)
		analyser.formats = word.NewDigesterFromJSON(d2)
		analyser.keywords = word.NewDigesterFromJSON(d3)
		list = append(list, analyser)
	}
	return list
}

func (analyser SubjectAnalyser) digest(text string) *word.Digester {
	digester := word.NewDigester(text)
	return word.NewGrammarFilter(digester).GetData()
}
