package main

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
///**
//
//   postman 调用
//POST /post?id=1234&page=1 HTTP/1.1
//Content-Type: application/x-www-form-urlencoded
//
//name=manu&message=this_is_grea
// */
//func main() {
//  r:=gin.Default()
//  r.POST("/post", func(c *gin.Context) {
//	 id:= c.Query("id") //请求户路径获取数据
//  	 page:=c.DefaultQuery("page","0")
//  	 name:=c.PostForm("name")
//  	  message:=c.PostForm("message")
//	  fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
//
//  })
//	r.Run(":8080")
//
//}
