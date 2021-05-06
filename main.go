/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-06 07:14:03
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-06 07:40:27
 * @Description:
 * @FilePath: /Lin/main.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"lin"
	"net/http"
)

// func main() {
// 	r := lin.New()

// 	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
// 		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
// 	})

// 	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
// 		for k, v := range req.Header {
// 			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 		}
// 	})

// 	r.Run(":9090")
// }

func main() {
	r := lin.New()

	r.GET("/", func(c *lin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello lin</h1>")
	})

	r.GET("/hello", func(c *lin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *lin.Context) {
		c.JSON(http.StatusOK, lin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9090")
}
