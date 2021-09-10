package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
  r:=gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

   r.GET("/foo", func(c *gin.Context) {
	   c.JSON(http.StatusOK,"被重定向了。。。。")
   })
  r.POST("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})


  r.Run(":8080")
}
