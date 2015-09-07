package controllers

type DefaultController struct {
	WisplyController
}

func (c *DefaultController) ShowAboutPage() {
	c.Layout = "general/layout.tpl"
	c.TplNames = "general/page/about.tpl"
}

func (c *DefaultController) ShowContactPage() {
	c.Layout = "general/layout.tpl"
	c.TplNames = "general/page/contact.tpl"
}

func (c *DefaultController) ShowIndexPage() {
	c.Layout = "general/layout.tpl"
	c.TplNames = "general/page/index.tpl"
}

func (c *DefaultController) ShowWebsciencePage() {
	c.Layout = "general/layout.tpl"
	c.TplNames = "general/page/webscience.tpl"
}

func (c *DefaultController) ShowSamplePage() {
	c.TplNames = "general/sample.tpl"
}
