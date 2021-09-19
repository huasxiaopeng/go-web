package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

/**
   空值处理
 */
func main() {
	var (
		userName  string = "root"
		password  string = "1234"
		ipAddrees string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go"
		charset   string = "utf8"
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)

	db, err := sqlx.Open("mysql", dsn)
	if err!=nil{
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	//queryRow(db)

	queryRows(db)
}

func queryRows(db *sqlx.DB) {

	type person struct {
		Uid int `db:"uid"`
		UserName string `db:"username"`

	}

	var perData person
	rows, err := db.Queryx("select uid,username from person")
	if err != nil {
		fmt.Printf("query data failed, error is [%v]", err.Error())
		return
	}
	var perDataSlice[] person
	for rows.Next(){
		err:=rows.StructScan(&perData)
		if err != nil {
			fmt.Printf("scan data failed, error is [%v]", err.Error())
			return
		}
		perDataSlice=append(perDataSlice,perData)
	}
	fmt.Println("多行查询结果为：",perDataSlice)
}
func queryRow(Db *sqlx.DB) {
	row := Db.QueryRow("select uid, username, create_time from person where uid=13")

	var uid int
	var userName string
	//var createTime sql.NullString //防止null 时间报错
	var createTime time.Time  //时间为空，报错
	err := row.Scan(&uid, &userName, &createTime)
	if err != nil {
		fmt.Printf("scan data failed, error is [%v]", err.Error())
		return
	}
	fmt.Println(uid, userName, createTime)
}

