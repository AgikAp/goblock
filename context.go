package goblock

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Request *http.Request
	Writer  *ResponseWriter
	params  map[string]string
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *Context) Param(key string) string {
	return c.params[key]
}

func (c *Context) Json(statusCode int, data map[string]interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(statusCode)

	err := json.NewEncoder(c.Writer).Encode(data)
	c.HandleError(err)
}

func (c *Context) setParams(params map[string]string) {
	c.params = params
}

func (c *Context) reset(w http.ResponseWriter, r *http.Request) {
	c.Writer = NewResponseWriter(w)
	c.Request = r
}
