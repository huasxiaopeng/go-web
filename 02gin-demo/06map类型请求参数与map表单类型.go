package main

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
///**
//POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
//Content-Type: application/x-www-form-urlencoded
//names[first]=thinkerou&names[second]=tianou
// */
//func main() {
//	r:=gin.Default()
//	r.POST("/post", func(c *gin.Context) {
//		ids:=c.QueryMap("ids")
//		names:=c.PostFormMap("names")
//		fmt.Printf("ids: %v; names: %v", ids, names)
//	})
//	r.Run(":8080")
//
//}
