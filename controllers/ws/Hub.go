package websockets

// Hub manages the connections
type Hub struct {
	// Registered connections.
	connections map[*Connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *Connection

	// Unregister requests from connections.
	unregister chan *Connection
}

func (hub *Hub) run() {
	for {
		select {
		case c := <-currentHub.register:
			currentHub.connections[c] = true
		case c := <-currentHub.unregister:
			delete(currentHub.connections, c)
			close(c.send)
		case m := <-currentHub.broadcast:
			for c := range currentHub.connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(currentHub.connections, c)
				}
			}
		}
	}
}
