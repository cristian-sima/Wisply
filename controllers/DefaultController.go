package controllers

type DefaultController struct {
	WisplyController
}

func (c *DefaultController) ShowAbout() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/about.tpl"
}

func (c *DefaultController) ShowContact() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/contact.tpl"
}

func (c *DefaultController) ShowIndex() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/index.tpl"
}

func (c *DefaultController) ShowWebscience() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/webscience.tpl"
}

func (c *DefaultController) ShowSample() {
	c.TplNames = "site/default/sample.tpl"
}

func (c *DefaultController) ShowAccessibility() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/accessibility.tpl"
}
