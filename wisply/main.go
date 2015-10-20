package main

import (
	_ "github.com/cristian-sima/Wisply/wisply/docs"
	_ "github.com/cristian-sima/Wisply/wisply/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
