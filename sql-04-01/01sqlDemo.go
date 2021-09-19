package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

/**
    在学习orm 过程中，发现有些遍历方式，以及api 有点不清楚，现在补下标准库中的sql操作
 */
func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err!=nil{
	 panic("链接失败，请重试")
	}
	addRecord(db)
	//updateRecord(db)
	//deleteRecord(db)
}
/**
  删除操作
 */
func deleteRecord(db *sql.DB) {
	//删除uid=2的数据
	result, err := db.Exec("delete from userinfo where uid = 2")
	if err != nil {
		fmt.Printf("delete faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("delete success, affected rows:[%d]\n", num)
}

//更新操作
func updateRecord(db *sql.DB) {
	result, err := db.Exec("update person set username ='lktbz' where uid=1")
	if err != nil {
		fmt.Printf("update faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected() //根据返回的 值可以看出是否更新成功
	fmt.Printf("update success, affected rows:[%d]\n", num)
}
/**
   数据的插入
 */
func addRecord(db *sql.DB) {
	for i := 0; i <10 ; i++ {
		result, err := db.Exec("insert into person  values(?,?,?,?,?,?)",0, "2019-07-06 11:45:20", "johnys", "123456", "技术部", "123456@163.com")
		if err != nil {
			fmt.Printf("data insert faied, error:[%v]", err.Error())
			return
		}
		id,_ := result.LastInsertId()
		fmt.Printf("insert success, last id:[%d]\n", id)
	}

}

