package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/adapter"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/education"
)

// Program represents a program for a institution
type Program struct {
	id          int
	institution int
	title       string
	code        string
	year        string
	ucasCode    string
	level       string
	content     string
	subject     int
}

// GetID returns the ID of the program
func (program Program) GetID() int {
	return program.id
}

// GetSubject returns the subject area of the program of study
func (program Program) GetSubject() *education.Subject {
	subject, err := education.NewSubject(strconv.Itoa(program.subject))
	if err != nil {
		fmt.Println(err)
	}
	return subject
}

// GetContent returns the content of the program
func (program Program) GetContent() string {
	return program.content
}

// GetTitle returns title
func (program Program) GetTitle() string {
	return program.title
}

// GetCode returns the code
func (program Program) GetCode() string {
	return program.code
}

// GetYear returns the year
func (program Program) GetYear() string {
	return program.year
}

// GetUCASCode returns the UCAS code
func (program Program) GetUCASCode() string {
	return program.ucasCode
}

// GetLevel returns the level of the program
func (program Program) GetLevel() string {
	return program.level
}

// GetInstitution returns institution
func (program Program) GetInstitution() int {
	return program.institution
}

// Delete removes the program from database
func (program *Program) Delete() error {
	sql := "DELETE FROM `institution_program` WHERE id=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = query.Exec(program.id)
	return err
}

// AddModule adds a module to the program of study
func (program Program) AddModule(moduleID string) error {

	fieldList := "`program`, `module`"
	questions := "?, ?"
	sql := "INSERT INTO `institution_program_session` (" + fieldList + ") VALUES (" + questions + ")"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		return err
	}
	_, err = query.Exec(program.GetID(), moduleID)
	return err
}

// DeleteModule removes a module from a program of study
func (program *Program) DeleteModule(moduleID string) error {
	sql := "DELETE FROM `institution_program_session` WHERE `program`=? AND `module`=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = query.Exec(program.GetID(), moduleID)
	return err
}

// Modify changes the details of the program
func (program Program) Modify(details map[string]interface{}) (adapter.WisplyError, error) {
	problems := adapter.WisplyError{}
	result := hasValidProgramModifyDetails(details)
	if !result.IsValid {
		problems.Data = result
		return problems, errors.New("Problem with the fields")
	}
	setClause := "SET `title`=?, `code`=?, `year`=?, `ucas_code`=?, `level`=?, `content`=?, `subject`=? "
	whereClause := "WHERE `id`= ?"
	sql := "UPDATE `institution_program` " + setClause + " " + whereClause
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	title := details["program-title"].(string)
	code := details["program-code"].(string)
	year := details["program-year"].(string)
	ucasCode := details["program-ucas-code"].(string)
	level := details["program-level"].(string)
	content := details["program-content"].(string)
	subject := details["program-subject"].(string)
	query.Exec(title, code, year, ucasCode, level, content, subject, program.id)
	return problems, err
}

// GetModules returns the modules of the program
func (program Program) GetModules() []Module {
	var list []Module
	fieldList := "module.`id`, module.`title`, module.`content`, module.`code`, module.`credits`, module.`year`, module.`institution`"
	orderClause := "ORDER BY module.`year` ASC, module.`title` ASC"
	whereClause := "WHERE `program` = ?"
	join := "INNER JOIN `institution_program_session` AS session ON session.module = module.id"
	sql := "SELECT " + fieldList + " FROM `institution_module` AS module " + join + " " + whereClause + " " + orderClause
	rows, err := database.Connection.Query(sql, strconv.Itoa(program.GetID()))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		item := Module{}
		rows.Scan(&item.id, &item.title, &item.content, &item.code, &item.credits, &item.year, &item.institution)
		list = append(list, item)
	}
	return list
}

// NewProgram creates a new program
func NewProgram(ID string) (Program, error) {
	program := Program{}
	fieldList := "`id`, `institution`, `title`, `code`, `year`, `ucas_code`, `level`, `content`, `subject`"
	sql := "SELECT " + fieldList + " FROM `institution_program` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return program, err
	}
	query.QueryRow(ID).Scan(&program.id, &program.institution, &program.title, &program.code, &program.year, &program.ucasCode, &program.level, &program.content, &program.subject)
	return program, nil
}

// CreateProgram adds a new program into database
func CreateProgram(details map[string]interface{}) (adapter.WisplyError, error) {
	problems := adapter.WisplyError{}
	result := hasValidProgramModifyDetails(details)
	if !result.IsValid {
		problems.Data = result
		return problems, errors.New("Invalid details for the program")
	}
	fieldList := "`institution`, `title`, `code`, `year`, `ucas_code`, `level`, `content`, `subject`"
	questions := "?, ?, ?, ?, ?, ?, ?, ?"
	sql := "INSERT INTO `institution_program` (" + fieldList + ") VALUES (" + questions + ")"
	query, err := database.Connection.Prepare(sql)

	if err != nil {
		return problems, err
	}
	title := details["program-title"].(string)
	code := details["program-code"].(string)
	year := details["program-year"].(string)
	institution := details["program-institution"].(int)
	ucasCode := details["program-ucas-code"].(string)
	level := details["program-level"].(string)
	content := details["program-content"].(string)
	subject := details["program-subject"].(string)
	_, err = query.Exec(institution, title, code, year, ucasCode, level, content, subject)

	return problems, err
}
