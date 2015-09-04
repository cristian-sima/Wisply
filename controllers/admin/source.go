package admin

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "regexp"
    "fmt"
)

type SourceController struct {
	beego.Controller
}

func (c *SourceController) List() {

    type Source2 struct {
        Id int
        Name string
        Url string
        Description string
    }


    var (
        sources []Source2
        exists bool = false
        )
    o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database
 
    num, err := o.Raw("SELECT id, name, url, description FROM source").QueryRows(&sources)

    if err != nil { 
        c.Data["messageContent"] = "It's a shame... There was a problem. Maybe you want to try again?"  
        c.TplNames = "general/message/error.tpl"
        c.Data["messageLink"] = "/admin/source";
    } else {

        if num != 0 {
            exists = true
        }


        fmt.Println(sources);

        c.Data["anything"] = exists
        c.Data["sources"] = sources
        c.TplNames = "general/source/list.tpl"
    }
    
    c.Layout = "general/admin.tpl"
}

func (c *SourceController) Get() {
    c.Layout = "general/admin.tpl"
	c.TplNames = "general/source/add.tpl"
}



func isURL(url string) bool{
	Re := regexp.MustCompile(`^(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$`)
	return Re.MatchString(url)
}

func (c *SourceController) Post() {


    var (
        name string = c.GetString("source-name")
        description string = c.GetString("source-description")
        URL string = c.GetString("source-URL")
    )

    o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database
    
    
    // VALIDATE
    
    if ;len(name) > 255 || name == ""  || 
       len(URL) > 255 || URL == "" ||
       len(description) > 255 ||
       !isURL(URL) {
        c.Data["messageContent"] = "There was a problem with fields. Try again"  
        c.TplNames = "general/message/error.tpl"
        c.Data["messageLink"] = "/admin/source/add"
    } else {
        // STORE

        elememts := []string{name, description, URL}

        _, err := o.Raw("INSERT INTO `source` (`name`, `description`, `url`) VALUES (?, ?, ?)", elememts).Exec()

        if err == nil {
            c.Data["messageContent"] = "The source has been added!" 
            c.TplNames = "general/message/success.tpl"
            c.Data["messageLink"] = "/admin/source"
        } else {       
            c.Data["messageContent"] = "It's a shame... There was a problem. Maybe you want to try again?"  
            c.TplNames = "general/message/error.tpl"
            c.Data["messageLink"] = "/admin/source/add";
        }

    }


    c.Layout = "general/status.tpl"
}

