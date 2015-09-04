package admin

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "regexp"
    "strconv"
    _ "fmt"
)

func isURL(url string) bool{
	Re := regexp.MustCompile(`^(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$`)
	return Re.MatchString(url)
}

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

        c.Data["anything"] = exists
        c.Data["sources"] = sources
        
        c.TplNames = "general/source/list.tpl"
    }
    
    c.Layout = "general/admin.tpl"
}

func (c *SourceController) AddNewSource() {

    c.Data["action"] = "Add"
    c.Data["legend"] = "Add a new source"
    c.Data["actionURL"] = "";
    c.Data["actionType"] = "POST";

    c.Layout = "general/admin.tpl"
	c.TplNames = "general/source/form.tpl"
}

func (c *SourceController) Insert() {

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
        c.Data["messageLink"] = "/admin/sources/add"
    } else {
        // STORE
        elememts := []string{name, description, URL}

        _, err := o.Raw("INSERT INTO `source` (`name`, `description`, `url`) VALUES (?, ?, ?)", elememts).Exec()

        if err == nil {
            c.Data["messageContent"] = "The source has been added!" 
            c.TplNames = "general/message/success.tpl"
            c.Data["messageLink"] = "/admin/sources"
        } else {       
            c.Data["messageContent"] = "It's a shame... There was a problem. Maybe you want to try again?"  
            c.TplNames = "general/message/error.tpl"
            c.Data["messageLink"] = "/admin/sources/add";
        }
    }
    c.Layout = "general/status.tpl"
}


func (c *SourceController) Edit() {
   
   o := orm.NewOrm()
   o.Using("default") // Using default, you can use other database
        
   type Source struct {
        Name string
        Url string
        Description string
    }

   id := c.Ctx.Input.Param(":id")

    if _, error := strconv.Atoi(id); error != nil {
        c.Abort("404")       
    } else {
        source := new(Source)
        error := o.Raw("SELECT name, url, description FROM source WHERE id = ?", id).QueryRow(&source)
        if error != nil {
            c.Abort("404")
        } else {
            c.Data["action"] = "Modify"
            c.Data["legend"] = "Modify details"
            c.Data["actionURL"] = ""
            c.Data["actionType"] = "POST"

            c.Data["sourceName"] = source.Name
            c.Data["sourceUrl"] = source.Url
            c.Data["sourceDescription"] = source.Description

            c.Layout = "general/admin.tpl"
            c.TplNames = "general/source/form.tpl"
        }
    }
}

func (c *SourceController) Update() {

    var (
        id string = c.Ctx.Input.Param(":id")
        name string = c.GetString("source-name")
        description string = c.GetString("source-description")
        URL string = c.GetString("source-URL")
    )

    o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database
        
    // VALIDATE


    if len(name) > 255 || name == ""  || 
       len(URL) > 255 || URL == "" ||
       len(description) > 255 ||
       !isURL(URL) {
        c.Data["messageContent"] = "There was a problem with fields. Try again"  
        c.TplNames = "general/message/error.tpl"
        c.Data["messageLink"] = "/admin/sources/modify/" + id
    } else {
        // STORE
        elememts := []string{name, description, URL, id}

        _, err := o.Raw("UPDATE `source` SET name=?, description=?, url=? WHERE id=?", elememts).Exec()

        if err == nil {
            c.Data["messageContent"] = "The source has been modified!" 
            c.TplNames = "general/message/success.tpl"
            c.Data["messageLink"] = "/admin/sources"
        } else {       
            c.Abort("404")
        }
    }
    c.Layout = "general/status.tpl"
}


func (c *SourceController) Delete () {
    
    var (
        id string = c.Ctx.Input.Param(":id")
    )

    o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database
        

    elememts := []string{id}

    _, err := o.Raw("DELETE from `source` WHERE id=?", elememts).Exec()

    if err == nil {
        c.Data["messageContent"] = "The source has been deleted!" 
        c.TplNames = "general/message/success.tpl"
        c.Data["messageLink"] = "/admin/sources"
    } else {       
        c.Abort("404")
    }
    
    c.Layout = "general/status.tpl"
}