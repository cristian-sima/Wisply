// Package word contains the objects and functions for processing words
package word

// GetFilterList returns the list of words which are rejected
func GetFilterList(listName string) []string {
	list := []string{}
	switch listName {
	case "articles":
		list = articles
		break
	case "conjunctions":
		list = conjunctions
		break
	case "education":
		list = education
		break
	case "prepositions":
		list = prepositions
		break
	case "pronouns":
		list = pronouns
		break
	case "redundant":
		list = redundantWords
		break
	default:
		panic("No list for " + listName)
	}
	return list
}
