package public

import InstitutionModel "github.com/cristian-sima/Wisply/models/institution"

// InstitutionController managers the operations for displaying
type InstitutionController struct {
	Controller
	model InstitutionModel.Model
}

// List shows all the institutions
func (controller *InstitutionController) List() {
	var exists bool
	list := controller.model.GetAll()
	exists = (len(list) != 0)
	controller.Data["anything"] = exists
	controller.Data["institutions"] = list
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/public/institution/list.tpl"
	controller.Layout = "site/public-layout.tpl"
}

// ShowInstitution shows the details regarding an institution
func (controller *InstitutionController) ShowInstitution() {
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.Layout = "site/public-layout.tpl"
	controller.TplNames = "site/public/institution/form.tpl"
}
