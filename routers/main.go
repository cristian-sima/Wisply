// Package routers contains all the addresses of the application
package routers

import "github.com/cristian-sima/Wisply/routers/account"

func init() {
	load()
}

func load() {
	loadPublic()
	loadAdmin()
	loadDevelopers()
	account.Load()
}
