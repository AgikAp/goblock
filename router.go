package goblock

import (
	"net/http"
	"path"
)

var allMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodHead,
	http.MethodOptions,
}

type Route struct {
	handlers   []HandlerFunc
	paramNames []string
	params     map[string]string
}

func NewRouter(handlers []HandlerFunc, paramNames []string) *Route {
	return &Route{
		handlers:   handlers,
		paramNames: paramNames,
	}
}

type IRouterGroup interface {
	Use(...HandlerFunc) IRouterGroup
	Get(string, ...HandlerFunc) IRouterGroup
	Post(string, ...HandlerFunc) IRouterGroup
	Put(string, ...HandlerFunc) IRouterGroup
	Delete(string, ...HandlerFunc) IRouterGroup
	Patch(string, ...HandlerFunc) IRouterGroup
	Options(string, ...HandlerFunc) IRouterGroup
	Head(string, ...HandlerFunc) IRouterGroup
	Connect(string, ...HandlerFunc) IRouterGroup
	All(string, ...HandlerFunc) IRouterGroup
}

type RouterGroup struct {
	goBlock   *GoBlock
	rootGroup bool
	path      string
	handlers  []HandlerFunc
}

func (r *RouterGroup) Group(path string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		goBlock:   r.goBlock,
		rootGroup: false,
		path:      r.combinePaths(path),
		handlers:  r.combineHandlers(handlers...),
	}
}

func (r *RouterGroup) Use(middlewares ...HandlerFunc) IRouterGroup {
	r.handlers = append(r.handlers, middlewares...)

	return r.this()
}

func (r *RouterGroup) Get(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	r.goBlock.registerRoute(
		http.MethodGet,
		path.Join(r.path, relativePath),
		r.combineHandlers(handlers...),
	)

	return r.this()
}

func (r *RouterGroup) Post(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	r.goBlock.registerRoute(
		http.MethodPost,
		path.Join(r.path, relativePath),
		r.combineHandlers(handlers...),
	)

	return r.this()
}

func (r *RouterGroup) Put(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	r.goBlock.registerRoute(
		http.MethodPut,
		path.Join(r.path, relativePath),
		r.combineHandlers(handlers...),
	)

	return r.this()
}

func (r *RouterGroup) Patch(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	r.goBlock.registerRoute(
		http.MethodPatch,
		path.Join(r.path, relativePath),
		r.combineHandlers(handlers...),
	)

	return r.this()
}

func (r *RouterGroup) Delete(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	r.goBlock.registerRoute(
		http.MethodDelete,
		path.Join(r.path, relativePath),
		r.combineHandlers(handlers...),
	)

	return r.this()
}

func (r *RouterGroup) Connect(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	r.goBlock.registerRoute(
		http.MethodConnect,
		path.Join(r.path, relativePath),
		r.combineHandlers(handlers...),
	)

	return r.this()
}

func (r *RouterGroup) Options(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	r.goBlock.registerRoute(
		http.MethodOptions,
		path.Join(r.path, relativePath),
		r.combineHandlers(handlers...),
	)

	return r.this()
}

func (r *RouterGroup) Head(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	r.goBlock.registerRoute(
		http.MethodHead,
		path.Join(r.path, relativePath),
		r.combineHandlers(handlers...),
	)

	return r.this()
}

func (r *RouterGroup) All(relativePath string, handlers ...HandlerFunc) IRouterGroup {
	for _, method := range allMethods {
		r.goBlock.registerRoute(
			method,
			path.Join(r.path, relativePath),
			r.combineHandlers(handlers...),
		)
	}

	return r.this()
}

func (r *RouterGroup) combinePaths(relativePath string) string {
	if relativePath == "" {
		return r.path
	}

	return path.Join(r.path, relativePath)
}

func (r *RouterGroup) combineHandlers(handlers ...HandlerFunc) []HandlerFunc {
	mergedHandlers := make([]HandlerFunc, len(r.handlers)+len(handlers))

	copy(mergedHandlers, r.handlers)
	copy(mergedHandlers[len(r.handlers):], handlers)

	return mergedHandlers
}

func (r *RouterGroup) this() IRouterGroup {
	if r.rootGroup {
		return r.goBlock
	}

	return r
}
