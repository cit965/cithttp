package cithttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Context is the most important part of ctp. It allows us get http request and response
type Context struct {
	Request  *http.Request
	Writer   http.ResponseWriter
	handlers []HandlerFunc
	index    int
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		Request: r,
		Writer:  w,
		index:   -1,
	}
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.Request.Context().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.Request.Context().Done()
}

func (ctx *Context) Err() error {
	return ctx.Request.Context().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.Request.Context().Value(key)
}

func (c *Context) String(s string) {
	c.Writer.Header().Set("Content-Type", "plain/txt")
	c.Writer.WriteHeader(200)
	c.Writer.Write([]byte(s))
}

func (ctx *Context) Json(status int, obj interface{}) {

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.Writer.WriteHeader(500)
		log.Print(err)
	}
	ctx.Writer.Write(byt)
}

// 将body文本解析到obj结构体中
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.Request != nil {
		// 读取文本
		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			return err
		}
		// 重新填充request.Body，为后续的逻辑二次读取做准备
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// 解析到obj结构体中
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

// 核心函数，调用context的下一个函数
func (ctx *Context) Next() {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		ctx.handlers[ctx.index](ctx)
	}
}

func (ctx *Context) SetHandlers(handlers []HandlerFunc) {
	ctx.handlers = handlers
}
