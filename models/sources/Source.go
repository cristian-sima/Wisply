package sources

import validity "github.com/connor4312/validity"

type Source struct {
    Id          int
    Name        string
    Url         string
    Description string
}

func CheckSourceDetails(rawData map[string]interface{}) *validity.ValidationResults {
    
    rules := validity.ValidationRules{
        "name": []string{"String", "between:1,255"},
        "url": []string{"String", "url", "between:1,2083"},
        "description": []string{"String", "max:255"},
    }

    return validity.ValidateMap(rawData, rules)
}

func CheckIndex(rawIndex string) bool {
    
    rawData := make(map[string]interface{})
    rawData["id"] = rawIndex

    rules := validity.ValidationRules{
        "id": []string{"Int"},
    }

    result := validity.ValidateMap(rawData, rules)

	return result.IsValid
}