// Package routers contains all the addresses for the application
// In case a request is made by the client, the router checks if the address
// matches any of the routes. If so, it calls the respective controller
package routers

import (
	"github.com/cristian-sima/Wisply/routers/account"
	"github.com/cristian-sima/Wisply/routers/admin"
	"github.com/cristian-sima/Wisply/routers/developer"
	"github.com/cristian-sima/Wisply/routers/public"
)

func init() {
	account.Load()
	admin.Load()
	developer.Load()
	public.Load()
}
