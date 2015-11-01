package education

import (
	wisply "github.com/cristian-sima/Wisply/models/adapter"
	validity "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"name": {"String", "between_inclusive:3,300"},
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

func areValidProgramDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name": rules["name"],
	}
	return wisply.Validate(details, rules)
}
