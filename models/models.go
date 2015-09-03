package models

import (
    "github.com/astaxie/beego/orm"
)

type Source struct {
    Id          int
    Name        string
    URL         string
    Description string
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(Source))
}