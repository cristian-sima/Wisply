package education

import (
	"errors"

	"github.com/cristian-sima/Wisply/models/database"
)

// Definition of study contains multiple Modules
type Definition struct {
	id      int
	content string
	source  string
	subject int
}

// GetID returns the ID of the definition
func (definition Definition) GetID() int {
	return definition.id
}

// GetContent returns content
func (definition Definition) GetContent() string {
	return definition.content
}

// GetSource returns the source
func (definition Definition) GetSource() string {
	return definition.source
}

// GetSubjectID returns the ID of the subject
func (definition Definition) GetSubjectID() int {
	return definition.subject
}

// Delete removes the definition of study and any information about it
func (definition Definition) Delete() error {
	sql := "DELETE FROM `subject_area_definition` WHERE id = ? "
	stmt, err := database.Connection.Prepare(sql)
	stmt.Exec(definition.id)
	return err
}

// Modify changes the details of the definition
func (definition Definition) Modify(details map[string]interface{}) error {
	result := hasDefinationValidDetails(details)
	if !result.IsValid {
		return errors.New("Problem with the fields")
	}
	setClause := "SET `content`=?, `source`=? "
	whereClause := "WHERE `id`= ?"
	sql := "UPDATE `subject_area_definition` " + setClause + " " + whereClause
	query, err := database.Connection.Prepare(sql)
	content := details["definition-content"].(string)
	source := details["definition-source"].(string)
	query.Exec(content, source, definition.id)
	return err
}

// NewDefinition creates a new definition by ID
func NewDefinition(ID string) (*Definition, error) {
	definition := &Definition{}
	fieldList := "`id`, `content`, `source`, `subject`"
	sql := "SELECT " + fieldList + " FROM `subject_area_definition` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return definition, err
	}
	query.QueryRow(ID).Scan(&definition.id, &definition.content, &definition.source, &definition.subject)
	return definition, nil
}

// CreateDefinition adds a new definition into database
func CreateDefinition(details map[string]interface{}) error {
	result := hasDefinationValidDetails(details)
	if !result.IsValid {
		return errors.New("Invalid name for the definition")
	}
	sql := "INSERT INTO `subject_area_definition` (`content`,`source`,`subject`) VALUES (?,?,?)"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	content := details["definition-content"].(string)
	source := details["definition-source"].(string)
	subject := details["definition-subject"].(int)
	_, err = query.Exec(content, source, subject)
	return err
}
