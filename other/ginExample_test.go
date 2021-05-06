/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-06 07:01:38
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-06 07:04:41
 * @Description:
 * @FilePath: /Lin/other/ginExample.go
 * @GitHub: https://github.com/liuyuchen777
 */
package other

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	r.Run(":9090")
}
