package auth

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int
	Username      string
	Password      string
	Email         string
	Administrator bool
}

func NewUser(id string) User {
	var user User
	db := orm.NewOrm()

	db.Raw("SELECT id, username, password, administrator FROM user").QueryRow(&user)
	return user
}

func (this *User) IsAdministrator() bool {
	return this.Administrator
}
