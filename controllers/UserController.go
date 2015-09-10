package controllers

import (
	"fmt"
	. "github.com/cristian-sima/Wisply/models/auth"
	"strings"
)

type UserController struct {
	AdminController
	model AuthModel
}

func (controller *UserController) ListUsers() {

	var exists bool = false

	users := controller.model.GetAll()

	exists = (len(users) != 0)

	controller.Data["anything"] = exists
	controller.Data["users"] = users
	controller.TplNames = "general/user/list.tpl"
	controller.Layout = "general/admin.tpl"
}

func (controller *UserController) Modify() {

	var id string

	id = controller.Ctx.Input.Param(":id")

	user, err := controller.model.GetUserById(id)

	if err != nil {
		fmt.Println(err)
		controller.Abort("databaseError")
	} else {
		controller.showModifyForm(user)
	}
}

func (controller *UserController) Update() {

	var userId, isAdministrator string
	rawData := make(map[string]interface{})

	userId = controller.Ctx.Input.Param(":id")

	isAdministrator = strings.TrimSpace(controller.GetString("modify-administrator"))
	rawData["administrator"] = isAdministrator

	_, err := controller.model.GetUserById(userId)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		problems, err := controller.model.ValidateModifyUser(rawData)
		if err != nil {
			controller.DisplayErrorMessages(problems)
		} else {
			databaseError := controller.model.UpdateUserType(userId, isAdministrator)
			if databaseError != nil {
				controller.Abort("databaseError")
			} else {
				controller.DisplaySuccessMessage("The source has been modified!", "/admin/users/")
			}
		}
	}
}

func (controller *UserController) Delete() {
	var id string
	id = controller.Ctx.Input.Param(":id")
	user, err := controller.model.GetUserById(id)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		databaseError := controller.model.DeleteUserById(id)
		if databaseError != nil {
			controller.Abort("databaseError")
		} else {
			controller.DisplaySuccessMessage("The user ["+user.Username+"] has been deleted. Well done!", "/admin/users/")
		}
	}
}

func (controller *UserController) showModifyForm(user *User) {
	controller.GenerateXsrf()
	controller.Data["userUsername"] = user.Username

	if user.Administrator {
		controller.Data["isAdministrator"] = true
	} else {
		controller.Data["isUser"] = true
	}
	controller.Layout = "general/admin.tpl"
	controller.TplNames = "general/user/form.tpl"
}
