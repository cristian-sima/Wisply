package auth

import (
	validity "github.com/cristian-sima/validity"

	wisply "github.com/cristian-sima/Wisply/models/adapter"
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

func isValidRegister(rawData map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"name":     rules["name"],
		"password": rules["password"],
		"email":    rules["email"],
	}
	return wisply.Validate(rawData, rules)
}

func isValidLogin(rawData map[string]interface{}) *validity.ValidationResults {
	rules := validity.ValidationRules{
		"email":    rules["email"],
		"password": rules["password"],
	}
	return wisply.Validate(rawData, rules)
}

func isValidAdminType(email string) *validity.ValidationResults {
	data := make(map[string]interface{})
	data["email"] = email
	rules := validity.ValidationRules{
		"administrator": rules["administrator"],
	}
	return wisply.Validate(data, rules)
}

func isValidEmail(email string) *validity.ValidationResults {
	data := make(map[string]interface{})
	data["email"] = email
	rules := validity.ValidationRules{
		"email": rules["email"],
	}
	return wisply.Validate(data, rules)
}

func isValidID(id string) *validity.ValidationResults {
	data := make(map[string]interface{})
	data["id"] = id
	rules := validity.ValidationRules{
		"id": rules["id"],
	}
	return wisply.Validate(data, rules)
}
