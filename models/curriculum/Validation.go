package curriculum

import (
	wisply "github.com/cristian-sima/Wisply/models/adapter"
	validity "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"name": {"String", "between_inclusive:3,100"},
	}
)

func isValidName(name string) *validity.ValidationResults {
	rawData := make(map[string]interface{})
	rawData["name"] = name
	rules := validity.ValidationRules{
		"name": rules["name"],
	}
	return wisply.Validate(rawData, rules)
}
