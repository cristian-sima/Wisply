// Package websockets encapsulates the functionality for the Websockets connecton
package websockets

// CreateHub creates a new Hub
func CreateHub() *Hub {
	hub := &Hub{
		broadcast:   make(chan *Message, MaxMessageSize),
		Register:    make(chan *Connection, 1),
		Unregister:  make(chan *Connection, 1),
		connections: make(map[*Connection]bool),
	}
	return hub
}
