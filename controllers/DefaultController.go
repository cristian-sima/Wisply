package controllers

type DefaultController struct {
	WisplyController
}

func (c *DefaultController) ShowAboutPage() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/about.tpl"
}

func (c *DefaultController) ShowContactPage() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/contact.tpl"
}

func (c *DefaultController) ShowIndexPage() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/index.tpl"
}

func (c *DefaultController) ShowWebsciencePage() {
	c.Layout = "site/layout.tpl"
	c.TplNames = "site/default/webscience.tpl"
}

func (c *DefaultController) ShowSamplePage() {
	c.TplNames = "site/default/sample.tpl"
}
