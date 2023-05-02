package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	e := Default()
	e.GET("foo", Foo)
	http.ListenAndServe(":8888", e)
}

type MyHandler func(res http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]MyHandler
}

func Foo(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("你 正 在 请求 foo 函数"))
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
	r(res, req)
}

func (e *Engine) GET(path string, h MyHandler) {
	e.router[path] = h
}
