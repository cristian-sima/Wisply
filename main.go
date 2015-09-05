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
        mysqlHost string    = "127.0.0.1"
        mysqlPort string    = "3306"
        mysqlAddress string = mysqlHost + ":" +  mysqlPort
        mysqlDatabase string = "wisply"
        databaseString string = mysqlUsername + ":" + mysqlPassword +  "@" + "(" + mysqlAddress +")/" + mysqlDatabase+ "?charset=utf8";
    )   


    orm.RegisterDriver("mysql", orm.DR_MySQL)
    
    go connectToDatabase(databaseString);
}

func connectToDatabase (databaseString string) {    
    var connected bool = false    
	for !connected {
        error := orm.RegisterDataBase("default", "mysql", databaseString)
        if error == nil {
            fmt.Println("Connected to database");
            connected = true;
        } else {
            fmt.Println("Problem trying to connect to database. Try again in 3 seconds");
            time.Sleep(3000 * time.Millisecond)
        }
	}
}





func loadPageNotFound(rw http.ResponseWriter, r *http.Request){
    t,_:= template.ParseFiles(beego.ViewsPath+"/errors/404.html")
    data :=make(map[string]interface{})
    t.Execute(rw, data)
}

func loadDatabaseError(rw http.ResponseWriter, r *http.Request){
    t,_:= template.ParseFiles(beego.ViewsPath+"/errors/database.html")
    data :=make(map[string]interface{})
    t.Execute(rw, data)
}

func main() {
    beego.Errorhandler("404", loadPageNotFound)
    beego.Errorhandler("databaseError", loadDatabaseError)
	beego.Run()
}
