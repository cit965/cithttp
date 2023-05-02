### Cit云原生社区 web 框架

```go
func main() {
	e := Default()
	e.GET("foo", Foo)
	e.GET("bff", Bff)
	e.Run(":8888")
}
```
