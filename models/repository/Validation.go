package repository

import (
	adapter "github.com/cristian-sima/Wisply/models/adapter"
	validity "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"name":        {"String", "between_inclusive:3,500"},
		"url":         {"String", "url", "between_inclusive:3,2083"},
		"description": {"String", "max:1000"},
		"id":          {"Int"},
		"institution": {"Int"},
		"wikiID":      {"String", "Regexp:^(NULL)$|^((\\d)+)$"},
		"category":    {"String", "Regexp:^(EPrints)$"},
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

func hasValidModificationDetails(details map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name":        rules["name"],
		"description": rules["description"],
		"url":         rules["url"],
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
