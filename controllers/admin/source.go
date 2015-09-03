package admin

import (
    "fmt"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "net/url"
)


type SourceController struct {
	beego.Controller
}

func (c *SourceController) List() {
    c.Layout = "general/admin.tpl"
	c.TplNames = "general/source/list.tpl"
}

func (c *SourceController) Get() {
    c.Layout = "general/admin.tpl"
	c.TplNames = "general/source/add.tpl"
}



func (c *SourceController) Post() {

    var (
            name string = c.GetString("source-name")
            description string = c.GetString("source-description")
            URL string = c.GetString("source-URL")
    )
    
    _, errorURL := url.Parse(URL)
    

    o := orm.NewOrm()
    o.Using("default") // Using default, you can use other database
    
    
    // VALIDATE
    
    if ;len(name) > 255 || name == ""  || 
       len(URL) > 255 || URL == "" ||
       len(description) > 255 ||
       errorURL != nil {
        c.Data["messageContent"] = "There was a problem with fields. Try again"  
        c.TplNames = "general/message/error.tpl"
    } else {
        // STORE

        elememts := []string{name, description, URL}

        _, err := o.Raw("INSERT INTO `source` (`name`, `description`, `URL`) VALUES (?, ?, ?)", elememts).Exec()

        if err == nil {
            c.Data["messageContent"] = "The source has been added!" 
            c.TplNames = "general/message/success.tpl"

        } else {       
            fmt.Println("mysql row affected nums: ", err) 
            c.Data["messageContent"] = "It's a shame... There was a problem. Maybe you want to try again?"  
            c.TplNames = "general/message/error.tpl"
        }

    }


    c.Data["messageLink"] = "/admin/source";
    c.Layout = "general/status.tpl"
}

