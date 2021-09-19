package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	strs:="hellos"
	file.Write([]byte(strs))//写入字节
	file.WriteString("\nhello 写入的是字符串")

}
