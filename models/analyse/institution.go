package analyse

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/analyse/word"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// InstitutionAnalyser analyse the data for an institution
type InstitutionAnalyser struct {
	parent       Analyser
	institution  repository.Institution
	moduleBuffer *database.SQLBuffer
}

// Start starts the process
func (analyser *InstitutionAnalyser) Start() {

	// get all modules
	modules := analyser.institution.GetModules()

	for _, module := range modules {
		child := analyser.CreateModuleAnalyser(module)
		child.start()
	}
}

func (analyser *InstitutionAnalyser) insertModuleData(module repository.Module, digesters map[string]*word.Digester) {

	columns := "`analyse`, `module`, `keywords`, `formats`, `description`"
	tableName := "digest_module"
	analyser.moduleBuffer = database.NewSQLBuffer(tableName, columns)
	analyser.moduleBuffer.ChangeLimit(30)

	d1 := digesters["keywords"]
	d2 := digesters["formats"]
	d3 := digesters["description"]
	analyser.moduleBuffer.AddRow(analyser.parent.id, module.GetID(), d1.GetJSON(), d2.GetJSON(), d3.GetJSON())
	err := analyser.moduleBuffer.Exec()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Finished")
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
