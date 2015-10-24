package curriculum

import (
	"errors"

	"github.com/cristian-sima/Wisply/models/database"
)

// Program of study contains multiple Modules
type Program struct {
	id   int
	name string
}

// GetID returns the ID of the program
func (program Program) GetID() int {
	return program.id
}

// GetName returns the name of the program
func (program Program) GetName() string {
	return program.name
}

// NewProgram creates a new program by ID
func NewProgram(ID string) (*Program, error) {
	program := &Program{}
	fieldList := "`id`, `name`"
	sql := "SELECT " + fieldList + " FROM `program_of_study` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return program, err
	}
	query.QueryRow(ID).Scan(&program.id, &program.name)
	return program, nil
}

// CreateProgram adds a new program into database
func CreateProgram(name string) error {
	result := isValidName(name)
	if !result.IsValid {
		return errors.New("Invalid name for the program")
	}
	sql := "INSERT INTO `program_of_study` (`name`) VALUES (?)"
	query, err := database.Connection.Prepare(sql)
	query.Exec(name)
	return err
}

// GetAllPrograms returns a list with all the programs of study
func GetAllPrograms() []Program {
	var list []Program
	fieldList := "`id`, `name`"
	sql := "SELECT " + fieldList + " FROM `program_of_study` "
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		item := Program{}
		rows.Scan(&item.id, &item.name)
		list = append(list, item)

	}
	return list
}
