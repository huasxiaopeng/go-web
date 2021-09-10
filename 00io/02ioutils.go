package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)
/**
   ioutils 读取完整的文件
 */
func main() {
   //ioReadAll()
	//ioReadDir()
	//ioReadFile()
	ioWriteFile()
}
func ioWriteFile()  {
  message:=	[]byte("HELLO jinGOM")
	err:=ioutil.WriteFile("hello.txt",message,0644)
	if err != nil {
		log.Fatal(err)
	}
}
func ioReadFile(){
	file, err := ioutil.ReadFile("E:\\记事本.txt")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s",file)
}
func ioReadDir()  {
	//files, err := ioutil.ReadDir(".")
	files, err := ioutil.ReadDir("D:\\") //windows
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
func ioReadAll()  {
	r :=  strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	all, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s",all)
}