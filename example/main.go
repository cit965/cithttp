package main

import (
	"github.com/cit965/cithttp"
)

func main() {
	r := cithttp.Default()
	r.GET("/foo", cithttp.Recovery(), cithttp.Log(), FooControllerHandler)
	g := r.Group("/boo")
	g.Use(cithttp.Log())
	{
		g.GET("/hello", FooControllerHandler)
		g.GET("/xx/:id", FooControllerHandler)
	}
	r.Run(":8000")

}

type R struct {
	Name string
}

func FooControllerHandler(c *cithttp.Context) {
	var r R
	c.BindJson(&r)

	c.Json(200, r)
}
