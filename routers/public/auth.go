package public

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers/public"
)

func getAuth() *beego.Namespace {
	ns := beego.NewNamespace("auth",
		beego.NSNamespace("/login",
			beego.NSRouter("", &public.Auth{}, "GET:ShowLoginPage"),
			beego.NSRouter("", &public.Auth{}, "POST:LoginAccount"),
		),
		beego.NSNamespace("/register",
			beego.NSRouter("", &public.Auth{}, "GET:ShowRegisterForm"),
			beego.NSRouter("", &public.Auth{}, "POST:CreateNewAccount"),
		),
		beego.NSNamespace("/logout",
			beego.NSRouter("", &public.Auth{}, "POST:Logout"),
		),
	)
	return ns
}
