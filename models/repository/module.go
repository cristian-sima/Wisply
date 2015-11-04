package repository

import (
	"errors"
	"fmt"

	"github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
)

// Module represents a module within a program of study
type Module struct {
	id      int
	title   string
	content string
	code    string
	module  string
	credits float64
	program int
	year    string
}

// GetID returns the ID of the module
func (module Module) GetID() int {
	return module.id
}

// GetTitle returns title
func (module Module) GetTitle() string {
	return module.title
}

// GetContent returns the content of the module
func (module Module) GetContent() string {
	return module.content
}

// GetCode returns the code
func (module Module) GetCode() string {
	return module.code
}

// GetCredits returns the academic credits
func (module Module) GetCredits(category string) float64 {
	// It is raporting to credits
	points := 0.0
	switch category {
	case "CATS":
		points = module.credits
		break
	case "ECTS":
		points = module.credits / 2
		break
	case "US":
		points = module.credits / 4
		break
	default:
		points = module.credits
		break
	}
	return points
}

// GetYear returns the year of the module
func (module Module) GetYear() string {
	return module.year
}

// Delete removes the module from database
func (module *Module) Delete() error {
	sql := "DELETE FROM `institution_module` WHERE `id`=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = query.Exec(module.id)
	fmt.Println(err)
	return err
}

// Modify changes the details of the module
func (module Module) Modify(details map[string]interface{}) (adapter.WisplyError, error) {
	problems := adapter.WisplyError{}
	result := hasValidModuleModifyDetails(details)
	if !result.IsValid {
		problems.Data = result
		return problems, errors.New("Problem with the fields")
	}
	setClause := "SET `title`=?, `content`=?, `code`=?, `program`=?, `credits`=?, `year`=?"
	whereClause := "WHERE `id`= ?"
	sql := "UPDATE `institution_module` " + setClause + " " + whereClause
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	title := details["module-title"].(string)
	content := details["module-content"].(string)
	code := details["module-code"].(string)
	program := details["module-program"].(int)
	credits := details["module-credits"].(string)
	year := details["module-year"].(string)
	query.Exec(title, content, code, program, credits, year, module.id)
	return problems, err
}

// NewModule creates a new module
func NewModule(ID string) (*Module, error) {
	module := &Module{}
	fieldList := "`id`, `title`, `content`, `code`, `program`, `credits`, `year`"
	sql := "SELECT " + fieldList + " FROM `institution_module` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return module, err
	}
	query.QueryRow(ID).Scan(&module.id, &module.title, &module.content, &module.code, &module.program, &module.credits, &module.year)
	return module, nil
}

// CreateModule adds a new module into database
func CreateModule(details map[string]interface{}) (adapter.WisplyError, error) {
	problems := adapter.WisplyError{}
	result := hasValidModuleModifyDetails(details)
	if !result.IsValid {
		problems.Data = result
		return problems, errors.New("Invalid details for the module")
	}
	fieldList := "`title`, `content`, `code`, `program`, `credits`, `year`"
	questions := "?, ?, ?, ?, ?, ?"
	sql := "INSERT INTO `institution_module` (" + fieldList + ") VALUES (" + questions + ")"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		return problems, err
	}
	title := details["module-title"].(string)
	content := details["module-content"].(string)
	code := details["module-code"].(string)
	program := details["module-program"].(int)
	credits := details["module-credits"].(string)
	year := details["module-year"].(string)
	_, err = query.Exec(title, content, code, program, credits, year)

	fmt.Println(err)
	return problems, err
}
