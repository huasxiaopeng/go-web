package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
bufio.NewWriter  写入
 */
func main() {
	file, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i <10 ; i++ {
		writer.WriteString("hello 回龙观\n")

	}
	writer.Flush()
}
