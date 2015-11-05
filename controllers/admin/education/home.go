package education

import "github.com/cristian-sima/Wisply/models/analyse"

// Home represents manages the home page for education
type Home struct {
	Controller
}

// Display shows the home page for the education
func (controller *Home) Display() {
	controller.Data["analyses"] = analyse.GetAll()
	controller.SetCustomTitle("Admin - Education")
	controller.LoadTemplate("home")
}

// Analyse analyses the data from an institution
func (controller *Home) Analyse() {
	analyse := analyse.CreateAnalyser()
	controller.ShowBlankPage()
	controller.RemoveLayout()
	go analyse.Start()
	controller.Redirect("/admin/education", 303)
}

// DeleteAnalyser removes an analyser from database
func (controller *Home) DeleteAnalyser() {
	ID := controller.Ctx.Input.Param(":analyser")
	analyser := analyse.NewAnalyser(ID)
	analyser.Delete()
	controller.ShowBlankPage()
	controller.RemoveLayout()
}
