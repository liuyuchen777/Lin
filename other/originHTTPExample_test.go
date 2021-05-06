/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-06 06:39:17
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-06 06:54:36
 * @Description:
 * @FilePath: /Lin/other/originHTTPExample_test.go
 * @GitHub: https://github.com/liuyuchen777
 */
package other

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestOrigin(t *testing.T) {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
