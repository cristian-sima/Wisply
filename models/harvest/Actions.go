package harvest

// Actions contains the id of the actions
var Actions = map[string]int{
	"testing":     3,
	"identifying": 4,
}

// Action represents the state (finish) and the number
type Action struct {
	Finished bool `json:"Finished"`
	Number   int  `json:"Number"`
}
