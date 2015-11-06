package analyse

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InstitutionAnalyser analyse the data for an institution
type InstitutionAnalyser struct {
	parent        Analyser
	institution   repository.Institution
	moduleBuffer  *database.SQLBuffer
	programBuffer *database.SQLBuffer
}

// Start starts the process
func (analyser *InstitutionAnalyser) Start() {
	analyser.performModules()
	analyser.performPrograms()
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

func (analyser *InstitutionAnalyser) insertModuleData(moduleAnalyser ModuleAnalyser) {

	columns := "`analyse`, `module`, `keywords`, `formats`, `description`"
	tableName := "digest_module"
	analyser.moduleBuffer = database.NewSQLBuffer(tableName, columns)
	analyser.moduleBuffer.ChangeLimit(30)

	d1 := moduleAnalyser.GetKeywordsDigest()
	d2 := moduleAnalyser.GetFormatsDigest()
	d3 := moduleAnalyser.GetDescriptionDigest()
	analyser.moduleBuffer.AddRow(analyser.parent.id, moduleAnalyser.GetModule().GetID(), d1.GetPlainJSON(), d2.GetPlainJSON(), d3.GetPlainJSON())
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
