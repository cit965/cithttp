package main

import (
	"fmt"
	"github.com/cit965/cithttp"
	"net/http"
	"strings"
)

func main() {
	e := Default()
	e.GET("foo", Foo)
	e.GET("bff", Bff)
	http.ListenAndServe(":8888", e)
}

type MyHandler func(ctx *cithttp.Context)

type Engine struct {
	router map[string]MyHandler
}

func Foo(ctx *cithttp.Context) {

	inlineStruct := struct {
		Sss int `json:"sss"`
	}{
		2343,
	}
	ctx.Json(200, inlineStruct)
}

func Bff(ctx *cithttp.Context) {

	ctx.String("bfff  handler")
}

func Default() *Engine {
	return &Engine{
		router: map[string]MyHandler{},
	}
}

func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("----", req.URL.Path)
	path := strings.Trim(req.URL.Path, "/")
	r, ok := e.router[path]
	if !ok {
		res.WriteHeader(404)
		res.Write([]byte("404"))
		return
	}
	r(cithttp.NewContext(req, res))
}

func (e *Engine) GET(path string, h MyHandler) {
	e.router[path] = h
}
