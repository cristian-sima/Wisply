package controllers

import (
	. "github.com/cristian-sima/Wisply/models/auth"
	"strings"
)

type AuthController struct {
	DefaultController
	Model AuthModel
}

func (c *AuthController) ShowLoginForm() {
	c.GenerateXsrf()
	c.TplNames = "general/auth/login.tpl"
	c.Layout = "general/layout.tpl"
}

func (c *AuthController) ShowRegisterForm() {
	c.GenerateXsrf()
	c.TplNames = "general/auth/register.tpl"
	c.Layout = "general/layout.tpl"
}

func (c *AuthController) CreateNewUser() {
	var (
		username, password, confirmPassowrd, email string
	)

	username = strings.TrimSpace(c.GetString("register-username"))
	password = strings.TrimSpace(c.GetString("register-password"))
	email = strings.TrimSpace(c.GetString("register-email"))
	confirmPassowrd = strings.TrimSpace(c.GetString("register-password-confirm"))

	rawData := make(map[string]interface{})
	rawData["username"] = username
	rawData["password"] = password
	rawData["email"] = email

	if confirmPassowrd != password {
		c.DisplayErrorMessage("The passwords do not match!")
	} else {
		problems, err := c.Model.ValidateRegisterDetails(rawData)
		if err != nil {
			c.DisplayErrorMessages(problems)
		} else {
			usernameAlreadyExists := c.Model.CheckUsernameExists(username)
			if usernameAlreadyExists {
				c.DisplayErrorMessage("The username is already used. Try another")
			} else {
				databaseError := c.Model.CreateNewUser(rawData)
				if databaseError != nil {
					c.Abort("databaseError")
				} else {
					c.DisplaySuccessMessage("Your account is ready!", "/auth/login/")
				}
			}
		}
	}
}

func (c *AuthController) LoginUser() {
	rawData := make(map[string]interface{})
	rawData["username"] = strings.TrimSpace(c.GetString("login-username"))
	rawData["password"] = strings.TrimSpace(c.GetString("login-password"))

	problems, err := c.Model.ValidateLoginDetails(rawData)
	if err != nil {
		c.DisplayErrorMessages(problems)
	} else {
		user, err := c.Model.TryLoginUser(rawData)
		if err != nil {
			c.DisplayErrorMessage("There was a problem while login. We think the username or the password were not good.")
		} else {
			c.saveLoginDetails(user)
			c.Redirect("/", 302)
		}
	}
}

func (c *AuthController) saveLoginDetails(user *User) {
	c.SetSession("user", user.Username)
}

func (controller *AuthController) Logout() {
	controller.DelSession("user")
	controller.DestroySession()
	controller.Redirect("/", 200)
}
