package main

import (
	_ "Wisply/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"html/template"
	"fmt"
	"time"
)

func init() {

	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	var (
		mysqlUsername string = "wisply"
		mysqlPassword string = "DNeaMKvz4t4DtL6b"
		mysqlHost string		= "127.0.0.1"
		mysqlPort string		= "3306"
		mysqlAddress string = mysqlHost + ":" +	mysqlPort
		mysqlDatabase string = "wisply"
		databaseString string = mysqlUsername + ":" + mysqlPassword +	"@" + "(" + mysqlAddress +")/" + mysqlDatabase+ "?charset=utf8";
		)

		orm.RegisterDriver("mysql", orm.DR_MySQL)
		go connectToDatabase(databaseString);
	}

	func connectToDatabase (databaseString string) {
		var connected bool = false
		for !connected {
			fmt.Println("Connecting to database...")
			error := orm.RegisterDataBase("default", "mysql", databaseString)
			if error == nil {
				fmt.Println("[Success]: Connected to database!\n");
				connected = true;
				} else {
					fmt.Println("[Error]: Problem trying to connect to database. Wisply tries again in 3 seconds...\n");
					time.Sleep(3000 * time.Millisecond)
				}
			}
		}

		func loadPageNotFound(rw http.ResponseWriter, r *http.Request){
			path := "/errors/404.html"
			loadError(rw, r, path);
		}

		func loadDatabaseError(rw http.ResponseWriter, r *http.Request){
			path := "/errors/database.html"
			loadError(rw, r, path);
		}

		func loadError(rw http.ResponseWriter, r *http.Request, path string) {
			t,_:= template.ParseFiles(beego.ViewsPath+path)
			data :=make(map[string]interface{})
			t.Execute(rw, data)
		}

		func main() {
			beego.Errorhandler("404", loadPageNotFound)
			beego.Errorhandler("databaseError", loadDatabaseError)
			beego.Run()
		}
