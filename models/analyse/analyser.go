package analyse

import (
	"fmt"
	"log"
	"time"

	"github.com/cristian-sima/Wisply/models/analyse/word"
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Analyser processes the data for an institution
type Analyser struct {
	id             string
	moduleDigester *word.Digester
	start          int64
	end            int64
}

// Start starts the analyse process
func (analyser Analyser) Start() {
	start := time.Now()
	analyser.processModules()
	elapsed := time.Since(start)
	log.Printf("All searches have taken %s", elapsed)
}

func (analyser Analyser) processModules() {

	// get all institutions

	institutions := repository.GetAllInstitutions()

	for _, institution := range institutions {
		analyser.CreateInstitutionAnalyser(institution)
	}

	// for each subject
	analyser.Finish()
}

// GetID returns the id
func (analyser Analyser) GetID() string {
	return analyser.id
}

// IsFinished checks if the analyser has finished
func (analyser Analyser) IsFinished() bool {
	return analyser.end != 0
}

// GetStartDate returns the start date of the action in a human readable form
func (analyser *Analyser) GetStartDate() string {
	return analyser.getDate(analyser.start)
}

// GetEndDate returns the end date of the action in a human readable form
func (analyser *Analyser) GetEndDate() string {
	if analyser.end == 0 {
		return "Still working..."
	}
	return analyser.getDate(analyser.end)
}

func (analyser *Analyser) getDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(dateFormat)
}

// Finish sets the analyser as finished
func (analyser Analyser) Finish() {
	sql := "UPDATE `analyse` SET `end`=? WHERE id=?"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	finish := getCurrentTimestamp()
	_, err = query.Exec(finish, analyser.id)

}

// Delete analyser
func (analyser Analyser) Delete() {
	sql := "DELETE FROM `analyse` WHERE id = ? "
	stmt, _ := database.Connection.Prepare(sql)
	stmt.Exec(analyser.id)
}

// CreateInstitutionAnalyser creates a new analyser for institution
func (analyser Analyser) CreateInstitutionAnalyser(institution repository.Institution) InstitutionAnalyser {
	institutionAnalyser := InstitutionAnalyser{
		parent:      analyser,
		institution: institution,
	}
	institutionAnalyser.Start()
	return institutionAnalyser
}

// NewAnalyser creates a new Analyser object for the institution
func NewAnalyser(id string) Analyser {
	analyser := Analyser{}
	fieldList := "`id`, `start`, `end`"
	sql := "SELECT " + fieldList + " FROM `analyse` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	query.QueryRow(id).Scan(&analyser.id, &analyser.start, &analyser.end)

	digester := word.NewDigester("")
	analyser.moduleDigester = &digester

	return analyser
}

// CreateAnalyser inserts into database the procsss and starts it
func CreateAnalyser() Analyser {
	sql := "INSERT INTO `analyse` (`start`) VALUES (?)"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	timestamp := getCurrentTimestamp()
	_, err = query.Exec(timestamp)

	// get the id
	id := ""
	fieldList := "`id`"
	sql = "SELECT " + fieldList + " FROM `analyse` WHERE start=? "
	query, err = database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	err = query.QueryRow(timestamp).Scan(&id)
	fmt.Println(err)
	return NewAnalyser(id)
}

func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// GetAll returns a list with all the analyses
func GetAll() []Analyser {
	var list []Analyser
	fieldList := "`id`, `start`, `end`"
	orderClause := "ORDER BY `start` DESC"
	sql := "SELECT " + fieldList + " FROM `analyse` " + orderClause
	rows, _ := database.Connection.Query(sql)
	for rows.Next() {
		item := Analyser{}
		rows.Scan(&item.id, &item.start, &item.end)
		list = append(list, item)
	}
	return list
}
