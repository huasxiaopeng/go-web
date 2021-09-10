package main

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
///**
//   多版本管理
// */
//func main() {
//	router.Run()
//}
//var (
//	router = gin.Default()
//)
//
//// Run will start the server
//func Run() {
//	getRoutes()
//	router.Run(":8888")
//}
//
//// getRoutes will create our routes of our entire application
//// this way every group of routes can be defined in their own file
//// so this one won't be so messy
//func getRoutes() {
//	v1 := router.Group("/v1")
//	addUserRoutes(v1)
//	addPingRoutes(v1)
//
//	v2 := router.Group("/v2")
//	addPingRoutes(v2)
//}
//func addPingRoutes(rg *gin.RouterGroup) {
//	ping := rg.Group("/ping")
//
//	ping.GET("/", func(c *gin.Context) {
//		c.JSON(http.StatusOK, "pong")
//	})
//}
//func addUserRoutes(rg *gin.RouterGroup) {
//	users := rg.Group("/users")
//
//	users.GET("/", func(c *gin.Context) {
//		c.JSON(http.StatusOK, "users")
//	})
//	users.GET("/comments", func(c *gin.Context) {
//		c.JSON(http.StatusOK, "users comments")
//	})
//	users.GET("/pictures", func(c *gin.Context) {
//		c.JSON(http.StatusOK, "users pictures")
//	})
//}