package search

import (
	"fmt"
	"strconv"

	"github.com/cristian-sima/Wisply/models/curriculum"
	"github.com/cristian-sima/Wisply/models/database"
)

// CurriculaSearch searches for curricula
type CurriculaSearch struct {
	*search
}

// Perform gets the results
func (search CurriculaSearch) Perform() {
	allCurriculums := search.getFromDB()
	for _, curriculum := range allCurriculums {
		result := &Result{
			Title:       curriculum.GetName(),
			URL:         search.getURL(curriculum.GetID()),
			Description: "",
			Icon:        "/static/img/public/search/curriculum.png",
			Category:    "Curriculum",
		}
		search.response.AppendItem(result)
	}
}

// gets the curricula
func (search CurriculaSearch) getFromDB() []*curriculum.Program {
	var list []*curriculum.Program

	limitClause := search.options.GetLimit()
	fieldsList := "`id`"
	whereClause := "WHERE `name` LIKE ?"
	sql := "SELECT DISTINCT " + fieldsList + " FROM `program_of_study` " + whereClause + space + limitClause

	rows, err := database.Connection.Query(sql, search.likeQuery())

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var ID string
		rows.Scan(&ID)
		item, _ := curriculum.NewProgram(ID)
		list = append(list, item)
	}
	return list
}

func (search CurriculaSearch) getURL(curriculumID int) string {
	path := "/curriculum/"
	action := path + strconv.Itoa(curriculumID)
	return action
}

// NewCurriculaSearch creates a search object for curricula
func NewCurriculaSearch(search *search) CurriculaSearch {
	return CurriculaSearch{
		search: search,
	}
}
