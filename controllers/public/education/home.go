package education

// Home manages the operations for education
type Home struct {
	Controller
}

// ShowCurricula shows all the curricula
func (controller *Home) Display() {
	controller.SetCustomTitle("Wisply - Curricula")
	controller.LoadTemplate("home")
}
