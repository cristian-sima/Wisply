package developer

import (
	wisply "github.com/cristian-sima/Wisply/models/adapter"
	validity "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"description": {"String", "between_inclusive:3,65535"},
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
