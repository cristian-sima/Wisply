package analyse

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/analyse/word"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/education"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InstitutionAnalyser analyse the data for an institution
type InstitutionAnalyser struct {
	id            string
	parent        Analyser
	institution   repository.Institution
	moduleBuffer  *database.SQLBuffer
	programBuffer *database.SQLBuffer
	keywords      *word.Digester
	formats       *word.Digester
	description   *word.Digester
	parentID      string
	subject       string
}

// GetSubject returns the subject object
func (analyser InstitutionAnalyser) GetSubject() *education.Subject {
	subject, _ := education.NewSubject(analyser.subject)
	return subject
}

// GetKeywordsDigest returns the digest for the keywords
func (analyser InstitutionAnalyser) GetKeywordsDigest() *word.Digester {
	return analyser.keywords
}

// GetGeneral returns the combination of both the digesters
func (analyser InstitutionAnalyser) GetGeneral() *word.Digester {
	return analyser.keywords.Combine(analyser.description)
}

// GetParent returns the parent of the module
func (analyser InstitutionAnalyser) GetParent() Analyser {
	return NewAnalyser(analyser.parentID)
}

// GetFormatsDigest returns the digest for the formats
func (analyser InstitutionAnalyser) GetFormatsDigest() *word.Digester {
	return analyser.formats
}

// GetDescriptionDigest returns the digest for the description
func (analyser InstitutionAnalyser) GetDescriptionDigest() *word.Digester {
	return analyser.description
}

// Start starts the process
func (analyser *InstitutionAnalyser) Start() {
	analyser.performModules()
	analyser.performPrograms()
	analyser.perform()
}

func (analyser *InstitutionAnalyser) perform() {

	institution := analyser.institution

	subjects := education.GetAllSubjects()

	for _, subject := range subjects {
		exists := false
		programs := institution.GetProgramsBySubjectID(subject.GetID())

		keywords := word.NewDigester("")
		formats := word.NewDigester("")
		description := word.NewDigester("")

		for _, program := range programs {
			exists = true
			fieldList := "`description`, `formats`, `keywords`"
			sql := "SELECT " + fieldList + " FROM `digest_program` WHERE program=? AND analyse=? "
			rows, _ := database.Connection.Query(sql, program.GetID(), analyser.parent.id)
			for rows.Next() {
				var s1, s2, s3 string
				rows.Scan(&s1, &s2, &s3)
				description = description.Combine(word.NewDigesterFromJSON(s1))
				formats = formats.Combine(word.NewDigesterFromJSON(s2))
				keywords = keywords.Combine(word.NewDigesterFromJSON(s3))
			}
		}
		if exists {
			keywords.SortByCounter("DESC")
			description.SortByCounter("DESC")
			formats.SortByCounter("DESC")

			columns := "`institution`, `subject`, `description`, `formats`, `keywords`, `analyse`"
			table := "digest_institution"

			buffer := database.NewSQLBuffer(table, columns)
			buffer.AddRow(institution.ID, subject.GetID(), description.GetPlainJSON(), formats.GetPlainJSON(), keywords.GetPlainJSON(), analyser.parent.id)

			err := buffer.Exec()

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (analyser *InstitutionAnalyser) performModules() {
	modules := analyser.institution.GetModules()
	for _, module := range modules {
		child := analyser.CreateModuleAnalyser(module)
		child.start()
	}
}

func (analyser *InstitutionAnalyser) performPrograms() {
	programs := analyser.institution.GetPrograms()
	for _, program := range programs {
		child := analyser.CreateProgramAnalyser(program)
		child.start()
	}
}

func (analyser *InstitutionAnalyser) insertModuleData(InstitutionAnalyser ModuleAnalyser) {

	columns := "`analyse`, `module`, `keywords`, `formats`, `description`"
	tableName := "digest_module"
	analyser.moduleBuffer = database.NewSQLBuffer(tableName, columns)
	analyser.moduleBuffer.ChangeLimit(30)

	d1 := InstitutionAnalyser.GetKeywordsDigest()
	d2 := InstitutionAnalyser.GetFormatsDigest()
	d3 := InstitutionAnalyser.GetDescriptionDigest()
	analyser.moduleBuffer.AddRow(analyser.parent.id, InstitutionAnalyser.GetModule().GetID(), d1.GetPlainJSON(), d2.GetPlainJSON(), d3.GetPlainJSON())
	err := analyser.moduleBuffer.Exec()
	if err != nil {
		fmt.Println(err)
	}
}

func (analyser *InstitutionAnalyser) insertProgramData(programAnalyser ProgramAnalyser) {

	columns := "`analyse`, `program`, `keywords`, `formats`, `description`"
	tableName := "digest_program"
	analyser.programBuffer = database.NewSQLBuffer(tableName, columns)
	analyser.programBuffer.ChangeLimit(15)

	d1 := programAnalyser.GetKeywordsDigest()
	d2 := programAnalyser.GetFormatsDigest()
	d3 := programAnalyser.GetDescriptionDigest()
	analyser.programBuffer.AddRow(analyser.parent.id, programAnalyser.GetProgram().GetID(), d1.GetPlainJSON(), d2.GetPlainJSON(), d3.GetPlainJSON())
	err := analyser.programBuffer.Exec()
	if err != nil {
		fmt.Println(err)
	}
}

// GetInstitutionAnalysersByInstitution gets all the analyses for the institution
func GetInstitutionAnalysersByInstitution(institutionID int) []InstitutionAnalyser {
	var list []InstitutionAnalyser
	fieldList := "`id`, `description`, `formats`, `keywords`, `analyse`, `subject`"
	sql := "SELECT " + fieldList + " FROM `digest_institution` WHERE institution=? ORDER by analyse DESC"
	rows, err := database.Connection.Query(sql, institutionID)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var d1, d2, d3 string
		analyser := InstitutionAnalyser{}
		rows.Scan(&analyser.id, &d1, &d2, &d3, &analyser.parentID, &analyser.subject)
		analyser.description = word.NewDigesterFromJSON(d1)
		analyser.formats = word.NewDigesterFromJSON(d2)
		analyser.keywords = word.NewDigesterFromJSON(d3)
		fmt.Println(analyser.GetParent())
		list = append(list, analyser)
	}
	return list
}

// CreateModuleAnalyser creates a new module analyser
func (analyser InstitutionAnalyser) CreateModuleAnalyser(module repository.Module) ModuleAnalyser {
	moduleAnalyser := ModuleAnalyser{
		parent: analyser,
		module: module,
	}
	return moduleAnalyser
}

// CreateProgramAnalyser creates a new module analyser
func (analyser InstitutionAnalyser) CreateProgramAnalyser(program repository.Program) ProgramAnalyser {
	programAnalyser := ProgramAnalyser{
		parent:  analyser,
		program: program,
	}
	return programAnalyser
}
