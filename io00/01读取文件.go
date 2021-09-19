package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("readme.md")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	 defer file.Close()
	//读取数据
	var tmp=make([]byte,128)
	n, err := file.Read(tmp)
	if err==io.EOF{
		fmt.Println("文件读取完毕")
		return
	}
	if err!=nil{
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}
