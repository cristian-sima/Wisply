package sources

import (
	. "github.com/cristian-sima/Wisply/models/adapter"
	. "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"name":        {"String", "between_inclusive:3,255"},
		"url":         {"String", "url", "between_inclusive:3,2083"},
		"description": {"String", "max:255"},
		"id":          {"Int"},
	}
)

func HasValidDetails(details map[string]interface{}) *ValidationResults {
	rules := ValidationRules{
		"name":     rules["name"],
		"password": rules["password"],
		"email":    rules["email"],
	}
	return Validate(details, rules)
}

func IsValidId(id string) *ValidationResults {
	data := make(map[string]interface{})
	data["id"] = id
	rules := ValidationRules{
		"id": rules["id"],
	}
	return Validate(data, rules)
}
