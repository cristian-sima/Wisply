package wisply

// Collection is used to group resources
type Collection struct {
	ID          int    `json:"ID"`
	Spec        string `json:"Spec"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Repository  int    `json:"Repository"`
}
