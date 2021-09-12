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
	queryRow(db)
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