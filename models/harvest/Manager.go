package harvest

// VerificationManager manages the operations for modification
type VerificationManager struct {
	WisplyManager
}

// Start starts the manager
func (manager *VerificationManager) Start() {

}

// GetName returns the name of the manager
func (manager *VerificationManager) GetName() string {
	return "Verification"
}

// End finishes is fired when the manager has finished the work
func (manager *VerificationManager) End() {
}

// Notify is called by a harvest repository with a message
func (manager *VerificationManager) Notify(message *Message) {

}
