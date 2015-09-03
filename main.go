package main

import (
	_ "Wisply/routers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {

    // [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
    // begoo.AppConfig.String("mysqluser") + ":" + begoo.AppConfig.String("mysqlpassword") + "@" + begoo.AppConfig.String("mysqlhost") + begoo.AppConfig.String("mysqldb") + "?charset=utf8"
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
    orm.RegisterDataBase("default", "mysql", databaseString)

}


func main() {
	beego.Run()
}