package subject

import (
	"github.com/cristian-sima/Wisply/models/analyse"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Subject manages the operations for curriculum.
type Subject struct {
	Controller
}

// Display shows the public page for a subject of study
func (controller *Subject) Display() {
	subject := controller.GetSubject()
	controller.SetCustomTitle(subject.GetName())
	controller.Data["institutions"] = repository.GetAllInstitutions()
	controller.LoadTemplate("home")
	controller.Data["subjectAnalyses"] = analyse.GetSubjectAnalysersBySubject(subject.GetID())

}
