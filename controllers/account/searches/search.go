package searches

// List is the controller which manages the operation for the search queries
type List struct {
	Controller
}

// Display List shows the list of all the searches
func (controller *List) Display() {
	controller.Data["searches"] = controller.Account.GetSearches().GetAll()
	controller.SetCustomTitle("Account - Activity")
	controller.LoadTemplate("list")
}

// Clear clears the entire history
func (controller *List) Clear() {
	controller.Account.GetSearches().Clear()
	controller.ShowBlankPage()
}
