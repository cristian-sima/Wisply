package auth

import (
	. "github.com/cristian-sima/Wisply/models/adapter"
	. "github.com/cristian-sima/validity"
)

var (
	rules = map[string][]string{
		"name":          {"String", "full_name", "between_inclusive:3,25"},
		"password":      {"String", "alpha_dash", "between_inclusive:6,25"},
		"email":         {"String", "email", "between_inclusive:3,60"},
		"id":            {"Int"},
		"administrator": {"String", "Regexp:^(true)|(false)$"},
	}
)

func IsValidRegister(rawData map[string]interface{}) *ValidationResults {
	rules := ValidationRules{
		"name":     rules["name"],
		"password": rules["password"],
		"email":    rules["email"],
	}
	return Validate(rawData, rules)
}

func IsValidLogin(rawData map[string]interface{}) *ValidationResults {
	rules := ValidationRules{
		"email":    rules["email"],
		"password": rules["password"],
	}
	return Validate(rawData, rules)
}

func IsValidAdminType(email string) *ValidationResults {
	data := make(map[string]interface{})
	data["email"] = email
	rules := ValidationRules{
		"administrator": rules["administrator"],
	}
	return Validate(data, rules)
}

func IsValidEmail(email string) *ValidationResults {
	data := make(map[string]interface{})
	data["email"] = email
	rules := ValidationRules{
		"email": rules["email"],
	}
	return Validate(data, rules)
}

func IsValidId(id string) *ValidationResults {
	data := make(map[string]interface{})
	data["id"] = id
	rules := ValidationRules{
		"id": rules["id"],
	}
	return Validate(data, rules)
}
