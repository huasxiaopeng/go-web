package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
/**
    文件读取步骤：
    首先指定文件名
    读取文件的方式，是读取二进制还是串
    然后打印读取的内容
 */
func main() {
	file, err := os.Open("readme.md")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for  {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}
