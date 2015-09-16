package database

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	config "github.com/cristian-sima/Wisply/models/config"

	// used by orm
	_ "github.com/go-sql-driver/mysql"
)

// Database manages the connection to database
type Database struct {
}

// Init registers the driver and tries to connect to database
func (database *Database) Init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	go database.connect()
}

// It tries to connect to database. If it can not, tries again
func (database *Database) connect() {
	var (
		connected        bool
		databaseString   string
		delayMiliseconds time.Duration = 3000
	)

	databaseString = database.getString()

	for !connected {
		fmt.Println("[INFO] Connecting to database...")
		error := orm.RegisterDataBase("default", "mysql", databaseString)
		if error == nil {
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

func (database *Database) getString() string {
	var (
		mysqlAddress   string
		databaseString string
	)
	configuration := config.GetDatabase()
	mysqlAddress = configuration.Host + ":" + configuration.Port
	databaseString = configuration.Username + ":" + configuration.Password + "@" + "(" + mysqlAddress + ")/" + configuration.Database + "?charset=utf8"
	return databaseString
}
