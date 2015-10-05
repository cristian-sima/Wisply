package history

// Event encapsulates the information about an event
type Event struct {
	ID            int
	Repository    int
	Duration      float32
	Content       string
	OperationName string
	OperationType string
	Timestamp     string
}

// GUIEvent represents an event displayed by gui.
type GUIEvent struct {
	Event
	RepositoryName string
}
