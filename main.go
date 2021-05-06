/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-06 06:42:10
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-06 07:09:02
 * @Description:
 * @FilePath: /Lin/main.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"Lin"
	"fmt"
	"net/http"
)

func main() {
	r := lin.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9090")
}
