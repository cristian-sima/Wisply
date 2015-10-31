package action

// Controller works with many processes
type Controller interface {
	GetConduit() chan ProcessMessager
}
