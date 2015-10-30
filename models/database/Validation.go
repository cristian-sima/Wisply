package database

import (
	wisply "github.com/cristian-sima/Wisply/models/adapter"
	validity "local-projects/validity"
)

var (
	rules = map[string][]string{
		"limit-min": {"Int", "min:-1"},
		"limit-max": {"Int", "min:-1"},
	}
)

func areValidLimits(min, max string) bool {
	rules := validity.ValidationRules{
		"min": rules["limit-min"],
		"max": rules["limit-max"],
	}
	result := wisply.Validate(map[string]interface{}{
		"min": min,
		"max": max,
	}, rules)

	return result.IsValid
}
