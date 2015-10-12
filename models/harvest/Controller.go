package harvest

import "github.com/cristian-sima/Wisply/models/action"

// Controller ... represents a controller
type Controller interface {
	GetConduit() chan ProcessMessager
}

// ProcessMessager ... is a very simple message
type ProcessMessager interface {
	action.ProcessMessager
	GetRepository() int
}

// ProcessMessage encapsulates the message to communicate with controller
type ProcessMessage struct {
	action.ProcessMessage
	Repository int
}

// GetRepository returns the ID of the repository
func (message *ProcessMessage) GetRepository() int {
	return message.Repository
}

// Message can be passed into createProcessMessage in order to create a complex message from process to controller
type Message struct {
	Name  string
	Value interface{}
}
