package main

import (
	"fmt"
	"io"

	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err!=nil{
		fmt.Println("响应失败")
		return
	}
	defer resp.Body.Close()
	all, _ := io.ReadAll(resp.Body)
	fmt.Println(string(all))


}
