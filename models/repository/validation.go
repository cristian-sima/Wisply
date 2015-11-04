package repository

import (
	adapter "github.com/cristian-sima/Wisply/models/adapter"
	validity "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"name":              {"String", "between_inclusive:3,500"},
		"url":               {"String", "url", "between_inclusive:3,2083"},
		"description":       {"String", "max:1000"},
		"id":                {"Int"},
		"institution":       {"Int"},
		"wikiID":            {"String", "Regexp:^(NULL)$|^((\\d)+)$"},
		"category":          {"String", "Regexp:^(EPrints)$"},
		"program-title":     {"String", "between_inclusive:3,200"},
		"program-code":      {"String", "between_inclusive:3,10"},
		"program-ucas-code": {"String", "between_inclusive:0,20"},
		"program-level":     {"String", "Regexp:^(undergraduate)|(postgraduate)$"},
		"module-CATS":       {"String", "between_inclusive:0,5"},
		"module-year":       {"String", "between_inclusive:1,2"},
	}
)

func hasValidInstitutionInsertDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name":        rules["name"],
		"url":         rules["url"],
		"description": rules["description"],
		"institution": rules["institution"],
		"logoURL":     rules["url"],
		"wikiURL":     rules["url"],
		"wikiID":      rules["wikiID"],
	}
	return adapter.Validate(details, rules)
}

func hasValidInstitutionModifyDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name":        rules["name"],
		"description": rules["description"],
		"institution": rules["institution"],
		"logoURL":     rules["url"],
		"wikiURL":     rules["url"],
		"wikiID":      rules["wikiID"],
	}
	return adapter.Validate(details, rules)
}

func hasValidInsertDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name":        rules["name"],
		"url":         rules["url"],
		"description": rules["description"],
		"institution": rules["institution"],
		"category":    rules["category"],
		"public-url":  rules["url"],
	}
	return adapter.Validate(details, rules)
}

func hasValidModuleModifyDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"module-title": rules["program-title"],
		"module-code":  rules["program-code"],
		"module-CATS":  rules["module-CATS"],
		"module-year":  rules["module-year"],
	}
	return adapter.Validate(details, rules)
}

func hasValidProgramModifyDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"program-title":     rules["program-title"],
		"program-code":      rules["program-code"],
		"program-ucas-code": rules["program-ucas-code"],
		"program-level":     rules["program-level"],
	}
	return adapter.Validate(details, rules)
}

func isValidStatus(status string) bool {
	if status == "unverified" ||
		status == "ok" ||
		status == "verified" ||
		status == "verification-failed" ||
		status == "problems" ||
		status == "verifying" ||
		status == "updating" ||
		status == "initializing" {

		return true
	}
	return false

}

func isValidURL(URL string) *validity.ValidationResults {
	data := make(map[string]interface{})
	data["url"] = URL
	rules := validity.ValidationRules{
		"url": rules["url"],
	}
	return adapter.Validate(data, rules)
}

func isValidID(ID string) *validity.ValidationResults {
	data := make(map[string]interface{})
	data["id"] = ID
	rules := validity.ValidationRules{
		"id": rules["id"],
	}
	return adapter.Validate(data, rules)
}
