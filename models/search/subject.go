package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/education"
)

// SubjectSearch searches for Subjects
type SubjectSearch struct {
	*search
}

// Perform gets the results
func (search SubjectSearch) Perform() {
	allSubjects := search.getFromDB()
	for _, subject := range allSubjects {
		result := &Result{
			Title:       subject.GetName(),
			URL:         search.getURL(subject.GetID()),
			Description: "",
			Icon:        "/static/img/public/search/curriculum.png",
			Category:    "Subject area",
			IsVisible:   true,
		}
		search.response.AppendItem(result)
	}
}

// it gets the subjects from database
func (search SubjectSearch) getFromDB() []*education.Subject {
	var list []*education.Subject

	limitClause := search.options.GetLimit()
	fieldsList := "`id`"
	whereClause := "WHERE `name` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM `subject_area` " + whereClause + space + limitClause

	rows, err := database.Connection.Query(sql, search.likeQuery())

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var ID string
		rows.Scan(&ID)
		item, _ := education.NewSubject(ID)
		list = append(list, item)
	}
	return list
}

func (search SubjectSearch) getURL(curriculumID int) string {
	path := "/education/subjects/"
	action := path + strconv.Itoa(curriculumID)
	return action
}

// NewSubjectsSearch creates a search object for subjects
func NewSubjectsSearch(search *search) SubjectSearch {
	return SubjectSearch{
		search: search,
	}
}
