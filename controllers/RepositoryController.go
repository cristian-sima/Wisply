package controllers

import (
	"fmt"

	oai "github.com/cristian-sima/Wisply/models/oai"
)

// RepositoryController It manages the operations for sources (list, delete, add)
type RepositoryController struct {
	AdminController
}

func dump(resp *oai.Response) {
	fmt.Printf("%#v\n", resp)

}

// Test It shows all the sources
func (controller *RepositoryController) Test() {
	// Print the OAI Response object to stdout

	req := (&oai.Request{
		BaseURL: "http://www.edshare.soton.ac.uk/cgi/oai2",
		Verb:    "Identify",
	})
	req.Harvest(dump)

	controller.TplNames = "site/source/list.tpl"
	controller.Layout = "site/admin.tpl"

}
