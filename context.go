package goblock

import (
	"encoding/json"
	"net/http"
)

type context interface {
	GetWriter() http.ResponseWriter

	GetRequest() *http.Request

	Json(statusCode int, data map[string]interface{})
}

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) reset(w http.ResponseWriter, r *http.Request) {
	c.writer = w
	c.request = r
}

func (c *Context) GetWriter() http.ResponseWriter {
	return c.writer
}

func (c *Context) GetRequest() *http.Request {
	return c.request
}

func (c *Context) Json(statusCode int, data map[string]interface{}) {
	c.writer.Header().Set("Content-Type", "application/json")
	c.writer.WriteHeader(statusCode)

	err := json.NewEncoder(c.writer).Encode(data)
	if err != nil {
		panic(err)
	}
}
