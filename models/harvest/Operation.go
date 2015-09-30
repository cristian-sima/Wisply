package harvest

import "fmt"

// Controller represents a controller
type Controller interface {
	Notify(*Message)
}

type Log struct {
}

func (logger *Log) show(message string) {
	fmt.Println("<-->  " + message)
}

type WisplyProcess struct {
	Log
	name       string
	controller *Controller
}

func (process *WisplyProcess) log(message string) {
	process.Log.show(process.getType() + " " + process.name + ": " + message)
}

func (process *WisplyProcess) GetController() *Controller {
	return process.controller
}

func (process *WisplyProcess) getType() string {
	return "Process"
}

func (operation *WisplyProcess) SetName(name string) {
	operation.name = name
}
