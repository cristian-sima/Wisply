package harvest

import "fmt"

// WisplyController represents a controller
type WisplyController interface {
	Notify(*Message)
}

// WisplyProcessInterface ... defines the methods which must be implemented by a process
type WisplyProcessInterface interface {
	Start()
	ManagerFinished()
}

// Log manages the operations for displaying information
type Log struct {
}

func show(message string) {
	fmt.Println("<-->  " + message)
}

// WisplyProcess is a basic process. A process does a series of actions using managers
type WisplyProcess struct {
	Log
	name       string
	controller *WisplyController
}

func (process *WisplyProcess) log(message string) {
	show(process.getType() + " " + process.name + ": " + message)
}

// GetController returns the reference to the controller which manages the process
func (process *WisplyProcess) GetController() *WisplyController {
	return process.controller
}

func (process *WisplyProcess) getType() string {
	return "Process"
}

// SetName sets the name of a process
func (process *WisplyProcess) SetName(name string) {
	process.name = name
}

// ManagerInterface ... defines the set of methods which must be implemented by a harvest manager
type ManagerInterface interface {
	Start()
	End()
	Save()
}

// WisplyManager represents a general manager
type WisplyManager struct {
	process *WisplyProcess
}

// GetProcess returns the proces of the manager
func (manager *WisplyManager) GetProcess() *WisplyProcess {
	return manager.process
}
