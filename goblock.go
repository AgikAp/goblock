package goblock

import (
	"net/http"
	"sync"
	"time"
)

type HandlerFunc func(*Context)

type Config struct {
}

type GoBlock struct {
	RouterGroup

	pool   sync.Pool
	routes *RouterTree
}

func New(config ...Config) *GoBlock {
	g := &GoBlock{
		RouterGroup: RouterGroup{
			path:     "/",
			handlers: nil,
		},
	}

	g.RouterGroup.goBlock = g
	g.RouterGroup.rootGroup = true

	g.pool.New = func() interface{} {
		return NewContext()
	}
	g.routes = NewRouterTree()

	return g
}

func (g *GoBlock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := g.pool.Get().(*Context)

	c.reset(w, r)
	g.handleRequest(c)

	g.pool.Put(c)
}

func (g *GoBlock) Handler() http.Handler {
	return g
}

func (g *GoBlock) Listen(addr string) (err error) {
	printLogo()

	err = http.ListenAndServe(addr, g.Handler())
	return
}

func (g *GoBlock) ListenTLS(addr string, certFile string, keyFile string) (err error) {
	printLogo()

	err = http.ListenAndServeTLS(addr, certFile, keyFile, g.Handler())
	return
}

func (g *GoBlock) registerRoute(method string, path string, handlers []HandlerFunc) {
	g.routes.Insert(method, path, handlers)
}

func (g *GoBlock) handleRequest(c *Context) {
	start := time.Now()

	defer func() {
		if r := recover(); r != nil {
			switch err := r.(type) {
			case *HttpError:
				c.Json(err.StatusCode, G{
					"message": err.Message,
				})
			default:
				c.Json(http.StatusInternalServerError, G{
					"message": "Internal server error",
				})
			}
		}
	}()

	method := c.Request.Method
	path := c.Request.URL.Path

	route, err := g.routes.Search(method, path)
	if err != nil {
		panic(err)
	}

	c.setParams(route.params)

	for _, handler := range route.handlers {
		handler(c)
	}

	logRequest(method, path, c.Writer.statusCode, time.Since(start))
}
