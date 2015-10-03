package history

// Event encapsulates the information about an event
type Event struct {
	Content       string
	OperationName string
	OperationType string
	Repository    int
}
