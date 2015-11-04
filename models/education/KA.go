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
	subject int
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

// GetSubjectID returns the ID of the subject
func (ka KA) GetSubjectID() int {
	return ka.subject
}

// Delete removes the ka of study and any information about it
func (ka KA) Delete() error {
	sql := "DELETE FROM `subject_area_ka` WHERE id = ? "
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
	sql := "UPDATE `subject_area_ka` " + setClause + " " + whereClause
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
	fieldList := "`id`, `content`, `source`, `code`, `title`, `subject`"
	sql := "SELECT " + fieldList + " FROM `subject_area_ka` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return ka, err
	}
	query.QueryRow(ID).Scan(&ka.id, &ka.content, &ka.source, &ka.code, &ka.title, &ka.subject)
	return ka, nil
}

// CreateKA adds a new ka into database
func CreateKA(details map[string]interface{}) error {
	result := hasDefinationValidDetails(details)
	if !result.IsValid {
		return errors.New("Invalid name for the ka")
	}
	sql := "INSERT INTO `subject_area_ka` (`content`, `source`, `code`, `title`, `subject`) VALUES (?, ?, ?, ?, ?)"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	content := details["ka-content"].(string)
	source := details["ka-source"].(string)
	subject := details["ka-subject"].(int)
	code := details["ka-code"].(string)
	title := details["ka-title"].(string)
	_, err = query.Exec(content, source, code, title, subject)
	return err
}
