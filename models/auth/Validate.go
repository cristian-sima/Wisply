package auth

import (
	validity "github.com/cristian-sima/validity"
)

func ValidateSourceDetails(rawData map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name":        []string{"String", "between_inclusive:3,255"},
		"url":         []string{"String", "url", "between_inclusive:3,2083"},
		"description": []string{"String", "max:255"},
	}
	return validity.ValidateMap(rawData, rules)
}

func ValidateIndex(rawIndex string) bool {

	rawData := make(map[string]interface{})
	rawData["id"] = rawIndex

	rules := validity.ValidationRules{
		"id": []string{"Int"},
	}

	result := validity.ValidateMap(rawData, rules)

	return result.IsValid
}
