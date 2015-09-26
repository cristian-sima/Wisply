package remote

import "net/http"
import local "github.com/cristian-sima/Wisply/models/repository"

// Standard ... defines the method to be implemented by a standard (remote repository)
type Standard interface {
	StartProcess()
	GetLocalRepository() *local.Repository
}

type HarvestController interface {
	Notify(*Message)
}

// Message encapsulates the message to communicate with controller
type Message struct {
	Name       string
	Content    string
	Value      string
	Repository int
}

// Represents a remote repository
type Repository struct {
	Controller HarvestController
}

func (repository *Repository) SetController(controller HarvestController) {
	repository.Controller = controller
}

func validateURL(URL string) bool {
	var isOk bool
	isOk = true
	request, err := http.Get(URL)
	if request == nil || err != nil {
		isOk = false
	} else if http.StatusOK != request.StatusCode {
		isOk = false
	}
	return isOk
}
