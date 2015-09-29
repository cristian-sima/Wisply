package public

import repository "github.com/cristian-sima/Wisply/models/repository"

// RepositoryController managers the operations for displaying repositories
type RepositoryController struct {
	Controller
	model repository.Model
}

// List shows all the institutions
func (controller *RepositoryController) List() {
	var exists bool
	list := controller.model.GetAllInstitutions()
	exists = (len(list) != 0)
	controller.Data["anything"] = exists
	controller.Data["institutions"] = list
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/public/institution/list.tpl"
	controller.Layout = "site/public-layout.tpl"
}

// ShowRepository shows the details regarding a repository
func (controller *RepositoryController) ShowRepository() {
	ID := controller.Ctx.Input.Param(":id")
	rep, err := repository.NewRepository(ID)
	institution := rep.GetInstitution()
	identification := rep.GetIdentification()
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["repository"] = rep
		controller.Data["institution"] = institution
		controller.Data["identification"] = identification
		controller.Layout = "site/public-layout.tpl"
		controller.TplNames = "site/public/repository/repository.tpl"
	}
}
