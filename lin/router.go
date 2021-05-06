/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-06 07:20:02
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-06 07:41:07
 * @Description:
 * @FilePath: /Lin/lin/router.go
 * @GitHub: https://github.com/liuyuchen777
 */
package lin

import (
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
