package sources

import (
	adapter "github.com/cristian-sima/Wisply/models/adapter"
	validity "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"name":        {"String", "between_inclusive:3,255"},
		"url":         {"String", "url", "between_inclusive:3,2083"},
		"description": {"String", "max:255"},
		"id":          {"Int"},
	}
)

func hasValidDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name":     rules["name"],
		"password": rules["password"],
		"email":    rules["email"],
	}
	return adapter.Validate(details, rules)
}

func isValidID(ID string) *validity.ValidationResults {
	data := make(map[string]interface{})
	data["id"] = ID
	rules := validity.ValidationRules{
		"id": rules["id"],
	}
	return adapter.Validate(data, rules)
}
