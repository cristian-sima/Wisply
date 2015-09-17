package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/cristian-sima/Wisply/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestPageHome(t *testing.T) {
	url := "/"
	name := "Home"
	checkPage(t, url, name)
}

func TestPageSample(t *testing.T) {
	url := "/sample"
	name := "Sample"
	checkPage(t, url, name)
}

func TestPageAbout(t *testing.T) {
	url := "/about"
	name := "About"
	checkPage(t, url, name)
}

func TestPageWebscience(t *testing.T) {
	url := "/webscience"
	name := "Webscience"
	checkPage(t, url, name)
}

func TestPageContact(t *testing.T) {
	url := "/contact"
	name := "Contact"
	checkPage(t, url, name)
}

func TestPageLogin(t *testing.T) {
	url := "/auth/login"
	name := "Login"
	checkPage(t, url, name)
}

func TestPageRegister(t *testing.T) {
	url := "/auth/register"
	name := "Register"
	checkPage(t, url, name)
}

func TestPageHelp(t *testing.T) {
	url := "/help"
	name := "Help"
	checkPage(t, url, name)
}

func TestPageAccesibility(t *testing.T) {
	url := "/accessibility"
	name := "Accessibility"
	checkPage(t, url, name)
}

// TestMain is a sample to run an endpoint test
func checkPage(t *testing.T, url string, name string) {
	r, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "Test", name)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}
