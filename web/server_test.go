package web_test

import (
	"net/http"
	"testing"

	"github.com/jcarley/httpserver/web"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServerSetup(t *testing.T) {

	MyMiddleware1 := func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}
	MyMiddleware2 := func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}

	Convey("Can add middleware to pipelines", t, func() {
		server := web.NewHttpServer()
		server.Pipeline("browser", MyMiddleware1, MyMiddleware2)
		So(len(server.Pipelines["browser"]), ShouldEqual, 2)
	})

	Convey("Setup routes for a pipeline", t, func() {
		server := web.NewHttpServer()
		server.Pipeline("browser", MyMiddleware1, MyMiddleware2)
		server.Scope("/").PipeThrough("browser")
		So(server.Scope("/").Pipeline(), ShouldEqual, "browser")
	})

}
