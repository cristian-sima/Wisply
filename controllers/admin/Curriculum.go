package admin

import "github.com/cristian-sima/Wisply/models/curriculum"

// Curriculum manages the operations for curriculum
type Curriculum struct {
	Controller
}

// ShowHomePage shows all the repositories
func (controller *Curriculum) ShowHomePage() {
	list := curriculum.GetAllPrograms()
	controller.Data["programs"] = list
	controller.SetCustomTitle("Admin - Curriculum")
	controller.TplNames = "site/admin/curriculum/list.tpl"
}
