package goblock

import (
	"net/http"
	"sync"
	"time"
)

type GoBlock struct {
	pool sync.Pool
}

func New() *GoBlock {
	g := &GoBlock{}

	g.pool.New = func() interface{} {
		return NewContext()
	}

	return g
}

func (g *GoBlock) handlerRequest(c *Context) {
	c.Json(200, map[string]interface{}{
		"message": "Hello, World",
	})
}

func (g *GoBlock) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logRequest(r.Method, r.URL.Path, time.Since(start))
	})
}

func (g *GoBlock) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := g.pool.Get().(*Context)

	c.reset(w, r)
	g.handlerRequest(c)

	g.pool.Put(c)
}

func (g *GoBlock) Handler() http.Handler {
	return g.loggerMiddleware(g)
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
