package cithttp

import (
	"net/http"
	"time"
)

type Context struct {
	Req *http.Request
	Res http.ResponseWriter
}

func NewContext(r *http.Request, res http.ResponseWriter) *Context {
	return &Context{
		Req: r,
		Res: res,
	}
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Req.Context().Deadline()
}
func (c *Context) Done() <-chan struct{} {
	return c.Req.Context().Done()
}
func (c *Context) Err() error {
	return c.Req.Context().Err()
}

func (c *Context) Value(key any) any {
	return c.Req.Context().Err()
}
