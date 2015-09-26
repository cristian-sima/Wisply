package websockets

import "time"

const (
	// WriteWaitPeriod Time allowed to write a message to the client.
	WriteWaitPeriod = 10 * time.Second

	// ReadWaitPeriod Time allowed to read the next message from the client.
	ReadWaitPeriod = 6000 * time.Second

	// PingInterval end pings to client with this period. Must be less than readWait.
	PingInterval = (ReadWaitPeriod * 9) / 10

	// MaxMessageSize sets the maximum message size allowed from client.
	MaxMessageSize = 512
)
