package harvest

// WisplyController represents a controller
type WisplyController interface {
	Notify(*Message)
}

// Message encapsulates the message to communicate with controller
type Message struct {
	Name       string
	Content    string
	Value      interface{}
	Repository int
}
