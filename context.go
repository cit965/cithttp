package cithttp

import (
	"encoding/json"
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

func (c *Context) Json(status int, obj any) {
	c.Res.Header().Set("Content-Type", "application/json")
	c.Res.WriteHeader(status)
	byteResult, err := json.Marshal(obj)
	if err != nil {
		c.Res.WriteHeader(500)
		return
	}
	c.Res.Write(byteResult)
}

func (c *Context) String(s string) {
	c.Res.Header().Set("Content-Type", "plain/txt")
	c.Res.WriteHeader(200)
	c.Res.Write([]byte(s))
}
