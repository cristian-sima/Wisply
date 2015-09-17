package database

import (
	"database/sql"
	"fmt"
	"time"

	config "github.com/cristian-sima/Wisply/models/config"

	// the driver
	_ "github.com/go-sql-driver/mysql"
)

// Database It represents the connection to the database
var (
	Database *sql.DB
)

// Db manages the connection to database
type Db struct {
}

// Init registers the driver and tries to connect to database
func (database *Db) Init() {
	go database.connect()
}

// It tries to connect to database. If it can not, tries again
func (database *Db) connect() {
	var (
		connected        bool
		databaseString   string
		delayMiliseconds time.Duration = 3000
	)

	databaseString = database.getString()

	for !connected {
		fmt.Println("[INFO] Connecting to database...")
		db, err := sql.Open("mysql", databaseString)
		Database = db
		// Open doesn't open a connection. Validate DSN data:
		err = db.Ping()
		if err == nil {
			fmt.Println("[SUCCESS]: Connected to database!")
			fmt.Println("")
			connected = true
		} else {
			fmt.Println("[ERROR]: Problem trying to connect to database. Wisply tries again in 3 seconds...")
			fmt.Println()
			time.Sleep(delayMiliseconds * time.Millisecond)
		}
	}
}

func (database *Db) getString() string {
	var (
		mysqlAddress   string
		databaseString string
	)
	configuration := config.GetDatabase()
	mysqlAddress = configuration.Host + ":" + configuration.Port
	databaseString = configuration.Username + ":" + configuration.Password + "@" + "(" + mysqlAddress + ")/" + configuration.Database + "?charset=utf8"
	return databaseString
}
