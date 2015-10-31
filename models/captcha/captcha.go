package captcha

// Captcha represents a image which is used for brute force attack
type Captcha struct {
	id string
}

// GetID returns the ID of the catcha
func (captcha Captcha) GetID() string {
	return captcha.id
}

// GetImageURL returns the path to the image
func (captcha Captcha) GetImageURL() string {
	return pathToImage + captcha.id + ".png"
}
