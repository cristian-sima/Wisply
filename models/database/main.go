// Package database contains the functions to connect to database
// Also, it provides some useful objects such as SQLBuffer or SQLOptions
package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cristian-sima/Wisply/models/config"

	// the driver
	_ "github.com/go-sql-driver/mysql"
)

// Connection represents the object which holds the connection to the database
var (
	Connection *sql.DB
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
		Connection = db
		// Open doesn't open a connection. Validate DSN data:
		err = db.Ping()
		if err == nil {
			fmt.Println("[SUCCESS]: Connected to database!")
			fmt.Println("")
			connected = true
		} else {
			tryAgain := "Wisply tries again in 3 seconds..."
			fmt.Println("[ERROR]: Problem trying to connect to database." + tryAgain)
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

	firstPart := configuration.Username + ":" + configuration.Password + "@"
	secondPart := "(" + mysqlAddress + ")/" + configuration.Database

	databaseString = firstPart + secondPart + "?charset=utf8"

	return databaseString
}
