package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//普通查询所使用

type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
}
//type CreditCard struct {
//	gorm.Model
//	Number   string
//	UserID   uint
//}
//type User struct {
//	gorm.Model
//	Name       string `gorm:"default:galeone"`
//	CreditCard CreditCard
//	Age int64 `gorm:"default:18"`  //默认值，当不插入数据时，会填写默认的
//}
func main() {

	dsn := "root:1234@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		fmt.Println("链接失败")
	}
	var user=&User{}
   //searchfirst(db,user)
   //searchTake(db,user)
	//searchLast(db,user)
	//searchById(db,user)
	searchFindIn(db,user)
}
/**
 sql in 写法
 */
func searchFindIn(db *gorm.DB, user *User) {
}
//// SELECT * FROM users WHERE id = 10; 等价于数据库
func searchById(db *gorm.DB, user *User) {
	 db.First(&user, 10).Scan(&user)
    fmt.Println(user)
}
/**
  查询最后一个元素
 */
func searchLast(db *gorm.DB, user *User) {
	db.Last(&user).Scan(&user)
	fmt.Println(user)
}

//不进行排序操作 都是查询单挑数据
func searchTake(db *gorm.DB, user User) {
	db.Take(&user).Scan(&user)
	fmt.Println(user.Age)
	fmt.Println(user.Name)
	fmt.Println(user.Email)
}
//简单查询
func searchfirst(db *gorm.DB,user User) {
	first := db.First(&user) //会根据指定主键进行排序
	//获取第一条数据
	fmt.Println(first.Row().Scan(&user))
	fmt.Println(user)
}



