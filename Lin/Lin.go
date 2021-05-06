/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-06 06:42:32
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-06 07:10:18
 * @Description:
 * @FilePath: /Lin/Lin/Lin.go
 * @GitHub: https://github.com/liuyuchen777
 */
package lin

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	/* router table */
	router map[string]HandlerFunc
}

/*
 * create a default no log router
 */
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

/*
 * add route to route map
 */

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

/*
 * engine realize ServeHTTP interface
 */
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 page not found: %s\n", req.URL)
	}
}
