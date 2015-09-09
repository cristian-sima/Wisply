package controllers

type AuthController struct {
	DefaultController
}

func (c *AuthController) ShowLoginForm() {

	c.TplNames = "general/auth/login.tpl"
	c.Layout = "general/layout.tpl"
}
