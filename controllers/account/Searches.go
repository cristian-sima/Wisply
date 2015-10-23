package account

// Search is the controller which manages the operation for the search queries
type Search struct {
	Controller
}

// DisplayHistory List shows the list of all the searches
func (controller *Search) DisplayHistory() {
	list := controller.Account.GetSearches().GetAll()
	controller.Data["searches"] = list
	controller.TplNames = "site/account/search/list.tpl"
}

// ClearHistory clears the entire history
func (controller *Search) ClearHistory() {
	controller.Account.GetSearches().Clear()
	controller.TplNames = "site/blank.tpl"
}
