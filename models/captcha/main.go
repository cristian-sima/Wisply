package captcha

import "github.com/dchest/captcha"

const (
	pathToImage          string = "/captcha/"
	allowedTimeForAction int64  = 60 * 30 // 30 minutes

)

var maxNumberOfTimes = map[string]int{
	"default": 1,
}

var currentListOfActions List

func getAllowedNumber(page string) int {
	number, exists := maxNumberOfTimes[page]
	if !exists {
		return maxNumberOfTimes["default"]
	}
	return number
}

// New creates a new capcha and returns the ID
func New() Captcha {
	d := struct {
		id string
	}{
		captcha.New(),
	}
	return Captcha{
		id: d.id,
	}
}
