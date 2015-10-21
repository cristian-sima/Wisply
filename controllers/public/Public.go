// Package public manages all the public pages that do not require a connection
package public

import (
	general "github.com/cristian-sima/Wisply/controllers/general"
)

// Controller can be accsed by the users who are not connected
type Controller struct {
	general.WisplyController
}
