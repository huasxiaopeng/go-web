package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

)

//简单查询使用
func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sqlx.Open("mysql", dsn)
	if err!=nil{
		panic("链接失败，请重试")
	}
	//queryData(db)
	//queryDataWithGet(db)
	querySelect(db)
}
/**
   查询多条数据
 */
func querySelect(db *sqlx.DB) {
	type person struct {
		Uid int `db:"uid"`
		UserName string `db:"username"`
		CreateTime string `db:"create_time"`
		Password string `db:"password"`
		Department string `db:"department"`
		Email string `db:"email"`
	}
	var userInfoSlice []person
	err := db.Select(&userInfoSlice, "select *from person")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}

	//遍历结构体切片
	for _, userData := range userInfoSlice {
		fmt.Println(userData.Uid, userData.CreateTime, userData.UserName,
			userData.Password, userData.Department, userData.Email)
	}
}
//映射结构体  查询单一数据
func queryDataWithGet(db *sqlx.DB) {
	type person struct {
		Uid int `db:"uid"`
		UserName string `db:"username"`
		CreateTime string `db:"create_time"`
		Password string `db:"password"`
		Department string `db:"department"`
		Email string `db:"email"`
	}
	var userData *person=new(person)
	err := db.Get(userData, "select *from person where uid=1")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	fmt.Println(userData.Uid, userData.CreateTime, userData.UserName,
		userData.Password, userData.Department, userData.Email)
}
/**
   简单查询
 */
func queryData(db *sqlx.DB) {
	rows, err := db.Query("select  *from person")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for rows.Next(){
		//使用变量接收数据
		var uid int
		var create_time, username, password, department, email string
		err := rows.Scan(&uid, &create_time, &username, &password, &department, &email)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		fmt.Println(uid, create_time, username, password, department, email)
	}

	rows.Close()
}
