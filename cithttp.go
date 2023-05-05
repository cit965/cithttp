package cithttp

import (
	"log"
	"net/http"
	"strings"
)

type MyHandler func(ctx *Context)

type Engine struct {
	router map[string]*Tree
}

func (e *Engine) Run(address string) {
	log.Println("start to listen on port:", address)
	http.ListenAndServe(address, e)
}

func Default() *Engine {

	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
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
	r := handlerMap.FindHandler(path)
	if r == nil {
		res.WriteHeader(404)
		log.Println("没有注册handler")
		return
	}
	r(NewContext(req, res))
}

type Group struct {
	engin  *Engine
	prefix string
}

func (e *Engine) GET(path string, h MyHandler) {
	e.router["GET"].AddRouter(path, h)
}

func (e *Engine) POST(path string, h MyHandler) {
	e.router["POST"].AddRouter(path, h)
}

func NewGroup(e *Engine, prefix string) *Group {
	return &Group{
		engin:  e,
		prefix: prefix,
	}
}
func (e *Engine) Group(prefix string) *Group {

	return NewGroup(e, prefix)
}

func (g *Group) GET(path string, h MyHandler) {
	g.engin.GET(g.prefix+path, h)
}

func (g *Group) POST(path string, h MyHandler) {
	g.engin.POST(g.prefix+path, h)
}
