package api

import (
	wisply "github.com/cristian-sima/Wisply/models/adapter"
	validity "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"description": {"String", "full_name", "between_inclusive:3,1000"},
	}
)

func isValidDescription(description string) bool {
	rules := validity.ValidationRules{
		"description": rules["description"],
	}
	rawData := make(map[string]interface{})
	rawData["description"] = description

	return wisply.Validate(rawData, rules).IsValid
}
