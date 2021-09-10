package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func main() {
//	router := gin.Default()
//   //http://localhost:8888/user/lktbz
//	// This handler will match /user/john but will not match /user/ or /user
//	router.GET("/user/:name", func(c *gin.Context) {
//		//获取请求参数中的name
//		name := c.Param("name")
//		c.String(http.StatusOK, "Hello %s", name)
//	})
//	router.Run(":8888")
//}
