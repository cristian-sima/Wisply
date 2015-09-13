package database

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	config "github.com/cristian-sima/Wisply/models/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Database struct {
}

func (database *Database) Init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	go database.connect()
}

func (database *Database) connect() {
	var (
		connected      bool = false
		databaseString string
	)

	databaseString = database.getString()

	for !connected {
		fmt.Println("[INFO] Connecting to database...")
		error := orm.RegisterDataBase("default", "mysql", databaseString)
		if error == nil {
			fmt.Println("[SUCCESS]: Connected to database!\n")
			connected = true
		} else {
			fmt.Println("[ERROR]: Problem trying to connect to database. Wisply tries again in 3 seconds...\n")
			time.Sleep(3000 * time.Millisecond)
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
