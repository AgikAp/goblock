package goblock

import (
	"log"
	"strings"
)

type RouterNode struct {
	key      string
	val      map[string]*Route
	children map[string]*RouterNode
	param    *RouterNode
}

func NewNode(key string) *RouterNode {
	return &RouterNode{
		key:      key,
		val:      make(map[string]*Route),
		children: make(map[string]*RouterNode),
	}
}

type RouterTree struct {
	root *RouterNode
}

func NewRouterTree() *RouterTree {
	return &RouterTree{root: &RouterNode{children: make(map[string]*RouterNode)}}
}

func (t *RouterTree) Insert(method string, path string, handlers []HandlerFunc) {
	node := t.root
	parts := splitPath(path)
	var paramNames []string

	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			paramNames = append(paramNames, part[1:])

			if node.param == nil {
				node.param = NewNode("*")
			}

			node = node.param
		} else {
			if _, exists := node.children[part]; !exists {
				node.children[part] = NewNode(part)
			}

			node = node.children[part]
		}
	}

	node.val[method] = NewRouter(handlers, paramNames)
}

func (t *RouterTree) Search(method string, path string) (*Route, error) {
	node := t.root
	parts := splitPath(path)
	params := make(map[string]string, len(parts))

	paramCounter := 0
	for _, part := range parts {
		if child, exists := node.children[part]; exists {
			node = child
		} else if node.param != nil {
			node = node.param
			paramNames := node.val[method].paramNames

			if paramCounter < len(paramNames) {
				params[paramNames[paramCounter]] = part
				paramCounter++
			}
		} else {
			return nil, NewHttpError(404, "Resource not found", nil)
		}
	}

	log.Println(node.val)

	if route, exists := node.val[method]; exists {
		route.params = params
		return route, nil
	}

	return nil, NewHttpError(405, "Resource not allowed", nil)
}
