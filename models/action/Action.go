package action

// Actioner ... defines the set of methods to be implemented by an action
type Actioner interface {
	Go()
	Finish()
}

// Action is the most basic type. It has a starting and ending timestamps and a content
type Action struct {
	Actioner
	ID      int    `json:"ID"`
	Start   string `json:"Start"`
	End     string `json:"End"`
	Content string `json:"Content"`
}
