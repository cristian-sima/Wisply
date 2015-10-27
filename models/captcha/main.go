package captcha

import "github.com/dchest/captcha"

const (
	pathToImage string = "/captcha/"
)

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
