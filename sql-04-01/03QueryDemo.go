package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var (
	userName  string = "root"
	password  string = "1234"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "go"
	charset   string = "utf8"
)
//查询学习
func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)

	db, err := sqlx.Open("mysql", dsn)
	if err!=nil{
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	
	//queryRowsDemo(db)
	//queryxRowsDemo(db)
	queryRowxDemo(db)
}

func queryRowxDemo(db *sqlx.DB) {
	//定义结构体保存数据
	type person struct {
		Uid int `db:"uid"`
		UserName string `db:"username"`
		CreateTime string `db:"create_time"`
	}

	var userData *person = new(person)
	row := db.QueryRowx("select uid, username, create_time from userinfo where uid = 1")
	err := row.StructScan(userData)
	if err != nil {
		fmt.Printf("scan data failed, error is [%v]", err.Error())
	}

	fmt.Println(userData.Uid, userData.UserName, userData.CreateTime)
}
//Queryx() 和 QueryRowx()，不仅支持 Scan() 方法，同时可将数据与结构体进行转换
func queryxRowsDemo(db *sqlx.DB) {
	//定义结构体保存数据
	type person struct {
		Uid int `db:"uid"`
		UserName string `db:"username"`
		CreateTime string `db:"create_time"`
	}

	var userData person
	rows, err := db.Queryx("select uid, username, create_time from person")
	if err != nil {
		fmt.Printf("query data failed, error is [%v]", err.Error())
		return
	}
	var userDataSlice []person
	for rows.Next() {
		err := rows.StructScan(&userData)
		if err != nil {
			fmt.Printf("scan data failed, error is [%v]", err.Error())
			return
		}
		userDataSlice = append(userDataSlice, userData)
	}
	fmt.Println(userDataSlice)
	err = rows.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func queryRowsDemo(db *sqlx.DB) {
	row := db.QueryRow("select uid, username, create_time from person")
	var uid int
	var userName, createTime string
	err := row.Scan(&uid, &userName, &createTime)
	if err != nil {
		fmt.Printf("scan data failed, error is [%v]", err.Error())
		return
	}
	fmt.Println(uid, userName, createTime)
}
