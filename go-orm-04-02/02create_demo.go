package main
//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
////普通查询所使用
//type User struct {
//	ID           uint
//	Name         string
//	Email        string
//	Age          uint8
//}
//
//
////type CreditCard struct {
////	gorm.Model
////	Number   string
////	UserID   uint
////}
////type User struct {
////	gorm.Model
////	Name       string `gorm:"default:galeone"`
////	CreditCard CreditCard
////	Age int64 `gorm:"default:18"`  //默认值，当不插入数据时，会填写默认的
////}
//func main() {
//	dsn := "root:1234@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err!=nil{
//		fmt.Println("链接失败")
//	}
//	//自动生成表
//	//db.AutoMigrate(&CreditCard{})
//	db.AutoMigrate(&User{})
//	//单挑记录插入
//	//user:=User{Name: "lktbz",Age: 18,Email: "666@66.com"}
//	//users := db.Create(&user) // 通过数据的指针来创建
//	//fmt.Println(user.ID)// 返回插入数据的主键
//	//fmt.Println(users.Error) // 返回 error
//	//fmt.Println(users.RowsAffected)// 返回插入记录的条数
//
//
//	//read
//	//var us User
//	//first := db.First(&us, 1)
//	//err = first.Row().Scan(&us)
//	//fmt.Println(us)
//
//	////创建记录，并根据指定字段插入
//	//user:=User{Name: "jk",Age: 20,Email: "777@66.com"}
//	//_ = db.Select("Name", "Age").Create(&user)
//
//	//批量插入字段
//	//var users=[]User{{Name: "zl",Age: 20},{Name: "wz"},{Name: "ks"}}
//	//db.Create(&users)
//	//for _, user := range users {
//	//	fmt.Println(user.ID) //查看插入的id
//	//}
//
//
//	//指定大小的批量
//	//var users = []User{{Name: "jinzhu_1"}, {Name: "jinzhu_10000"}}
//	//db.CreateInBatches(users,100)
//    // 上面的修改下
//
//    var users=make([]User,200)
//	for i := 0; i <200 ; i++ {
//		users[i]=User{Name: "beij"}
//	}
//	db.CreateInBatches(users,100)
//	//Map 方式创建
//	//db.Model(&User{}).Create(map[string]interface{}{
//	//	"Name":"GimiJiang","Age":18,
//	//})
//
//   //批量
//	//db.Model(&User{}).Create([]map[string]interface{}{
//	//	{"Name": "jinzhu_1", "Age": 18},
//	//	{"Name": "jinzhu_2", "Age": 20},
//	//})
//
//    //关联创建
//	//db.Create(&User{
//	//	Name: "zhizhu",
//	//	CreditCard: CreditCard{Number: "41111111"},
//	//})
//
//}
//
//
