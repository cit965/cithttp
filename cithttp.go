package cithttp

import (
	"log"
	"net/http"
	"strings"
)

type MyHandler func(ctx *Context)

type Engine struct {
	router map[string]map[string]MyHandler
}

func (e *Engine) Run(address string) {
	log.Println("start to listen on port:", address)
	http.ListenAndServe(address, e)
}

func Default() *Engine {

	router := map[string]map[string]MyHandler{}
	getRouter := map[string]MyHandler{}
	postRouter := map[string]MyHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	return &Engine{
		router,
	}
}

func (e *Engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := strings.Trim(req.URL.Path, "/")
	method := strings.ToUpper(req.Method)
	handlerMap, ok := e.router[method]
	if !ok {
		log.Println("not found method")
		return
	}
	r, ok := handlerMap[path]
	if !ok {
		res.WriteHeader(404)
		res.Write([]byte("404"))
		return
	}
	r(NewContext(req, res))
}

func (e *Engine) GET(path string, h MyHandler) {
	e.router["GET"][path] = h
}

func (e *Engine) POST(path string, h MyHandler) {
	e.router["POST"][path] = h
}
