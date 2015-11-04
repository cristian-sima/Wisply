// Package captcha is an adaptation of the package github.com/dchest/captcha
// to the framework beego. It includes more functionality such as allowed number
// of times before the captcha message is shown
package captcha

import "github.com/dchest/captcha"

const (
	pathToImage string = "/captcha/"
)

type configuration struct {
	count int
	time  int64
}

var (
	configurations = map[string]configuration{
		"default": configuration{
			count: 5,
			time:  (60 * 30),
		},
		"tools": configuration{
			count: 700,
			time:  (60 * 60 * 24), // 1 day
		},
	}
	currentListOfActions List
)

func getAllowedTime(page string) int64 {
	info, exists := configurations[page]
	if !exists {
		return configurations["default"].time
	}
	return info.time
}

func getAllowedNumber(page string) int {
	info, exists := configurations[page]
	if !exists {
		return configurations["default"].count
	}
	return info.count
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
