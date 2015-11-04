package education

import (
	"errors"
	"fmt"
	"html/template"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
)

// Subject represents a subject area
type Subject struct {
	id          int
	name        string
	description string
}

// GetID returns the ID of the subject
func (subject Subject) GetID() int {
	return subject.id
}

// GetName returns the name of the subject
func (subject Subject) GetName() string {
	return subject.name
}

// Delete removes the subject of study and any information about it
func (subject Subject) Delete() error {
	sql := "DELETE FROM `subject_area` WHERE id = ? "
	stmt, err := database.Connection.Prepare(sql)
	stmt.Exec(subject.id)
	return err
}

// GetHTMLDescription returns the description as HTML code
func (subject Subject) GetHTMLDescription() template.HTML {
	description := subject.GetDescription()
	return template.HTML([]byte(description))
}

// SetDescription sets the static description
func (subject Subject) SetDescription(description string) error {
	sql := "UPDATE `subject_area` SET description=? WHERE id=?"
	stmt, err := database.Connection.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(description, strconv.Itoa(subject.id))
	if err != nil {
		return err
	}
	return err
}

// GetDescription returns the static description from database
func (subject Subject) GetDescription() string {
	if subject.description != "" {
		return subject.description
	}
	fieldSet := "`description`"
	sql := "SELECT " + fieldSet + " FROM `subject_area` WHERE id = ? LIMIT 0,1"
	query, _ := database.Connection.Prepare(sql)
	query.QueryRow(subject.id).Scan(&subject.description)

	return subject.description
}

// Modify changes the details of the subject
func (subject Subject) Modify(details map[string]interface{}) error {
	result := areValidSubjectDetails(details)
	if !result.IsValid {
		return errors.New("Problem with the details")
	}
	setClause := "SET `name`=?"
	whereClause := "WHERE `id`= ?"
	sql := "UPDATE `subject_area` " + setClause + " " + whereClause
	query, err := database.Connection.Prepare(sql)
	query.Exec(details["name"].(string), subject.id)
	return err
}

// GetDefinitions returns the formal definitions for the subject
func (subject Subject) GetDefinitions() []Definition {
	var list []Definition
	fieldList := "`id`, `content`, `source`, `subject`"
	orderClause := "ORDER BY `content` ASC"
	whereClause := "WHERE `subject` = ?"
	sql := "SELECT " + fieldList + " FROM `subject_area_definition` " + whereClause + " " + orderClause
	rows, err := database.Connection.Query(sql, strconv.Itoa(subject.GetID()))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		item := Definition{}
		rows.Scan(&item.id, &item.content, &item.source, &item.subject)
		list = append(list, item)
	}
	return list
}

// GetKAs returns the knowledge areas for the subject
func (subject Subject) GetKAs() []KA {
	var list []KA
	fieldList := "`id`, `content`, `source`, `code`, `title`, `subject`"
	orderClause := "ORDER BY `content` ASC"
	whereClause := "WHERE `subject` = ?"
	sql := "SELECT " + fieldList + " FROM `subject_area_ka` " + whereClause + " " + orderClause
	rows, err := database.Connection.Query(sql, strconv.Itoa(subject.GetID()))
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		item := KA{}
		rows.Scan(&item.id, &item.content, &item.source, &item.code, &item.title, &item.subject)
		list = append(list, item)
	}
	return list
}

// NewSubject creates a new subject by ID
func NewSubject(ID string) (*Subject, error) {
	subject := &Subject{}
	fieldList := "`id`, `name`"
	sql := "SELECT " + fieldList + " FROM `subject_area` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		return subject, err
	}
	query.QueryRow(ID).Scan(&subject.id, &subject.name)
	return subject, nil
}

// CreateSubject adds a new subject into database
func CreateSubject(name string) error {
	result := isValidName(name)
	if !result.IsValid {
		return errors.New("Invalid name for the subject")
	}
	sql := "INSERT INTO `subject_area` (`name`) VALUES (?)"
	query, err := database.Connection.Prepare(sql)
	query.Exec(name)
	return err
}

// GetAllSubjects returns a list with all the subjects of study
func GetAllSubjects() []Subject {
	var list []Subject
	fieldList := "`id`, `name`"
	orderClause := "ORDER BY `name` ASC"
	sql := "SELECT " + fieldList + " FROM `subject_area` " + orderClause
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		item := Subject{}
		rows.Scan(&item.id, &item.name)
		list = append(list, item)

	}
	return list
}
