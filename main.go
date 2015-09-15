package main

import (
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
	. "github.com/cristian-sima/Wisply/models/database"
	_ "github.com/cristian-sima/Wisply/routers"
)

func init() {
	database := new(Database)
	database.Init()
}

func main() {
	beego.Errorhandler("404", loadPageNotFound)
	beego.Errorhandler("databaseError", loadDatabaseError)
	beego.Run()
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
