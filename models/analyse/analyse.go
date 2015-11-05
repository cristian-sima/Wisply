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
	timestamp      int64
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
	fieldList := "`id`, `timestamp`"
	sql := "SELECT " + fieldList + " FROM `analyse` WHERE id=? "
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	query.QueryRow(id).Scan(&analyser.id, &analyser.timestamp)

	digester := word.NewDigester("")
	analyser.moduleDigester = &digester

	return analyser
}

// CreateAnalyser inserts into database the procsss and starts it
func CreateAnalyser() Analyser {
	sql := "INSERT INTO `analyse` (`timestamp`) VALUES (?)"
	query, err := database.Connection.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	timestamp := getCurrentTimestamp()
	_, err = query.Exec(timestamp)

	// get the id
	id := ""
	fieldList := "`id`"
	sql = "SELECT " + fieldList + " FROM `analyse` WHERE timestamp=? "
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
