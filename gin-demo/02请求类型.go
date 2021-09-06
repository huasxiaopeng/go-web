package main

import "github.com/gin-gonic/gin"

/**
   构建请求的类型
 */
func main() {
	router := gin.Default()
	/**
	 请求的类型 get post put delete
	          patch head options
	 */
	router.GET("/somget")
	router.POST("/somePost")
	router.PUT("/somePut")
	router.DELETE("/someDelete")
	router.PATCH("/somePatch")
	router.HEAD("/someHead")
	router.OPTIONS("/someOptions")

	router.Run()
}
