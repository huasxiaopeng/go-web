package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
/**
  定义模型
 */
type Product struct {
	gorm.Model  //使用这个，创建的表会带几个字段CreatedAt UpdatedAt DeletedAt  ID
	Code  string
	Price uint
}
func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		fmt.Println("链接失败")
	}
	//自动生成表
	db.AutoMigrate(&Product{})
	//创建记录
	create := db.Create(&Product{Code: "D42", Price: 100})
	fmt.Println(create)
}
