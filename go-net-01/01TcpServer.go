package main

import (
	"bufio"
	"fmt"
	"net"
)
// TCP server端

// 处理函数
func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err!=nil {
		fmt.Println("监听失败")
		return
	}
	fmt.Println("启动成功")
	for  {
		conn, err := listen.Accept()//建立链接
		if err!=nil{
			fmt.Println("链接建立失败。。。")
			break
		}
        go procss(conn)// 启动一个goroutine处理连接
	}
}
func procss(conn net.Conn) {
	defer conn.Close() //处理完毕完毕链接
	reader:=bufio.NewReader(conn)
	var buf[128]byte
	read, err := reader.Read(buf[:])
	if err!=nil{
		fmt.Println("读取失败")
		return
	}
    rescStr:=string(buf[:read])
	fmt.Println("收到client端发来的数据：", rescStr)
	conn.Write([]byte(rescStr)) // 发送数据
}
