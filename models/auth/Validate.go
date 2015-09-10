package auth

import (
	"fmt"
	validity "github.com/cristian-sima/validity"
)

func ValidateNewUserDetails(rawData map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"username": []string{"String", "alpha_dash", "between_inclusive:3,25"},
		"password": []string{"String", "alpha_dash", "between_inclusive:6,25"},
		"email":    []string{"String", "email", "between_inclusive:3,25"},
	}
	return validity.ValidateMap(rawData, rules)
}

func ValidateLoginDetails(rawData map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"username": []string{"String", "alpha_dash", "between_inclusive:3,25"},
		"password": []string{"String", "alpha_dash", "between_inclusive:6,25"},
	}
	return validity.ValidateMap(rawData, rules)
}

func ValidateModify(rawData map[string]interface{}) *validity.ValidationResults {
	fmt.Println(rawData)
	rules := validity.ValidationRules{
		"administrator": []string{"String", "Regexp:^(true)|(false)$"},
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
