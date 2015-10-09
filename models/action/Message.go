package action

// Message is a simple type of message used by actions to communicate between them
type Message struct {
	Name  string
	Value interface{}
}

// GetName returns the name of the message
func (message *Message) GetName() string {
	return message.Name
}

// GetValue returns the value of the message
func (message *Message) GetValue() interface{} {
	return message.Value
}

// Messager ... defines the methods which must be implemented by a message
type Messager interface {
	GetName() string
	GetValue() interface{}
}

// OperationMessager ... defines the methods which must be implemented by a message sent from operation to process
type OperationMessager interface {
	Messager
	GetOperation() *Operation
}

// OperationMessage represents a message sent from operation to process
type OperationMessage struct {
	*Message
	Operation *Operation
}

// GetOperation returns the operation
func (message *OperationMessage) GetOperation() *Operation {
	return message.Operation
}
