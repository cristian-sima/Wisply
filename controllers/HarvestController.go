package controllers

type HarvestConstructor struct {
	DefaultController
}

func (c *SourceController) ShowOptions() {
	c.TplNames = "general/harvest/options.tpl"
	c.Layout = "general/admin.tpl"
}
