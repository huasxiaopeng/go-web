package main

import (
	"fmt"
	"io/ioutil"
)

/**
bufio.NewWriter  写入
 */
func main() {
	str := "hello 沙河"
	err := ioutil.WriteFile("hello.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
