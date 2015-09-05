package controllers


type AdminController struct {
	WisplyController
}

func (c *AdminController) ShowDashboard() {
    c.Layout = "general/admin.tpl"
	c.TplNames = "general/admin/dashboard.tpl"
}