package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/root/codecloud/workplace/code.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	// 使用Read方法读取数据
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("read successfully")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}
