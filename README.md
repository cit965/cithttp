### Cit云原生社区 web 框架


如何使用，请看下面的代码：
```go
package main

import (
	"github.com/cit965/cithttp"
)

func main() {
	e := cithttp.Default()
	e.GET("foo", Foo)
	e.GET("bff", Bff)
	e.Run(":8888")
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

```
