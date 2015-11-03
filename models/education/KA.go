package education

import (
	"errors"

	"github.com/cristian-sima/Wisply/models/database"
)

// KA represents a knowledge area
type KA struct {
	id      int
	content string
	code    string
	source  string
	title   string
	program int
}

// GetID returns the ID of the ka
func (ka KA) GetID() int {
	return ka.id
}

// GetContent returns content
func (ka KA) GetContent() string {
	return ka.content
}

// GetTitle returns title
func (ka KA) GetTitle() string {
	return ka.title
}

// GetSource returns the source
func (ka KA) GetSource() string {
	return ka.source
}

// GetCode returns the code
func (ka KA) GetCode() string {
	return ka.code
}

// GetProgramID returns the ID of the program
func (ka KA) GetProgramID() int {
	return ka.program
}

// Delete removes the ka of study and any information about it
func (ka KA) Delete() error {
	sql := "DELETE FROM `program_of_study_ka` WHERE id = ? "
	stmt, err := database.Connection.Prepare(sql)
	stmt.Exec(ka.id)
	return err
}

// Modify changes the details of the ka
func (ka KA) Modify(details map[string]interface{}) error {
	result := hasKAValidDetails(details)
	if !result.IsValid {
		return errors.New("Problem with the fields")
	}
	setClause := "SET `content`=?, `source`=?, `code`=?, `title`=? "
	whereClause := "WHERE `id`= ?"
	sql := "UPDATE `program_of_study_ka` " + setClause + " " + whereClause
	query, err := database.Connection.Prepare(sql)
	content := details["ka-content"].(string)
	source := details["ka-source"].(string)
	code := details["ka-code"].(string)
	title := details["ka-title"].(string)
	query.Exec(content, source, code, title, ka.id)
	return err
}

// NewKA creates a new ka by ID
func NewKA(ID string) (*KA, error) {
	ka := &KA{}
	fieldList := "`id`, `content`, `source`, `code`, `title`, `program`"
	sql := "SELECT " + fieldList + " FROM `program_of_study_ka` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return ka, err
	}
	query.QueryRow(ID).Scan(&ka.id, &ka.content, &ka.source, &ka.code, &ka.title, &ka.program)
	return ka, nil
}

// CreateKA adds a new ka into database
func CreateKA(details map[string]interface{}) error {
	result := hasDefinationValidDetails(details)
	if !result.IsValid {
		return errors.New("Invalid name for the ka")
	}
	sql := "INSERT INTO `program_of_study_ka` (`content`, `source`, `code`, `title`, `program`) VALUES (?, ?, ?, ?, ?)"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	content := details["ka-content"].(string)
	source := details["ka-source"].(string)
	program := details["ka-program"].(int)
	code := details["ka-code"].(string)
	title := details["ka-title"].(string)
	_, err = query.Exec(content, source, code, title, program)
	return err
}
