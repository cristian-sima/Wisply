package education

import (
	wisply "github.com/cristian-sima/Wisply/models/adapter"
	validity "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"name":               {"String", "between_inclusive:3,300"},
		"source":             {"String", "between_inclusive:3,200"},
		"definition-content": {"String", "between_inclusive:3,1000"},
		"ka-content":         {"String", "between_inclusive:3,65535"},
		"ka-code":            {"String", "between_inclusive:2,10"},
		"ka-title":           {"String", "between_inclusive:3,100"},
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

func areValidSubjectDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name": rules["name"],
	}
	return wisply.Validate(details, rules)
}

func hasKAValidDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"ka-content": rules["ka-content"],
		"ka-source":  rules["source"],
		"ka-code":    rules["ka-code"],
		"ka-title":   rules["ka-title"],
	}
	return wisply.Validate(details, rules)
}

func hasDefinationValidDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"definition-content": rules["definition-content"],
		"definition-source":  rules["source"],
	}
	return wisply.Validate(details, rules)
}
