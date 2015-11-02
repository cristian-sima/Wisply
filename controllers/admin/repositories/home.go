package repositories

import "github.com/cristian-sima/Wisply/models/repository"

// Home manages the default page for repositories
type Home struct {
	Controller
}

// Display shows a table with all the repositories
func (controller *Home) Display() {
	controller.SetCustomTitle("Admin - Repositories")
	controller.Data["repositories"] = repository.GetAllRepositories()
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.LoadTemplate("home")
}
