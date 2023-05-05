package main

import (
	"github.com/cit965/cithttp"
	"time"
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

func FooControllerHandler(c *cithttp.Context) {
	time.Sleep(time.Second * 3)
	c.Json(200, "success")
}
