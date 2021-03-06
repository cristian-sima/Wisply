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

// OperationMessager ... defines the methods which must be implemented by
// a message sent from operation to process
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

// ProcessMessager ... defines the methods which must be implemented by
// a message sent from process to controller
type ProcessMessager interface {
	Messager
	GetProcess() Processer
}

// ProcessMessage represents a message sent from process to its controller
type ProcessMessage struct {
	*Message
	Process Processer
}

// GetProcess returns the process
func (message *ProcessMessage) GetProcess() Processer {
	return message.Process
}

// Processer represents an interface for a process
type Processer interface {
}
