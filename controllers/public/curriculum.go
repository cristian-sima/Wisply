package public

import "github.com/cristian-sima/Wisply/models/curriculum"

// Curriculum manages the operations for curriculum
type Curriculum struct {
	Controller
}

// ShowCurricula shows all the curricula
func (controller *Curriculum) ShowCurricula() {
	controller.SetCustomTitle("Wisply - Curricula")
	controller.TplNames = "site/public/curriculum/curricula.tpl"
}

// ShowProgram shows the dashboard for a program
func (controller *Curriculum) ShowProgram() {
	controller.loadProgramToTemplate()
	controller.TplNames = "site/public/curriculum/program/home.tpl"
}

func (controller *Curriculum) loadProgramToTemplate() *curriculum.Program {
	ID := controller.Ctx.Input.Param(":id")
	program, err := curriculum.NewProgram(ID)
	if err != nil {
		controller.Abort("databaseError")
		return program
	}
	controller.Data["program"] = program
	controller.SetCustomTitle(program.GetName() + " Curriculum")
	return program
}
