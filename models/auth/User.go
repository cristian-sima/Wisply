package auth

import (
	"fmt"
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

	err := db.Raw("SELECT id, username, password, administrator FROM user").QueryRow(&user)

	fmt.Println(err)

	return user
}

func (this *User) IsAdministrator() bool {
	return this.Administrator
}
