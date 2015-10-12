package harvest

import "github.com/cristian-sima/Wisply/models/action"

// Controller ... represents a controller
type Controller interface {
	GetConduit() chan ProcessMessager
}

// ProcessMessager ... is a very simple message
type ProcessMessager interface {
	action.ProcessMessager
}

// MessageX encapsulates the message to communicate with controller
type MessageX struct {
	Name       string
	Content    string
	Value      interface{}
	Repository int
}
