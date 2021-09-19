package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端启动成功")
	inputReader:=bufio.NewReader(os.Stdin)
	for{
		intput,_:=inputReader.ReadString('\n')
		intputInfo:= strings.Trim(intput,"\r\n")
		if strings.ToUpper(intputInfo)=="Q"{
			return
		}
		_,err:=conn.Write([]byte(intputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
