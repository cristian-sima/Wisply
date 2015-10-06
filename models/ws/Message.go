package websockets

// Message represents a message from client to server
type Message struct {
	Name string `json:"Name"`
	// Repository represents the ID of the repository
	Repository int         `json:"Repository"`
	Value      interface{} `json:"Value"`
}
