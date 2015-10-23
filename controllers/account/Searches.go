package account

// Search is the controller which manages the operation for the search queries
type Search struct {
	Controller
}

// DisplayList List shows the list of all the searches
func (controller *Search) DisplayList() {
	list := controller.Account.GetSearches().GetAll()
	controller.Data["searches"] = list
	controller.TplNames = "site/account/search/list.tpl"
}
