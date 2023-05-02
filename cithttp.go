package cithttp

import (
	"log"
	"net/http"
	"strings"
)

type MyHandler func(ctx *Context)

type Engine struct {
	router map[string]MyHandler
}

func (e *Engine) Run(address string) {
	log.Println("start to listen on port:", address)
	http.ListenAndServe(address, e)
}

func Default() *Engine {
	return &Engine{
		router: map[string]MyHandler{},
	}
}

func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := strings.Trim(req.URL.Path, "/")
	r, ok := e.router[path]
	if !ok {
		res.WriteHeader(404)
		res.Write([]byte("404"))
		return
	}
	r(NewContext(req, res))
}

func (e *Engine) GET(path string, h MyHandler) {
	e.router[path] = h
}
