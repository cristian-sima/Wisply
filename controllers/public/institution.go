package public

import repository "github.com/cristian-sima/Wisply/models/repository"

// Institution managers the operations for displaying the institutions
// It displays the pages which are accessed by the public (no connection)
type Institution struct {
	Controller
	model repository.Model
}

// List shows all the institutions
func (controller *Institution) List() {
	var exists bool
	list := controller.model.GetAllInstitutions()
	exists = (len(list) != 0)
	controller.Data["anything"] = exists
	controller.Data["institutions"] = list
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.SetCustomTitle("Institutions")
	controller.TplNames = "site/public/institution/list.tpl"
}

// ShowInstitution shows the details regarding an institution
func (controller *Institution) ShowInstitution() {
	ID := controller.Ctx.Input.Param(":id")
	institution, err := repository.NewInstitution(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.SetCustomTitle(institution.Name)
		controller.Data["institution"] = institution
		controller.Data["repositories"] = institution.GetRepositories()
		controller.TplNames = "site/public/institution/institution.tpl"
	}
}
