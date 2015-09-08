package test

import (
	"github.com/astaxie/beego"
	_ "github.com/cristian-sima/Wisply/routers"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
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

func checkPage(t *testing.T, url string, pageName string) {
	r, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("Test static page: ", pageName)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("The page "+pageName+" can be accessed", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The page should not be empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
