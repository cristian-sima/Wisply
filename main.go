package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/cristian-sima/Wisply/routers"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"os"
	"time"
)

type SQLConfiguration struct {
	Username  string
	Password string
	Host string
	Port string
	Database string
}

func getConfigurationFromFile() SQLConfiguration {
	fileName := "conf/database.json"
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	configuration := SQLConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("There was an error with the configuration file (conf/database.conf):", err)
	}
	return configuration;
}


/*
 * [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
 */
func getSQLString() string {
	configuration := 	getConfigurationFromFile();
	var (
		mysqlAddress   string = configuration.Host + ":" + configuration.Port
		databaseString string = configuration.Username + ":" + configuration.Password + "@" + "(" + mysqlAddress + ")/" + configuration.Database + "?charset=utf8"
	)
	return databaseString
}

func init() {
	databaseString := getSQLString()
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	go connectToDatabase(databaseString)
}

func connectToDatabase(databaseString string) {
	var connected bool = false
	for !connected {
		fmt.Println("Connecting to database...")
		error := orm.RegisterDataBase("default", "mysql", databaseString)
		if error == nil {
			fmt.Println("[Success]: Connected to database!\n")
			connected = true
		} else {
			fmt.Println("[Error]: Problem trying to connect to database. Wisply tries again in 3 seconds...\n")
			time.Sleep(3000 * time.Millisecond)
		}
	}
}

func loadPageNotFound(rw http.ResponseWriter, r *http.Request) {
	path := "/errors/404.html"
	loadError(rw, r, path)
}

func loadDatabaseError(rw http.ResponseWriter, r *http.Request) {
	path := "/errors/database.html"
	loadError(rw, r, path)
}

func loadError(rw http.ResponseWriter, r *http.Request, path string) {
	t, _ := template.ParseFiles(beego.ViewsPath + path)
	data := make(map[string]interface{})
	t.Execute(rw, data)
}

func main() {
	beego.Errorhandler("404", loadPageNotFound)
	beego.Errorhandler("databaseError", loadDatabaseError)
	beego.Run()
}
