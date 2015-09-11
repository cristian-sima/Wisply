package routers

import (
	"github.com/astaxie/beego"
	"github.com/cristian-sima/Wisply/controllers"
)

func init() {

	beego.Router("/", &controllers.DefaultController{}, "*:ShowIndex")
	beego.Router("/about", &controllers.DefaultController{}, "*:ShowAbout")
	beego.Router("/contact", &controllers.DefaultController{}, "*:ShowContact")
	beego.Router("/webscience", &controllers.DefaultController{}, "*:ShowWebscience")
	beego.Router("/sample", &controllers.DefaultController{}, "*:ShowSample")
	beego.Router("/accessibility", &controllers.DefaultController{}, "*:ShowAccessibility")

	// ----------------------------- Authentification --------------------------------------

	authNamespace := beego.NewNamespace("/auth",
		beego.NSNamespace("/login",
			beego.NSRouter("", &controllers.AuthController{}, "GET:ShowLoginForm"),
			beego.NSRouter("", &controllers.AuthController{}, "POST:LoginUser"),
		),
		beego.NSNamespace("/register",
			beego.NSRouter("", &controllers.AuthController{}, "GET:ShowRegisterForm"),
			beego.NSRouter("", &controllers.AuthController{}, "POST:CreateNewUser"),
		),
		beego.NSNamespace("/logout",
			beego.NSRouter("", &controllers.AuthController{}, "POST:Logout"),
		),
	)

	// ----------------------------- Admin --------------------------------------

	sourcesNamespace := beego.NSNamespace("/sources",
		beego.NSRouter("", &controllers.SourceController{}, "*:ListSources"),
		beego.NSNamespace("/add",
			beego.NSRouter("", &controllers.SourceController{}, "GET:AddNewSource"),
			beego.NSRouter("", &controllers.SourceController{}, "POST:InsertSource"),
		),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &controllers.SourceController{}, "GET:Modify"),
			beego.NSRouter(":id", &controllers.SourceController{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &controllers.SourceController{}, "POST:Delete"),
		),
	)

	usersNamespace := beego.NSNamespace("/users",
		beego.NSRouter("", &controllers.UserController{}, "*:ListUsers"),
		beego.NSNamespace("/modify",
			beego.NSRouter(":id", &controllers.UserController{}, "GET:Modify"),
			beego.NSRouter(":id", &controllers.UserController{}, "POST:Update"),
		),
		beego.NSNamespace("/delete",
			beego.NSRouter(":id", &controllers.UserController{}, "POST:Delete"),
		),
	)

	adminNamespace :=
		beego.NewNamespace("/admin",
			beego.NSRouter("", &controllers.AdminController{}, "*:ShowDashboard"),
			sourcesNamespace,
			usersNamespace,
		)

	// register namespace
	beego.AddNamespace(authNamespace)
	beego.AddNamespace(adminNamespace)

}
