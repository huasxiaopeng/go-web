package main

//import (
//	"github.com/gin-gonic/gin"
//	"log"
//)
//
//type Person struct {
//	Name    string `form:"name"`
//	Address string `form:"address"`
//}
////localhost:8085/testing?name=lktbz&address=北京  传递过来的参数，绑定到person 中
//func main() {
//  r:= gin.Default()
//  r.Any("/testing",startPage)
//	  r.Run(":8085")
//}
//
//func startPage(c *gin.Context) {
//	var person Person
//	//查看结构体是否为空
//	if c.ShouldBindQuery(&person)==nil{
//		log.Println("====== Only Bind By Query String ======")
//		log.Println(person.Name)
//		log.Println(person.Address)
//	}
//	//返回字符串
//	c.String(200,"success")
//}
