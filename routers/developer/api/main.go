// Package api contains the routes for the API requests
package api

import "github.com/astaxie/beego"

// Get returns the Namespace for API
func Get() func(*beego.Namespace) {
	ns := beego.NSNamespace("/api",
		getRepository(),
		getSearch(),
	)
	return ns
}
