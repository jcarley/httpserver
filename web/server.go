package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MiddlewareFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

type HttpServer struct {
	Pipelines map[string][]MiddlewareFunc
	Port      string
	Router    *mux.Router
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		Pipelines: make(map[string][]MiddlewareFunc),
	}
}

func (this *HttpServer) Pipeline(name string, middleware ...MiddlewareFunc) {
	tmp, ok := this.Pipelines[name]
	if !ok {
		tmp = []MiddlewareFunc{}
	}
	for _, m := range middleware {
		tmp = append(tmp, m)
	}
	this.Pipelines[name] = tmp
}
