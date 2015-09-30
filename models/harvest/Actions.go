package harvest

// Actions contains the id of the actions
var Actions = map[string]int{
	"testing":     3,
	"identifying": 4,
	"harvesting":  5,
}

// Action represents the state (finish) and the number
type Action struct {
	Finished  bool   `json:"Finished"`
	Type      string `json:"Type"`
	Count     int    `json:"Number"`
	IsCurrent bool   `json:"IsCurrent"`
}

// Update changes the count of action
func (action *Action) Update(newValue int) {
	action.Count = action.Count + newValue
}

// Finish marks the action as finished
func (action *Action) Finish() {
	action.Finished = true
	action.IsCurrent = false
}
